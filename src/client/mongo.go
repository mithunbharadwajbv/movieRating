package client

import (
	"context"
	"log"
	"time"

	"github.com/sanjeev/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	movieCollection *mongo.Collection
	userCollection  *mongo.Collection

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	mongoPort = config.Conf.MongoPort
)

func init() {
	clientOptions := options.Client().ApplyURI(mongoPort)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	movieCollection = client.Database("test_db").Collection("Movies")
	userCollection = client.Database("test_db").Collection("Users")

}

func GetMovieCollection() *mongo.Collection {
	return movieCollection
}

func GetUserCollection() *mongo.Collection {
	return userCollection
}
