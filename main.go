package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
	"versioner/api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CLIENT_URL")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Setup API
	port, err := strconv.Atoi(os.Getenv("HOST_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	err = api.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}
