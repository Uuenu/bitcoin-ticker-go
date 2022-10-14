package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GnrBearerTkn() string {

	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}
	collection := client.Database("BitcoinTicker").Collection("Bearer")

	token := "random string"

	result, err := collection.InsertOne(context.TODO(), bson.D{{Key: "token", Value: "Bearer " + token}})
	fmt.Println(result)
	if err != nil {
		return ""
	} else {
		return "Bearer cd25"
	}
}
