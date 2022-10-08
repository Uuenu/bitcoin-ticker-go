package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	TICKER_URL = "https://blockchain.info/ticker"
	FEE        = 0.02
)

var ()

type BitcoinPrice struct {
	//currency string  `json:"price"`
	price float64 `json:"price"`
}

type Ticker struct {
	Data             map[string]map[string]interface{}
	BitcoinPriceList map[string]float64
}

func (ticker *Ticker) SetTicker() {
	resp, err := http.Get(TICKER_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ticker.Data) // get ticker
	if err != nil {
		panic(err)
	}

	ticker.BitcoinPriceList = make(map[string]float64)
	for key, value := range ticker.Data {
		sell := fmt.Sprintf("%v", value["sell"])
		sellf, _ := strconv.ParseFloat(sell, 32)
		//btcPrice := BitcoinPrice{price: sellf}
		btcPrice := sellf
		ticker.BitcoinPriceList[string(key)] = btcPrice
		//fmt.Println(string(key), " ", ticker.BitcoinPriceList[string(key)])
	}

}
