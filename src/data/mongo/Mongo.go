package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoRepository struct {
	Questions   *QuestionRepository
	Topics      *TopicRepository
	Essays      *EssayRepository
	Tags        *TagRepository
	Collections []*mongo.Collection
}

func NewRepository(
	uri string,
	database string) *MongoRepository {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Connection to mongo successfull")
	db := client.Database(database)

	var collections []*mongo.Collection
	questionsCollection := db.Collection("questions")
	collections = append(collections, questionsCollection)
	relatedVideosCollection := db.Collection("relatedVideos")
	collections = append(collections, relatedVideosCollection)
	topicCollection := db.Collection("topic")
	collections = append(collections, topicCollection)
	essayCollection := db.Collection("essay")
	collections = append(collections, essayCollection)
	tagCollection := db.Collection("tag")
	collections = append(collections, tagCollection)

	return &MongoRepository{
		Questions: &QuestionRepository{
			questionsCollection,
			relatedVideosCollection,
		},
		Topics:      &TopicRepository{topicCollection},
		Essays:      &EssayRepository{essayCollection},
		Tags:        &TagRepository{tagCollection},
		Collections: collections,
	}
}

func (mongo MongoRepository) drop() {
	for _, collection := range mongo.Collections {
		_ = collection.Drop(context.Background())
	}
}

func insert(repo *mongo.Collection, doc interface{}) (aggregates.Id, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	res, err := repo.InsertOne(
		ctx,
		doc)
	if err != nil {
		return aggregates.Id{}, err
	}
	id := aggregates.Id{
		Value: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	log.Printf("Doc with id %s inserted on %s relatedVideosCollection", id.Value, repo.Name())
	return id, nil
}
