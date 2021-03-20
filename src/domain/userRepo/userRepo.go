package userRepo

import (
	"context"
	"fmt"

	"github.com/sanjeev/src/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name   string   `json:"name" bson:"_id,omitempty"`
	Type   string   `json:"type"`
	Rating []Rating `json:"rating"`
}

type Rating struct {
	MovieName string `json:"movieName"`
	Rating    int64  `json:"rating"`
	Comment   string `json:"comments"`
}

var userCollection *mongo.Collection = client.GetUserCollection()

func DropUserCollection() {
	userCollection.Drop(context.Background())
}

func GetDummyUser(name string) User {
	return User{Name: name, Type: "admin", Rating: []Rating{Rating{MovieName: "kichha", Rating: 5, Comment: "tullare ballare"}}}
}

// AddUser : adds user to mongo
func AddUser(user User) error {
	ctx := context.Background()
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error while inserting data to Mongo")
		return err
	}
	return nil
}

// UpdateMovie updates existing movie struct with new one
func UpdateUser(user User) error {
	ctx := context.Background()
	filter := bson.D{{"_id", user.Name}}
	update := bson.M{
		"$set": user,
	}
	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error while updating user data")
	}
	return nil
}

// GetUserByName takes name and returns user struct
func GetUserByName(name string) (user *User, err error) {

	ctx := context.Background()

	cur := userCollection.FindOne(ctx, bson.D{{"_id", name}})
	cur.Decode(&user)
	if user == nil {
		return nil, fmt.Errorf("No user found with name " + name)
	}
	return user, nil
}
