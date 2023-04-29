package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/red512/go-naru/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	v = utils.Config
)

func Client() *mongo.Client {
	clientOptions := options.Client().ApplyURI(v.GetString("db.mongoURI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
