package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/rts-core/versioner/api"
	accessors "github.com/rts-core/versioner/db/access/mongo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to Mongo
	credential := options.Credential{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CLIENT_URL")).SetAuth(credential))
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

	// Accessors
	applicationAccessor := accessors.NewMongoApplicationsAccessor(client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME")))

	err = api.Listen(port, applicationAccessor)
	if err != nil {
		log.Fatal(err)
	}
}
