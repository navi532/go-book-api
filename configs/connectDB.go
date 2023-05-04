package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODBURI")))

	if err != nil {
		log.Fatalln("Unable to connect to DB")
	}

	log.Println("Successfully connected to DB")

	DB = client.Database("golang")

}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
