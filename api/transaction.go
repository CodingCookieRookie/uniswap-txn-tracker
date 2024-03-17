package api

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/request"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/errors"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/gin-gonic/gin"
)

// Gets historical transactions from/to address from start to end time.
func GetHistoricalTxns(c *gin.Context) (any, error) {
	var historicalTxnReq request.HistoricalTxnReq
	if err := c.BindQuery(&historicalTxnReq); err != nil {
		log.Errorf("error binding historical txn request: %v", err)
		return nil, &errors.UserError{Msg: err.Error()}
	}
	// TODO: checks on input

	// TODO: service impl
	return nil, nil
}

// Gets transaction fee of transaction with corresponding transaction hash.
func GetTransactionFee(c *gin.Context) (any, error) {
	var transactionFeeRequest request.TxnFeeReq
	if err := c.BindQuery(&transactionFeeRequest); err != nil {
		log.Errorf("error binding transaction fee request: %v", err)
		return nil, &errors.UserError{Msg: err.Error()}
	}
	// TODO: service impl
	return nil, nil
}
