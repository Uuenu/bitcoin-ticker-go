package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gin-ticker/models"

	"github.com/gin-gonic/gin"
)

func getMethods(c *gin.Context) {

	method := c.Query("method")

	if method == "rates" {
		var ticker models.Ticker
		ticker.SetTicker()
		jsonString, err := json.Marshal(ticker.BitcoinPriceList)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, jsonString)
	}
}

func postMethods(c *gin.Context) {
	method := c.Query("method")
	if method == "convert" {
		cf, ct, val := c.Query("convert_from"), c.Query("convert_to"), c.Query("value")
		fmt.Println(cf, ct, val)
		var convert models.Convert
		valf, _ := strconv.ParseFloat(val, 32)
		convert.SetConvert(cf, ct, valf)
		convert.GetConvert()
		c.JSON(http.StatusOK, gin.H{
			"convert_from":    convert.CurrencyFrom,
			"convert_to":      convert.CurrencyTo,
			"value":           convert.Value,
			"converted_value": convert.ConvertedValue,
			"rate":            convert.Rate,
		})
	}
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.Use()
	{
		api.GET("/v1", getMethods)

		api.POST("/v1", postMethods)

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
