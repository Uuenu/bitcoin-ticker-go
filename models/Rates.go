package models

//"fmt"

type Rates struct {
	Ticker   Ticker
	Currency string
	Sell     float64
}

func (r *Rates) CheckRatesParams(params map[string][]string) bool {
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

func (r *Rates) GetRates(params map[string][]string) map[string]float64 {
	r.Ticker.SetTicker()
	if len(params) > 1 {
		//return make(map[string]BitcoinPrice{params["currency"][0], r.Ticker.BitcoinPriceList[params["currency"][0]]})
		//return make([1]BitcoinPrice{params["currency"][0], r.Ticker.BitcoinPriceList[] })
		fiatCurrency := make(map[string]float64)
		fiatCurrency[params["currency"][0]] = r.Ticker.BitcoinPriceList[params["currency"][0]]
		return fiatCurrency
	} else {
		return r.Ticker.BitcoinPriceList
	}
}
