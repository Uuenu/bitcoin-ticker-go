package models

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Convert struct {
	Request        Request
	Ticker         Ticker
	CurrencyFrom   string
	CurrencyTo     string
	Value          float64
	ConvertedValue float64
	Rate           float64
}

func (c *Convert) CheckParams(params map[string][]string) bool {

	fmt.Print("params len", len(params))

	if len(params) != 4 {
		return false
	}

	if params["convert_from"] == nil || params["convert_to"] == nil || params["value"] == nil {
		return false
	}

	if params["convert_from"][0] != "BTC" && params["convert_to"][0] != "BTC" {
		return false
	}

	fiatCheck := false

	c.Ticker.SetTicker()
	for key := range c.Ticker.BitcoinPriceList {
		if key == params["convert_from"][0] || key == params["convert_to"][0] {
			fiatCheck = true
		}
	}

	return fiatCheck

}

func (c *Convert) SetData(params map[string][]string) {

	c.CurrencyFrom = params["convert_from"][0]
	c.CurrencyTo = params["convert_to"][0]
	c.Value, _ = strconv.ParseFloat(params["value"][0], 64)

}

func (c *Convert) GetResult() {
	c.Ticker.SetTicker()

	switch c.CurrencyFrom {
	case "BTC":
		{
			c.ConvertedValue = c.Value*c.Ticker.BitcoinPriceList[c.CurrencyTo] - FEE*(c.Value*c.Ticker.BitcoinPriceList[c.CurrencyTo])
			c.Rate = 1
		}

	default:
		{
			c.ConvertedValue = c.Value/c.Ticker.BitcoinPriceList[c.CurrencyFrom] - FEE*(c.Value*c.Ticker.BitcoinPriceList[c.CurrencyTo])
			c.Rate = 1
		}
	}

}

func (c *Convert) ResultJSON() gin.H {
	return gin.H{
		"convert from":    c.CurrencyFrom,
		"convert to":      c.CurrencyTo,
		"value":           c.Value,
		"converted value": c.ConvertedValue,
		"rate":            c.Rate,
	}
}
