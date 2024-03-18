package service

import (
	"math"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/errors"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/log"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/model"
	"github.com/CodingCookieRookie/uniswap-txn-tracker/mysql"
)

// Gets historical transactions from/to address from start to end time.
func GetHistoricalTxnsService(startTime, endTime uint64) (*model.TxnsResp, error) {
	txns, err := mysql.GetTransactionsByTimestamp(startTime, endTime)

	if err != nil {
		log.Errorf("error getting transactions from db, error: %v", err)
		if len(txns) == 0 {
			return nil, &errors.ServerError{
				Msg: "Please wait some time for the server to pull your data if input is valid",
			}
		}
		return nil, &errors.ServerError{
			Msg: err.Error(),
		}
	}

	return &model.TxnsResp{
		Result: txns,
	}, err
}

func weiToEth(gasPrice uint64) float64 {
	return float64(gasPrice) / math.Pow(10, 18)
}

func calculateTransactionFee(txnFeeDetails *model.TxnFeeDetails) float64 {
	gasUsed := txnFeeDetails.GasUsed
	gasPrice := txnFeeDetails.GasPrice
	ethPrice := txnFeeDetails.EthPrice

	txnFee := float64(gasUsed) * weiToEth(gasPrice) * float64(ethPrice)

	return txnFee
}

// Gets transaction fee of transaction with corresponding transaction hash.
func GetTransactionFeeService(txnHash string) (*model.TxnFeeResp, error) {
	txnDetails, err := mysql.GetTxnDetailsByTxnHash(txnHash)
	if err != nil {
		log.Errorf("error getting transaction details from db, error: %v", err)
		return nil, &errors.ServerError{
			Msg: err.Error(),
		}
	}

	return &model.TxnFeeResp{
		TxnFee: calculateTransactionFee(txnDetails),
	}, nil
}
