package users

import (
	"appointy/dbservice"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(newUser User) *mongo.InsertOneResult {
	client, error := dbservice.GetMongoClient()
	fmt.Println(client, error)
	var userCollection = client.Database(dbservice.DB).Collection("users")
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}
