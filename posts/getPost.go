package posts

import (
	"appointy/dbservice"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPost(id string) Post {
	var post Post
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	postCollection.FindOne(context.TODO(), filter).Decode(&post)
	return post
}
