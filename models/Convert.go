package models

import "github.com/gin-gonic/gin"

type Convert struct {
	Request        Request
	Ticker         Ticker
	CurrencyFrom   string
	CurrencyTo     string
	Value          float64
	ConvertedValue float64
	Rate           float64
}

func (c *Convert) CheckConvertParams(params map[string][]string) bool {

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

func (c *Convert) SetConvert(currency_from, currency_to string, value float64) {
	c.CurrencyFrom = currency_from
	c.CurrencyTo = currency_to
	c.Value = value
}

func (c *Convert) GetConvert() {
	c.Ticker.SetTicker()

	switch c.CurrencyFrom {
	case "BTC":
		{
			c.ConvertedValue = c.Value * c.Ticker.BitcoinPriceList[c.CurrencyTo]
			c.Rate = 1
		}

	default:
		{
			c.ConvertedValue = c.Value / c.Ticker.BitcoinPriceList[c.CurrencyFrom]
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
