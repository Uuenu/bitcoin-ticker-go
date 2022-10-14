package routes

import (
	"gin-ticker/models"
	"gin-ticker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rates(c *gin.Context) {

	IsAuth := c.MustGet("Auth access").(bool)
	if !IsAuth && c.Query("method") != "token" {
		var req models.Request
		c.JSON(http.StatusOK, req.AuthErorrJSON())

	} else {
		method := c.Query("method")

		switch method {
		case "rates":
			params := c.Request.URL.Query()
			var rates models.Rates
			if rates.CheckParams(params) {
				c.JSON(http.StatusOK, rates.ResultJSON())
			} else {
				c.JSON(http.StatusOK, rates.Request.ErorrJSON())
			}
		case "token":
			var getToken models.GetToken
			getToken.Token = utils.GnrBearerTkn()
			params := c.Request.URL.Query()
			getToken.CheckParams(params)
			c.JSON(http.StatusOK, gin.H{
				"Token": getToken.Token,
			})

		}

	}
}
