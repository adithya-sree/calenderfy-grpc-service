package dao

import (
	"calendarfy-grpc-service/app/config"
	"calendarfy-grpc-service/app/logger"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var out = logger.GetLogger("database.go")

type Dao struct {
	mongo  *mongo.Database
	config config.Configs
}

func NewDao(c config.Configs) (Dao, error) {
	out.Println("initializing mongo client", c.MongoHost)

	opt := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", c.MongoUser, c.MongoPassword, c.MongoHost, c.MongoDatabase))

	client, err := mongo.NewClient(opt)
	if err != nil {
		out.Println("error while initializing client", err)
		return Dao{}, err
	}

	out.Println("attempting to connect to mongodb at host", c.MongoHost)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		out.Println("error while connecting to mongo db at host", c.MongoHost)
		return Dao{}, err
	}

	out.Println("successfully connected to mongo at host", c.MongoHost)

	return Dao{mongo: client.Database(c.MongoDatabase), config: c}, nil
}

func (d *Dao) FindByEmail(email string) (Profile, error) {
	out.Println("attempting to find profile for email", email)

	var p Profile

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := d.mongo.Collection(d.config.MongoCollection).FindOne(ctx, bson.M{"email": email}).Decode(&p)
	start := time.Now()
	if err != nil {
		out.Printf("error while querying mongo %s, unable to find profile %s, failed in %dms", err, p, time.Since(start))
		return Profile{}, err
	}

	out.Printf("successfully found profile %s for email %s in %dms", p, email, time.Since(start))

	return p, nil
}

func (d *Dao) InsertProfile(p Profile) error {
	out.Println("attempting to insert profile", p)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	insertResult, err := d.mongo.Collection(d.config.MongoCollection).InsertOne(ctx, p)
	start := time.Now()
	if err != nil {
		out.Printf("error while inserting to mongo %s, unable to find profile %s, failed in %dms", err, p, time.Since(start))
		return err
	}

	out.Printf("succesfully inserted profile %s, result %s in %dms", p, insertResult.InsertedID, time.Since(start))
	return nil
}

func (d *Dao) UpdateProfile(old Profile, new Profile) error {
	out.Printf("attempting to update profile %s", old)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	start := time.Now()
	updateResult, err := d.mongo.Collection(d.config.MongoCollection).UpdateOne(ctx, bson.M{"email": old.Email}, bson.D{
		{"$set", bson.D{
			{"email", new.Email},
			{"pushToken", new.PushToken},
			{"events", new.Events},
		}},
	})

	if err != nil {
		out.Printf("error while updating profile %s, error %v, failed in %sms", old, err, time.Since(start))
		return err
	}

	out.Printf("successfully updated profile %s, matched %v documents and updated %v documents in %sms", new, updateResult.MatchedCount, updateResult.ModifiedCount, time.Since(start))
	return nil
}
