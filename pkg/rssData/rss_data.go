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
	"github.com/FriendlyUser/go_fin_server/pkg/types"
)

// GetRssData godoc
// @Summary Get mongodb data and send it back
// @Description send mongodb data
// @Accept  json
// @Produce  json
// @Success 200 {object} types.FeedBody
// @Failure 400 {object} types.HTTPError
// @Failure 404 {object} types.HTTPError
// @Failure 500 {object} types.HTTPError
// @Router /rss-data [get]
func GetRssData(c *fiber.Ctx) error {
	client, err := GetMongoClient()
	err = client.Ping(context.TODO(), nil)
	// todo return error
    if err != nil {
        log.Fatal(err)
    }
	collection := client.Database("heroku_49s52xjc").Collection("feeds")
	var results []types.FeedData
	options := options.Find()
	// Limit by 10 documents only 
	options.SetLimit(0)
	cursor, _ := collection.Find(context.TODO(), bson.D{}, options)
	if cursor == nil {
		return nil
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	  }
	return c.JSON(types.FeedBody{Data: results})
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
