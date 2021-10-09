package posts

import (
	"appointy/dbservice"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var getPostLock sync.Mutex

func GetPost(id string) Post {
	getPostLock.Lock()
	defer getPostLock.Unlock()
	var post Post
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	postCollection.FindOne(context.TODO(), filter).Decode(&post)
	return post
}
