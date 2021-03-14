package rssData 

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	"os"
	"fmt"
	"log"
)

// GetRssData godoc
// @Summary Get mongodb data and send it back
// @Description send mongodb data
// @Accept  json
// @Produce  json
// @Success 200 {object} FeedBody
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /rss-data [get]
func GetRssData(c *fiber.Ctx) error {
	client, err := GetMongoClient()
	err = client.Ping(context.TODO(), nil)
	// todo return error
    if err != nil {
        log.Fatal(err)
    }
	collection := client.Database("heroku_49s52xjc").Collection("feeds")
	var results []FeedData
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	if cursor == nil {
		return nil
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	  }
	return c.JSON(FeedBody{Data: results})
}

func GetMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")
	db_uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.avxoy.mongodb.net/heroku_49s52xjc?retryWrites=true&w=majority", username, password)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_uri))
	defer cancel()
	return client, err
}

type FeedData struct {
	Title string `json:"title" bson:"title"`
	Url string `json:"url" bson:"url"`
	Channel string `json:"channel" bson:"channel"`
}

type FeedBody struct {
	Data[] FeedData `json:"data"`
}
