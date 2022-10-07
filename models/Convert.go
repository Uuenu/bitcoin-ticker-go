package models

type Convert struct {
	CurrencyFrom   string
	CurrencyTo     string
	Value          float64
	ConvertedValue float64
	Rate           float64
}

func (c *Convert) SetConvert(currency_from, currency_to string, value float64) {
	c.CurrencyFrom = currency_from
	c.CurrencyTo = currency_to
	c.Value = value
}

func (c *Convert) GetConvert() {
	var ticker Ticker
	ticker.SetTicker()

	switch c.CurrencyFrom {
	case "BTC":
		{
			c.ConvertedValue = c.Value * ticker.BitcoinPriceList[c.CurrencyTo].price
			c.Rate = 1
		}

	default:
		{
			c.ConvertedValue = c.Value / ticker.BitcoinPriceList[c.CurrencyFrom].price
			c.Rate = 1
		}
	}

}
