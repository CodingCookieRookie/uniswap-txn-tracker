package model

type TxnsResp struct {
	Result []*Txn `json:"result"`
}

type Txn struct {
	BlockNumber       string  `json:"blockNumber"`
	TimeStamp         string  `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             string  `json:"nonce"`
	BlockHash         string  `json:"blockHash"`
	From              string  `json:"from"`
	ContractAddress   string  `json:"contractAddress"`
	To                string  `json:"to"`
	Value             string  `json:"value"`
	TokenName         string  `json:"tokenName"`
	TokenSymbol       string  `json:"tokenSymbol"`
	TokenDecimal      string  `json:"tokenDecimal"`
	TransactionIndex  string  `json:"transactionIndex"`
	Gas               string  `json:"gas"`
	GasPrice          string  `json:"gasPrice"`
	GasUsed           string  `json:"gasUsed"`
	CumulativeGasUsed string  `json:"cumulativeGasUsed"`
	Input             string  `json:"input"`
	Confirmations     string  `json:"confirmations"`
	Ethprice          float64 `json:"ethprice"`
}

type TxnFeeResp struct {
	TxnFee float64 `json:"transactionFee"`
}

type TxnFeeDetails struct {
	GasPrice uint64  `json:"gasPrice"`
	GasUsed  int     `json:"gasUsed"`
	EthPrice float64 `json:"ethPrice"`
}
