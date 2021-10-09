package posts

import (
	"appointy/dbservice"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Note that it is quite inefficient to first get all the documents and then paginate manually,
a better approach would be to query documents according to the page-number
i.e, skip (pageNumber-1)*CountPerPage and limit CounterPerPage

But even this approach can become slow as the number of documents increase

The best way would be have a sorting order(like sorting by ID) and then using a range and
selecting only those documents from mongoDB that correspond to the right index (index <---derived---- pageNumber)
*/
func GetAllPostsByUserId(userid string, pageNumber int) []Post {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}
