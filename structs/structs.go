package structs

// Coinbase is a representation of the coinbase api https://api.coinbase.com/v2/exchange-rates
type Coinbase struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    struct {
			BCH string `json:"BCH"`
			BTC string `json:"BTC"`
			ETH string `json:"ETH"`
			LTC string `json:"LTC"`
		} `json:"rates"`
	} `json:"data"`
}
