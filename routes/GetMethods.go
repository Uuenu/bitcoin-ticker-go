package routes

import (
	"gin-ticker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rates(c *gin.Context) {

	method := c.Query("method")
	//fmt.Println("Auth: ", c.Query("Authorization"))
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
