package dao

import (
	"context"
	"time"

	"github.com/shubhamdixit863/golangApis/models"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shubhamdixit863/golangApis/database"
)

func Register(user *models.User) (*mongo.InsertOneResult, error) {
	db := database.Get()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	return db.Collection("users").InsertOne(ctx, user)

}

func Login(user *models.User) {

}
