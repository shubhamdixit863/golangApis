package dao

import (
	"context"
	"time"

	"github.com/shubhamdixit863/golangApis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shubhamdixit863/golangApis/database"
)

func CreateListing(listing *models.Listings) (*mongo.InsertOneResult, error) {
	db := database.Get()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	return db.Collection("listings").InsertOne(ctx, listing)

}

func GetListing() (*mongo.Cursor, error) {
	db := database.Get()
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)

	return db.Collection("listings").Find(ctx, bson.D{})
}
