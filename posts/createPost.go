package posts

import (
	"appointy/dbservice"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(newPost Post) *mongo.InsertOneResult {
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	insertResult, err := postCollection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}
