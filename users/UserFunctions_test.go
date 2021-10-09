package users

import (
	"appointy/dbservice"
	"context"
	"testing"
)

var testUser User = User{"srikanth", "test-user-1", "New York", 21, "myPassword"}

func TestGetUser(t *testing.T) {
	initializeTestData()
	defer deleteTestData()
	response := GetUser(testUser.Id)
	//TO-DO: compare the struct (i.e, all attributes should match and not just id)
	if response.Id != testUser.Id {
		t.Error("Attributes do not match")
	}
}

func TestCreateUser(t *testing.T) {
	response, err := CreateUser(testUser)
	defer deleteTestData()
	if err != nil || response == nil {
		t.Error("Did not create user")
	}
}

func TestEncryptPassword(t *testing.T) {
	plain_password := testUser.Password
	testUser.EncryptPassword()
	if plain_password == testUser.Password {
		t.Error("Did not encrypt password")
	}
}

func TestCreateHash(t *testing.T) {
	hashed := createHash("hello")
	if hashed == "hello" {
		t.Error("Incorrectly hashed passphrase")
	}
}

//utility function
func initializeTestData() {
	//skipping password encryption for data initialization
	client, _ := dbservice.GetMongoClient()
	var userCollection = client.Database(dbservice.DB).Collection(dbservice.USERS_COLLECTION)
	userCollection.InsertOne(context.TODO(), testUser)

}

//utility function
func deleteTestData() {
	client, _ := dbservice.GetMongoClient()
	var userCollection = client.Database(dbservice.DB).Collection(dbservice.USERS_COLLECTION)
	userCollection.DeleteOne(context.TODO(), testUser.Id)
}
