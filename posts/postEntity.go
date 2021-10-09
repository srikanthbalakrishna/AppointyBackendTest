package posts

type Post struct {
	Caption   string `json:"caption"`
	Id        string `json:"id"`
	URL       string `json:"url"`
	TimeStamp int    `json:"time"`
}
