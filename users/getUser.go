package users

import (
	"appointy/dbservice"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(id string) User {
	var user User
	client, error := dbservice.GetMongoClient()
	fmt.Println(client, error)
	var userCollection = client.Database(dbservice.DB).Collection("users")
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	userCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}
