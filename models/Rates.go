package models

import "github.com/gin-gonic/gin"

type Rates struct {
	Request  Request
	Ticker   Ticker
	Currency string
	Sell     float64
	Result   map[string]float64
}

func (r *Rates) CheckParams(params map[string][]string) bool {
	if len(params) > 2 {
		return false
	}

	if params["currency"] != nil {
		fiatCheck := false

		r.Ticker.SetTicker()
		for key := range r.Ticker.BitcoinPriceList {
			if key == params["currency"][0] {
				fiatCheck = true
			}
		}
		return fiatCheck
	}
	return true
}

func (r *Rates) SetData(params map[string][]string) {
	if len(params) > 1 {
		r.Currency = params["currency"][0]
	}
}

func (r *Rates) GetResult() {
	r.Ticker.SetTicker()
	if r.Currency != "" {
		fiatCurrency := make(map[string]float64)
		fiatCurrency[r.Currency] = r.Ticker.BitcoinPriceList[r.Currency]
		r.Result = fiatCurrency
	} else {
		r.Result = r.Ticker.BitcoinPriceList
	}
}

func (r *Rates) ResultJSON() gin.H {
	json := make(map[string]interface{})
	for key, value := range r.Result {
		json[key] = value
	}
	return json
}
