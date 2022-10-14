package main

import (
	"gin-ticker/auth"
	"gin-ticker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.Use(auth.Middleware)
	api.Use()
	{
		api.GET("/v1", routes.Rates)

		api.POST("/v1", routes.Convert)

	}

	router.Run(":5000") // listen and serve on 0.0.0.0:3000
}
