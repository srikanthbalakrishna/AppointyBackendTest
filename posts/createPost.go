package posts

import (
	"appointy/dbservice"
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var createPostLock sync.Mutex

func CreatePost(newPost Post) (*mongo.InsertOneResult, error) {
	createPostLock.Lock()
	defer createPostLock.Unlock()
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	insertResult, err := postCollection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult, err
}
