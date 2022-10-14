package models

import "github.com/gin-gonic/gin"

type Requester interface {
	CheckParams(params map[string]string)
	ResultJSON(c *gin.Context)
	GetResult()
	SetData()
	ErorrJSON()
	AuthErorrJSON()
	AuthStatus()
}
