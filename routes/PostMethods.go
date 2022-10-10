package routes

import (
	"gin-ticker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Convert(c *gin.Context) {
	method := c.Query("method")

	switch method {
	case "convert":
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
	default:
		var request models.Request
		c.JSON(http.StatusOK, request.ErorrJSON())
	}

}
