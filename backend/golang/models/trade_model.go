package models

type CurrencySymbol struct {
	Symbol string `json:"symbol"`
}

type Currency struct {
	jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
	Id      int `json:"id"`
}

type Result struct {
	Id                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
}