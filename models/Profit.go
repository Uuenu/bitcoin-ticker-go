package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Profit struct {
	Request         Request
	Ticker          Ticker
	Сurrency_from   string
	Currency_to     string
	Value_from      float64
	PriceInPurchase float64
	Value_crypto    float64
	ProfitValue     float64
}

func (p *Profit) CheckParams(params map[string][]string) bool {
	if len(params) != 5 {
		return false
	}
	if params["currency_from"] == nil || params["currency_to"] == nil || params["value_from"] == nil || params["price_purchase"] == nil {
		return false
	}

	if params["currency_to"][0] != "BTC" {
		return false
	}

	fiatCheck := false

	p.Ticker.SetTicker()
	for key := range p.Ticker.BitcoinPriceList {
		if key == params["currency_from"][0] {
			fiatCheck = true
		}
	}

	return fiatCheck

}

func (p *Profit) SetData(params map[string][]string) {
	p.Сurrency_from = params["currency_from"][0]
	p.Currency_to = params["currency_to"][0]
	p.Value_from, _ = strconv.ParseFloat(params["value_from"][0], 32)
	p.PriceInPurchase, _ = strconv.ParseFloat(params["price_purchase"][0], 32)
	p.Value_crypto = p.Value_from / p.PriceInPurchase

}

func (p *Profit) GetResult() {
	p.Ticker.SetTicker()
	p.ProfitValue = p.Ticker.BitcoinPriceList[p.Сurrency_from]*p.Value_crypto - p.Value_from
}

func (p *Profit) ResultJSON() gin.H {
	return gin.H{
		"Currency from":                               p.Сurrency_from,
		"Currency to":                                 p.Currency_to,
		"Value of selling fiat currency":              p.Value_from,
		"Value of buying crypto currency":             p.Value_crypto,
		"Price of crypto currency in purchase moment": p.PriceInPurchase,
		"Total Profin":                                p.ProfitValue,
	}
}
