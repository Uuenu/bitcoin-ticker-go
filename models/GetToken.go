package models

import "github.com/gin-gonic/gin"

type GetToken struct {
	IsGenerate bool
	Token      string
}

func (r *GetToken) CheckParams(params map[string][]string) bool {
	return len(params) == 1
}

func (r *GetToken) SetData(params map[string][]string) {}

func (r *GetToken) GetResult() {}

func (r *GetToken) ResultJSON() gin.H {
	return gin.H{
		"Token": r.Token,
	}
}
