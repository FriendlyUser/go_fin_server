package types

type HTTPError struct {
	status  string
	message string
}

type Account struct {
	Id string
}

type FeedData struct {
	Title string `json:"title" bson:"title"`
	Url string `json:"url" bson:"url"`
	Channel string `json:"channel" bson:"channel"`
}

type FeedBody struct {
	Data[] FeedData `json:"data"`
}

// from golang serverless function
type Message struct {
	Data [][]string `json:"data"`
	Columns [8]string `json:"columns"`
	Index []string `json:"index"`
}