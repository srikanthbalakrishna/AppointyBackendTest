package users

import (
	"appointy/dbservice"
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var createUserLock sync.Mutex

func CreateUser(newUser User) (*mongo.InsertOneResult, error) {
	createUserLock.Lock()
	defer createUserLock.Unlock()
	client, _ := dbservice.GetMongoClient()
	var userCollection = client.Database(dbservice.DB).Collection(dbservice.USERS_COLLECTION)
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult, err
}
