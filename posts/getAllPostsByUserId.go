package posts

import (
	"appointy/dbservice"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllPostsByUserId(userid string, pageNumber int) []Post {
	var posts []Post
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	filter := bson.D{primitive.E{Key: "userid", Value: userid}}

	cursor, err := postCollection.Find(context.TODO(), filter)
	fmt.Println(cursor, err)
	cursor.All(context.TODO(), &posts)

	return paginate(posts, ((pageNumber - 1) * CountPerPage), CountPerPage)
}

func paginate(x []Post, skip int, size int) []Post {
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}
