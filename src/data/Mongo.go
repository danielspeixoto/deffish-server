package data

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

type MongoRepository struct {
	Questions *mongo.Collection
}

func NewMongoRepository(
	uri string,
	database string,
	questionsCollection string) *MongoRepository {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	client, err := mongo.Connect(ctx, uri)
	if err != nil {
		panic(err)
	}
	log.Printf("Connection to mongo successfull")
	db := client.Database(database)
	return &MongoRepository{
		db.Collection(questionsCollection),
	}
}
