package api

import (
	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/request"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/errors"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/gin-gonic/gin"
)

// GetHistoricalTxns godoc
// @Summary		returns all historical transactions from start to end time
// @Description	returns all historical transactions from start to end time
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			address		query		string	true	"Address"
// @Param			start_time	query		string	true	"Start Time in ISO 8601 format"
// @Param			end_time	query		string	true	"End Time in ISO 8601 format"
// @Success		200			{object}	response.TxnsResp
// @Failure		400			{object}	errors.UserError	"Bad Request"
// @Failure		500			{object}	errors.ServerError	"Server Error"
// @Router			/transactions [get]
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

// GetTransactionFee godoc
// @Summary		returns transaction fee of transaction with corresponding transaction hash
// @Description	returns transaction fee of transaction with corresponding transaction hash
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			txn_hash		query		string	true	"Transaction hash"
// @Success		200			{object}	response.TxnFeeResp
// @Failure		500			{object}	errors.ServerError	"Server Error"
// @Router			/transaction/fee [get]
func GetTransactionFee(c *gin.Context) (any, error) {
	var transactionFeeRequest request.TxnFeeReq
	if err := c.BindQuery(&transactionFeeRequest); err != nil {
		log.Errorf("error binding transaction fee request: %v", err)
		return nil, &errors.UserError{Msg: err.Error()}
	}
	// TODO: service impl
	return nil, nil
}
