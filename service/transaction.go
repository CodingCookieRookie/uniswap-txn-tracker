package service

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/request"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/response"
	"github.com/gin-gonic/gin"
)

// Gets historical transactions from/to address from start to end time.
func GetHistoricalTxns(historicalTxnReq request.HistoricalTxnReq) (*response.TxnsResp, error) {
	// TODO: service impl
	return nil, nil
}

// Gets transaction fee of transaction with corresponding transaction hash.
func GetTransactionFee(c *gin.Context) (*response.TxnFeeResp, error) {
	// TODO: service impl
	return nil, nil
}
