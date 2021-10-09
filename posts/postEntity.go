package posts

import "sync"

// Assuming that each post is associated with a user
type Post struct {
	Caption   string `json:"caption"`
	Id        string `json:"id"`
	URL       string `json:"url"`
	TimeStamp int    `json:"time"`
	UserId    string `json:"userid"`
}

const CountPerPage int = 2

var lock sync.Mutex
