package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database *mongo.Database
)

// Connect with database
func Connect(db string, url string) *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//"mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	// Calling the  disconnect function using defer such that it gets called at the end
	/*defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	*/

	if err != nil {
		log.Fatal(err)
		return nil

	} else {
		fmt.Println("Connected With Db")
		database = client.Database(db)
		return client

	}

}

func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client.Disconnect(ctx)
}

func Get() *mongo.Database {
	return database
}
