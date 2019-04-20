package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

type Repository struct {
	Questions *QuestionRepository
	Topics *TopicRepository
	Essays *EssayRepository
	Tags *TagRepository
}

func NewRepository(
	uri string,
	database string) *Repository {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	client, err := mongo.Connect(ctx, uri)
	if err != nil {
		panic(err)
	}
	log.Printf("Connection to mongo successfull")
	db := client.Database(database)
	return &Repository{
		NewQuestionRepository(db.Collection("questions")),
		&TopicRepository{db.Collection("topic")},
		&EssayRepository{db.Collection("essay")},
		&TagRepository{db.Collection("tag")},
	}
}

func insert(repo *mongo.Collection, doc interface{}) (aggregates.Id, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	res, err := repo.InsertOne(
		ctx,
		doc)
	if err != nil { return aggregates.Id{}, err }
	id := aggregates.Id{
		Value: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	log.Printf("Doc with id %s inserted on %s collection", id.Value, repo.Name())
	return id,  nil
}