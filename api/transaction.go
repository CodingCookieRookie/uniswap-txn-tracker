package api

import (
	"time"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/request"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/errors"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/service"
	"github.com/gin-gonic/gin"
)

const timeLayout = "2006-01-02 15:04:05"

// GetHistoricalTxns godoc
// @Summary		returns all historical transactions from start to end time
// @Description	returns all historical transactions from start to end time
// @Tags			accounts
// @Accept			json
// @Produce		json
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

	startTime, err := time.Parse(timeLayout, historicalTxnReq.StartTime)
	if err != nil {
		return nil, &errors.UserError{
			Msg: "start time must be in correct format, eg. 2006-01-02 15:04:05",
		}
	}

	endTime, err := time.Parse(timeLayout, historicalTxnReq.EndTime)
	if err != nil {
		return nil, &errors.UserError{
			Msg: "end time must be in correct format, eg. 2006-01-02 15:04:05",
		}
	}

	startTimeUnix := uint64(startTime.Unix())
	endTimeUnix := uint64(endTime.Unix())
	log.Debugf("startTimeUnix: %v, endTimeUnix: %v", startTimeUnix, endTimeUnix)

	return service.GetHistoricalTxnsService(startTimeUnix, endTimeUnix)
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
	txnHash := c.Query("txn_hash")
	return service.GetTransactionFeeService(txnHash)
}
