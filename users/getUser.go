package users

import (
	"appointy/dbservice"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var getUserLock sync.Mutex

func GetUser(id string) User {
	getUserLock.Lock()
	defer getUserLock.Unlock()
	var user User
	client, _ := dbservice.GetMongoClient()
	var userCollection = client.Database(dbservice.DB).Collection(dbservice.USERS_COLLECTION)
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	userCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}
