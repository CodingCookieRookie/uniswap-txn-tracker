package model

type SwapEvent struct {
	TxnHash      string `json:"txnHash"`
	SqrtPriceX96 string `json:"sqrtPriceX96"`
	From         string `json:"from"`
	To           string `json:"to"`
}

type SwapEventWithToken struct {
	SqrtPriceX96 string `json:"sqrtPriceX96"`
	From         string `json:"from"`
	To           string `json:"to"`
	TokenName    string `json:"tokenName"`
	TokenSymbol  string `json:"tokenSymbol"`
	TokenDecimal string `json:"tokenDecimal"`
}

type SwapPrice struct {
	SwapEventWithToken *SwapEventWithToken `json:"swap_event"`
	Price_Text         string              `json:"price,omitempty"`
}

type SwapResponse struct {
	SwapPrices []*SwapPrice `json:"swap_prices"`
}
