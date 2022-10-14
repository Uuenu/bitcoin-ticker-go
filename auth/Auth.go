package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Middleware(c *gin.Context) {

	req := c.Request
	token := req.Header.Get("Authorization")

	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}
	collection := client.Database("BitcoinTicker").Collection("Bearer")

	var result bson.M
	collection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&result)

	//fmt.Println("Tokens ", token, " - ", result["token"])

	if result != nil {
		c.Set("Auth access", true)
	} else {
		c.Set("Auth access", false)
	}
}
