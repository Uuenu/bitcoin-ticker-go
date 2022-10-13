package routes

import (
	"gin-ticker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rates(c *gin.Context) {

	IsAuth := c.MustGet("Auth access").(bool)
	if !IsAuth {
		var req models.Request
		c.JSON(http.StatusOK, req.AuthErorrJSON())

	} else {
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
}
