package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Girish:Girish%40123@cluster0.gefkt.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	insta := client.Database("InstagramAPI")
	userscol := insta.Collection("users")
	postcol := insta.Collection("posts")
	_, _ = userscol, postcol
	res, err := userscol.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Girish"},
		{Key: "id", Value: "123"},
		{Key: "email", Value: "sabharigirish))1@gmail.com"},
		{Key: "password", Value: "1234"},
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = res
}
