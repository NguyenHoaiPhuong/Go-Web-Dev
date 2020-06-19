package models

// Response : struct
type Response struct {
	Data   *Data   `json:"data"`
	Status *Status `json:"status"`
}

// Data : struct
type Data struct {
	BTC  *Coin `json:"BTC"`
	ETH  *Coin `json:"ETH"`
	TOKO *Coin `json:"TOKO"`
	USDT *Coin `json:"USDT"`
}

// Status : struct
type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}

// Coin : struct
type Coin struct {
	ID     uint32 `json:"id,omitempty"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  *Quote `json:"quote"`
}

// Quote : struct
type Quote struct {
	USD *USD `json:"USD"`
}

// USD : struct
type USD struct {
	Price          float32 `json:"price"`
	PecentChange1h float32 `json:"percent_change_1h"`
}
