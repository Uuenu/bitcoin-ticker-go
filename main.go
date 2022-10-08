package main

import (
	"net/http"
	"strconv"

	"gin-ticker/models"

	"github.com/gin-gonic/gin"
)

func getMethods(c *gin.Context) {

	method := c.Query("method")

	if method == "rates" {
		params := c.Request.URL.Query()
		var rates models.Rates
		if rates.CheckRatesParams(params) {
			c.JSON(http.StatusOK, rates.GetRates(params))
		} else {
			c.JSON(http.StatusOK, rates.Request.ErorrJSON())
		}
	}
}

func postMethods(c *gin.Context) {
	method := c.Query("method")
	if method == "convert" {
		params := c.Request.URL.Query()
		var convert models.Convert
		if !convert.CheckConvertParams(params) {
			c.JSON(http.StatusOK, convert.Request.ErorrJSON())
		} else {
			value, _ := strconv.ParseFloat(params["value"][0], 32)
			convert.SetConvert(params["convert_from"][0], params["convert_to"][0], value)
			convert.GetConvert()
			c.JSON(http.StatusOK, convert.ResultJSON())
		}

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

	router.Run(":5000") // listen and serve on 0.0.0.0:3000
}
