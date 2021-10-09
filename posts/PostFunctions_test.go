package posts

import (
	"appointy/dbservice"
	"context"
	"testing"
)

var testPost Post = Post{"TEST Post!", "POS-TEST", "www.google.com", 1633744129, "USR-TEST"}

var testArray = []Post{
	{"Caption 1", "POS-TEST-1", "www.google.com", 1633744129, "USR-TEST"},
	{"Caption 2", "POS-TEST-2", "www.google.com", 1633744129, "USR-TEST"},
	{"Caption 3", "POS-TEST-3", "www.google.com", 1633744129, "USR-TEST"},
	{"Caption 4", "POS-TEST-4", "www.google.com", 1633744129, "USR-TEST"},
	{"Caption 5", "POS-TEST-5", "www.google.com", 1633744129, "USR-TEST"},
}

func TestGetPost(t *testing.T) {
	initializeTestData()
	defer deleteTestData()
	response := GetPost(testPost.Id)
	//TO-DO: compare the struct (i.e, all attributes should match and not just id)
	if response.Id != testPost.Id {
		t.Error("Attributes do not match")
	}
}

func TestCreatePost(t *testing.T) {
	response, err := CreatePost(testPost)
	defer deleteTestData()
	if err != nil || response == nil {
		t.Error("Did not create post")
	}
}

func TestGetAllPostsByUserId(t *testing.T) {

	for i := 0; i < CountPerPage+1; i++ {
		initializeTestData()
	}
	defer deleteTestData()

	response := GetAllPostsByUserId(testPost.UserId, 1)
	if response == nil || len(response) != CountPerPage {
		t.Error("Incorrect pagination")
	}

}

func TestPaginate(t *testing.T) {
	arr := paginate(testArray, 1, 3)
	if arr[0].Id != "POS-TEST-2" || len(arr) != 3 {
		t.Error("Incorrect pagination")
	}

}

//utility function
func initializeTestData() {
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	postCollection.InsertOne(context.TODO(), testPost)

}

//utility function
func deleteTestData() {
	client, _ := dbservice.GetMongoClient()
	var postCollection = client.Database(dbservice.DB).Collection(dbservice.POSTS_COLLECTION)
	postCollection.DeleteOne(context.TODO(), testPost.Id)
}
