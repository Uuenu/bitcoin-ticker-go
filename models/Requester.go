package models

import "github.com/gin-gonic/gin"

type Requester interface {
	CheckParams(params map[string][]string) bool
	GetResult()
	SetData(params map[string][]string)
	ResultJSON() gin.H
}
