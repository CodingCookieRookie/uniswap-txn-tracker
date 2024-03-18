package mysql

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
)

func ReplaceTransactionBulkByBatch(transactions []*model.Txn, timestampToEthPriceMap *map[uint64]string) error {
	currentLen := 0
	for currentLen < len(transactions) {
		if len(transactions) <= maxBatchSize {
			return ReplaceTransactionsBulk(transactions, timestampToEthPriceMap)
		}
		batch := transactions[:maxBatchSize]
		err := ReplaceTransactionsBulk(batch, timestampToEthPriceMap)
		if err != nil {
			log.Errorf("error replacing transaction bulk in batch function: %v", err)
		}
		currentLen += maxBatchSize
	}
	return nil
}

// Generates shortened and unique id to store transactions.
func generateUniqueTxnId(txn *model.Txn, length int) string {
	if txn == nil {
		return ""
	}
	input := fmt.Sprintf("%v-%v-%v-%v", txn.BlockNumber, txn.From, txn.To, txn.Hash)
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hashBytes)
	if length > 0 && length < len(encoded) {
		return encoded[:length]
	}
	return encoded
}

func ReplaceTransactionsBulk(txns []*model.Txn, timestampToEthPriceMap *map[uint64]string) error {
	if len(txns) == 0 {
		return nil
	}

	if timestampToEthPriceMap == nil {
		log.Error("timestampToEthPriceMap is nil")
		return nil
	}
	var sb strings.Builder
	args := []any{}

	for _, tx := range txns {
		id := generateUniqueTxnId(tx, 64)
		timeStamp, err := strconv.ParseInt(tx.TimeStamp, 10, 64)
		if err != nil {
			log.Errorf("error converting txn timestamp to int64: %v", err)
			continue
		}
		ethPrice, exists := (*timestampToEthPriceMap)[uint64(timeStamp)]
		if !exists {
			log.Errorf("timeStamp does not exist in eth mapping: %v", uint64(timeStamp))
			return nil
		}

		arg := []any{id, tx.BlockNumber, tx.TimeStamp, tx.Hash, tx.Nonce, tx.BlockHash, tx.From, tx.ContractAddress, tx.To, tx.Value, tx.TokenName, tx.TokenSymbol, tx.TokenDecimal, tx.TransactionIndex, tx.Gas, tx.GasPrice, tx.GasUsed, tx.CumulativeGasUsed, tx.Input, tx.Confirmations, ethPrice}
		sb.WriteString(fmt.Sprintf("%v,", returnPlaceHolderString(arg)))
		args = append(args, arg...)
	}
	query := sb.String()
	if index := strings.LastIndex(sb.String(), ","); index > 0 {
		query = query[:index]
	}
	return exec("REPLACE INTO uniswap_db.transactions (`id`, `blockNumber`, `timeStamp`, `hash`, `nonce`, `blockHash`, `fromAddress`, `contractAddress`, `toAddress`, `value`, `tokenName`, `tokenSymbol`, `tokenDecimal`, `transactionIndex`, `gas`, `gasPrice`, `gasUsed`, `cumulativeGasUsed`, `input`, `confirmations`, `ethPrice`) VALUES "+query, args...)
}

func GetTransactionsByTimestamp(startTime uint64, endTime uint64) ([]*model.Txn, error) {
	return query(func(tx *model.Txn) []any {
		return []any{
			&tx.BlockNumber, &tx.TimeStamp, &tx.Hash, &tx.Nonce, &tx.BlockHash, &tx.From, &tx.ContractAddress, &tx.To, &tx.Value, &tx.TokenName, &tx.TokenSymbol, &tx.TokenDecimal, &tx.TransactionIndex, &tx.Gas, &tx.GasPrice, &tx.GasUsed, &tx.CumulativeGasUsed, &tx.Input, &tx.Confirmations, &tx.Ethprice,
		}
	}, "SELECT `blockNumber`, `timeStamp`, `hash`, `nonce`, `blockHash`, `fromAddress`, `contractAddress`, `toAddress`, `value`, `tokenName`, `tokenSymbol`, `tokenDecimal`, `transactionIndex`, `gas`, `gasPrice`, `gasUsed`, `cumulativeGasUsed`, `input`, `confirmations`, `ethPrice` FROM uniswap_db.transactions WHERE timeStamp >= ? and timeStamp <= ?", startTime, endTime)
}

func GetTxnDetailsByTxnHash(txnHash string) (*model.TxnFeeDetails, error) {
	return queryRow(func(txnFeeDetails *model.TxnFeeDetails) []any {
		return []any{
			&txnFeeDetails.GasUsed, &txnFeeDetails.GasPrice, &txnFeeDetails.EthPrice,
		}
	}, "SELECT `gasUsed`, `gasPrice`, `ethPrice` FROM uniswap_db.transactions WHERE hash = ?", txnHash)
}
