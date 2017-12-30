package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	flag "github.com/spf13/pflag"
	"os"
)

type coinbase struct {
	Data struct {
		Rates    map[string]string `json:"rates"`
	} `json:"data"`
}

func main() {
	const url = "https://api.coinbase.com/v2/exchange-rates?currency=%s"

	var currency, baseCurrency string
	flag.StringVarP(&currency, "currency", "c", "", "Ticker symbol of the currency (ETH, LTC, BTC)")
	flag.StringVarP(&baseCurrency, "base", "b", "USD", "Base currency used to express the price")
	flag.Parse()

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	if currency == "" {
		flag.Usage()
		os.Exit(1)
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(url, currency), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	cb := coinbase{}
	err = json.Unmarshal(body, &cb)
	if err != nil {
		log.Fatal(err)
	}

	exchanged := cb.Data.Rates[baseCurrency]

	switch currency {
	case
		"ETH",
		"BTC",
		"LTC",
		"BCH":
		fmt.Printf("%s\n", exchanged)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
