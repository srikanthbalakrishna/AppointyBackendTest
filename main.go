package main

import (
	"appointy/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	//POST user
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newUser users.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		fmt.Println(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(users.CreateUser(newUser).InsertedID)
	})

	//GET user/{id}
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		fmt.Println(id)
		if id != "" {
			json.NewEncoder(w).Encode(users.GetUser(id))
		}

	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
