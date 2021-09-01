package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var clientOptions *options.ClientOptions
	clientOptions = options.Client().ApplyURI("mongodb://docker.for.mac.localhost:27001")
	clientOptions.SetMaxPoolSize(50)
	m, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err", err)
	}
	err = m.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err", err)
	}
	res := make(map[string]interface{})
	test := m.Database("job").Collection("test")
	_, err = test.InsertOne(context.TODO(), bson.M{"foo": "bar"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("inserted")
	opt := options.FindOne().SetSort(bson.M{"_id": -1})
	err = test.FindOne(context.TODO(), bson.M{}, opt).Decode(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	//block it!
	select {}
}
