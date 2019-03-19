package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/domain/topic"
	"deffish-server/src/helpers"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type TopicRepository struct {
	topics *mongo.Collection
}

var _ topic.IRepository = (*TopicRepository)(nil)

type Topic struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Content []string           `bson:"content"`
	Title    string           `bson:"title"`
}

func (repo TopicRepository) Insert(topic aggregates.Topic) (aggregates.Id, error) {
	id, err := insert(repo.topics, toMongoTopic(topic))
	return id,  err
}

func (repo TopicRepository) drop() {
	_ = repo.topics.Drop(context.Background())
}

func (repo TopicRepository) Find() ([]aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.topics.Find(ctx, nil)
	if err != nil { return nil, err }
	return fromCursorToTopics(cursor)
}

func (repo TopicRepository) Id(id aggregates.Id) (aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil { return aggregates.Topic{}, err }

	res := repo.topics.FindOne(ctx,
		bson.M{"_id": objId},
	)
	var mongoTopic Topic
	err = res.Decode(&mongoTopic)
	if err != nil { return aggregates.Topic{}, err }
	return fromMongoToTopic(mongoTopic), nil
}

func fromCursorToTopics(cursor mongo.Cursor) ([]aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.Topic
	for cursor.Next(ctx) {
		var doc Topic
		err := cursor.Decode(&doc)
		if err != nil { return nil, err }
		items = append(items, fromMongoToTopic(doc))
	}
	return items, nil
}

func fromMongoToTopic(doc Topic) aggregates.Topic {
	var contents []aggregates.Text
	for _, element := range doc.Content {
		contents = append(contents, aggregates.Text{
			Value: element,
		})
	}

	return aggregates.Topic{
		Id: aggregates.Id {
			Value: doc.Id.Hex(),
		},
		Title: aggregates.Title{
			Value: doc.Title,
		},
		Content:contents,
	}
}

func toMongoTopic(topic aggregates.Topic) Topic {
	return Topic{
		Title: topic.Title.Value,
		Content: helpers.TextArrToStringArray(topic.Content),
	}
}
