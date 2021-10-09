package main

import (
	"appointy/posts"
	"appointy/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
Note on Thread-safety: From what i read online,the net/http server automatically starts a new goroutine
for each client connection and executes request handlers in those goroutines.
So we would only need to secure the other functions that we are using,
like the dbservice and all the user&post functions

For simplicity and due to time-constraints for this project,
I chose to just make all those functions synchronised( i.e,only one thread at a time would be able to access it)
by using the Mutex provided by sync package

This is one significant area-of-future-improvement
*/

func main() {
	//POST user
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newUser users.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		newUser.EncryptPassword()
		fmt.Println(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdUser, _ := users.CreateUser(newUser)
		json.NewEncoder(w).Encode(createdUser.InsertedID)
	})

	//GET /user/{id}
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		fmt.Println(id)
		if id != "" {
			json.NewEncoder(w).Encode(users.GetUser(id))
		}
	})

	//POST /posts
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newPost posts.Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		fmt.Println(newPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdPost, _ := posts.CreatePost(newPost)
		json.NewEncoder(w).Encode(createdPost.InsertedID)
	})

	//GET /posts/{id} AND posts/users/{id}
	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/users/") {
			// GET /posts/users/{id}
			page, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil || page == 0 {
				page = 1
			}
			id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
			fmt.Println(id)
			if id != "" {
				json.NewEncoder(w).Encode(posts.GetAllPostsByUserId(id, page))
			}
		} else {
			// GET /posts/{id}
			id := strings.TrimPrefix(r.URL.Path, "/posts/")
			fmt.Println(id)
			if id != "" {
				json.NewEncoder(w).Encode(posts.GetPost(id))
			}
		}

	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
