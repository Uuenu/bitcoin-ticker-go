package routes

import (
	"gin-ticker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, r models.Requester) {

	params := c.Request.URL.Query()
	if !r.CheckParams(params) {
		var baseReq models.Request
		c.JSON(http.StatusOK, baseReq.ErorrJSON())
	} else {
		r.SetData(params)
		r.GetResult()
		c.JSON(http.StatusOK, r.ResultJSON())
	}

}

func Method(c *gin.Context) {

	IsAuth := c.MustGet("Auth access").(bool)
	method := c.Query("method")
	if !IsAuth && method != "token" {
		var req models.Request
		c.JSON(http.StatusOK, req.AuthErorrJSON())

	} else {
		switch method {
		case "profit":
			var profit models.Profit
			Response(c, &profit)
		case "convert":
			var convert models.Convert
			Response(c, &convert)
		case "token":
			var token models.GetToken
			Response(c, &token)
		case "rates":
			var rates models.Rates
			Response(c, &rates)
		default:
			var request models.Request
			c.JSON(http.StatusOK, request.UndefinedMethod())
		}

	}
}
