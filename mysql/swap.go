package mysql

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
)

func generateUniqueSwapId(swap *model.SwapEvent, length int) string {
	if swap == nil {
		return ""
	}
	input := fmt.Sprintf("%v-%v-%v", swap.From, swap.To, swap.TxnHash)
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hashBytes)
	if length > 0 && length < len(encoded) {
		return encoded[:length]
	}
	return encoded
}

func ReplaceSwapBulkByBatch(swaps []*model.SwapEvent) error {
	currentLen := 0
	for currentLen < len(swaps) {
		if len(swaps) <= maxBatchSize {
			return ReplaceSwapsBulk(swaps)
		}
		batch := swaps[:maxBatchSize]
		err := ReplaceSwapsBulk(batch)
		if err != nil {
			log.Errorf("error replacing swap bulk in batch function: %v", err)
		}
		currentLen += maxBatchSize
	}
	return nil
}

func ReplaceSwapsBulk(swaps []*model.SwapEvent) error {
	if len(swaps) == 0 {
		return nil
	}

	var sb strings.Builder
	args := []any{}

	for _, swap := range swaps {
		id := generateUniqueSwapId(swap, 64)
		arg := []any{id, swap.TxnHash, swap.SqrtPriceX96, swap.From, swap.To}
		sb.WriteString(fmt.Sprintf("%v,", returnPlaceHolderString(arg)))
		args = append(args, arg...)
	}
	query := sb.String()
	if index := strings.LastIndex(sb.String(), ","); index > 0 {
		query = query[:index]
	}
	return exec("REPLACE INTO uniswap_db.swaps (`id`, `txnHash`, `sqrtPriceX96`, `fromAddress`, `toAddress`) VALUES "+query, args...)
}

func GetSwapPricesByTxnHash(txnHash string) ([]*model.SwapEventWithToken, error) {
	return query(func(swap *model.SwapEventWithToken) []any {
		return []any{
			&swap.SqrtPriceX96, &swap.From, &swap.To, &swap.TokenName, &swap.TokenSymbol, &swap.TokenDecimal,
		}
	},

		`
		SELECT s.sqrtPriceX96, s.fromAddress, s.toAddress, t.tokenName, t.tokenSymbol, t.tokenDecimal
		FROM
		uniswap_db.swaps as s 
		JOIN 
		uniswap_db.transactions AS t
		ON s.txnHash = t.hash
		WHERE s.txnHash = ?`, txnHash)
}
