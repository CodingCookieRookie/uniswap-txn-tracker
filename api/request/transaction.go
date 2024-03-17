package request

// A TransactionsRequest will get all transactions from an address from start to end time.
type HistoricalTxnReq struct {
	Address   string `form:"start_time" binding:"required"`
	StartTime string `form:"start_time" binding:"required"` // start time in ISO 8601
	EndTime   string `form:"end_time" binding:"required"`   // end time in ISO 8601
}

// A TransactionRequest will get the transaction fee of the corresponding transaction in USDT.
type TxnFeeReq struct {
	TxnHash string `form:"txn_hash" binding:"required"` // transaction hash
}
