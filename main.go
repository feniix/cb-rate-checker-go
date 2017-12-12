package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/structs"
	. "github.com/feniix/cb-rate-checker-go/structs"
	"github.com/shopspring/decimal"
)

func main() {
	url := "https://api.coinbase.com/v2/exchange-rates"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	cb := Coinbase{}
	jsonErr := json.Unmarshal(body, &cb)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	rates := structs.Map(&cb.Data.Rates)
	one := decimal.NewFromFloat(1)
	eth, _ := decimal.NewFromString(rates["ETH"].(string))

	fmt.Printf("%v\n", one.DivRound(eth, 2))
}
