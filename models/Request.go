package models

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
}

func (r *Request) ErorrJSON() gin.H {
	return gin.H{"response": "Incorrect params"}
}
