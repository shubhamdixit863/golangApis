package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shubhamdixit863/golangApis/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Register() {
	db := database.Get()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := db.Collection("users").InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		fmt.Println("ERror ocuured", err)
		log.Fatal(err)
	} else {

		fmt.Println(res)

	}
}
