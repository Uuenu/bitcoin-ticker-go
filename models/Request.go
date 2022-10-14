package models

import (
	"github.com/gin-gonic/gin"
)

type BaseRequester interface {
	ErorrJSON() gin.H
	AuthStatus() bool
	AuthErorrJSON() gin.H
}

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

func (r *Request) UndefinedMethod() gin.H {
	return gin.H{"Response": "Undefined Method"}
}
