package movieRepo

import (
	"context"
	"fmt"

	"github.com/sanjeev/src/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Movie struct {
	Name            string `json:"name" bson:"_id,omitempty"`
	Discription     string `json:"discription"`
	Rating          float64
	Comments        []string
	NoOfPeopleRated int64
}

var movieCollection *mongo.Collection = client.GetMovieCollection()

func DropMovieCollection() {
	movieCollection.Drop(context.Background())
}

func GetDummyMovie(name string) Movie {
	return Movie{Name: name, Rating: 5, Comments: []string{"tullare"}, NoOfPeopleRated: 1}
}

func AddMovie(movie Movie) error {
	ctx := context.Background()
	_, err := movieCollection.InsertOne(ctx, movie)
	if err != nil {
		fmt.Println("Error while inserting data to Mongo")
		return err
	}
	return nil
}

// UpdateMovie updates existing movie struct with new one
func UpdateMovie(movie Movie) error {
	ctx := context.Background()
	filter := bson.D{{"_id", movie.Name}}
	update := bson.M{
		"$set": movie,
	}
	_, err := movieCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Error while updating movie data")
	}
	return nil
}

// GetMovieByName takes name and returns movie struct
func GetMovieByName(name string) (movie *Movie, err error) {

	ctx := context.Background()

	cur := movieCollection.FindOne(ctx, bson.D{{"_id", name}})
	cur.Decode(&movie)
	if movie == nil {
		return nil, fmt.Errorf("No movie object found")
	}
	return movie, nil
}
