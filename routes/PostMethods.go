package routes

import (
	"fmt"
	"gin-ticker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Convert(c *gin.Context) {

	IsAuth := c.MustGet("Auth access").(bool)
	if !IsAuth {
		var req models.Request
		c.JSON(http.StatusOK, req.AuthErorrJSON())

	} else {
		method := c.Query("method")
		params := c.Request.URL.Query()
		switch method {
		case "convert":
			var convert models.Convert
			if !convert.CheckParams(params) {
				c.JSON(http.StatusOK, convert.Request.ErorrJSON())
			} else {
				value, _ := strconv.ParseFloat(params["value"][0], 32)
				convert.SetData(params["convert_from"][0], params["convert_to"][0], value)
				convert.GetResult()
				c.JSON(http.StatusOK, convert.ResultJSON())
			}
		case "profit":
			fmt.Println("Hello")
			var profit models.Profit
			if !profit.CheckParams(params) {
				c.JSON(http.StatusOK, profit.Request.ErorrJSON())
			} else {
				profit.SetData(params)
				profit.GetResult()
				c.JSON(http.StatusOK, profit.ResultJSON())
			}
		default:
			var request models.Request
			c.JSON(http.StatusOK, request.ErorrJSON())
		}
	}

}
