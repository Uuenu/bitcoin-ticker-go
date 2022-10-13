package models

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
}

func (r *Request) ErorrJSON() gin.H {
	return gin.H{"response": "Incorrect params"}
}

func (r *Request) AuthStatus(authStatus bool) bool {
	return authStatus
}

func (r *Request) AuthErorrJSON() gin.H {
	return gin.H{"response": "Authorization Erorr"}
}
