package users

import (
	"appointy/dbservice"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(newUser User) (*mongo.InsertOneResult, error) {
	client, _ := dbservice.GetMongoClient()
	var userCollection = client.Database(dbservice.DB).Collection(dbservice.USERS_COLLECTION)
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult, err
}
