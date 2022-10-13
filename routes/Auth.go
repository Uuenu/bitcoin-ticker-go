package routes

import (
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	req := c.Request
	token := req.Header.Get("Authorization")

	// ctx := context.TODO()
	// opts := options.Client().ApplyURI("mongodb://localhost:27017")
	// client, err := mongo.Connect(ctx, opts)

	// if err != nil {
	// 	panic(err)
	// }

	//client.FindOne(context.TODO(), bson.M{"Token": token})

	if token != "Bearer cd25" {
		c.Set("Auth access", false)
	} else {
		c.Set("Auth access", true)
	}
}
