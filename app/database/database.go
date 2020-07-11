package database

import (
	"calenderfy-grpc-service/app/config"
	"calenderfy-grpc-service/app/logger"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var out *log.Logger = logger.GetLogger("database.go")

type Database struct {
	*mongo.Collection
}

func NewDatabase(c config.Configs) (Database, error) {
	out.Println("initializing mongo client", c.MongoHost)

	options := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s", c.MongoUser, c.MongoPassword, c.MongoHost))

	client, err := mongo.NewClient(options)
	if err != nil {
		out.Println("error while initializing client", err)
		return Database{}, err
	}

	out.Println("attempting to connect to mongodb at host", c.MongoHost)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		out.Println("error while connectiing to mongo db at host", c.MongoHost)
		return Database{}, err
	}

	out.Println("successfully connected to mongo at host", c.MongoHost)

	return Database{client.Database(c.MongoDatabase).Collection(c.MongoCollection)}, nil
}
