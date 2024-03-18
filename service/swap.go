package service

import (
	"fmt"
	"math/big"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/mysql"
)

func GetUniswapSwapPrice(txnHash string) (*model.SwapResponse, error) {
	swapEventsWithToken, err := mysql.GetSwapPricesByTxnHash(txnHash)
	swapPrices := make([]*model.SwapPrice, 0)
	if err == nil && swapEventsWithToken != nil {
		for _, e := range swapEventsWithToken {
			swapPrice := calculateSwapPrice(e)
			swapPrices = append(swapPrices, swapPrice)
		}
	}
	if err != nil {
		log.Errorf("error getting transactions from db, error: %v", err)
	}

	// try get from API
	return &model.SwapResponse{
		SwapPrices: swapPrices,
	}, nil
}

func calculateSwapPrice(swapEventWithToken *model.SwapEventWithToken) *model.SwapPrice {

	// Use the SetString method to create a big.Int from the string
	sqrtPriceX96 := new(big.Int)
	sqrtPriceX96.SetString(swapEventWithToken.SqrtPriceX96, 10)

	price1 := new(big.Float).SetInt(sqrtPriceX96)
	price1.Quo(price1, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)))
	price1.Mul(price1, price1) // square it to get price ratio

	price0 := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil))
	price0.Quo(price0, new(big.Float).SetInt(sqrtPriceX96))
	price0.Quo(price0, new(big.Float).SetInt(sqrtPriceX96)) // square root and then inverse it to get price ratio

	priceText := fmt.Sprintf("Price of token2 in terms of %v: %s\n", swapEventWithToken.TokenSymbol, price0.Text('f', 18))
	return &model.SwapPrice{
		SwapEventWithToken: swapEventWithToken,
		Price_Text:         priceText,
	}
}
