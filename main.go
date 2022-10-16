package main

import (
	"fmt"
	"gin-ticker/routes"
	"gin-ticker/tests"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	//api.Use(auth.Middleware)
	api.Use()
	{
		api.GET("/v1", routes.Method)

		api.POST("/v1", routes.Method)

	}

	go tests.GenerateTests(100)

	router.Run(":5000") // listen and serve on 0.0.0.0:3000

	fmt.Println("Hello")

}
