package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
	"deffish-server/src/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TopicRepository struct {
	collection *mongo.Collection
}

var _ topic.IRepository = (*TopicRepository)(nil)

type Topic struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Content []string           `bson:"content"`
	Title    string           `bson:"title"`
}

func (repo TopicRepository) Random(amount int) ([]aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	cursor, err := repo.collection.Aggregate(ctx,  bson.D{
		{"", bson.M{ "$sample": bson.M{"size": amount} }},
	})
	if err != nil { return nil, err }
	return fromCursorToTopics(cursor)
}

func (repo TopicRepository) Insert(topic aggregates.Topic) (aggregates.Id, error) {
	id, err := insert(repo.collection, toMongoTopic(topic))
	return id,  err
}

func (repo TopicRepository) drop() {
	_ = repo.collection.Drop(context.Background())
}

func (repo TopicRepository) FindAll() ([]aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.collection.Find(ctx, nil)
	if err != nil { return nil, err }
	return fromCursorToTopics(cursor)
}

func (repo TopicRepository) Id(id aggregates.Id) (aggregates.Topic, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil { return aggregates.Topic{}, err }

	res := repo.collection.FindOne(ctx,
		bson.M{"_id": objId},
	)
	var mongoTopic Topic
	err = res.Decode(&mongoTopic)
	if err != nil { return aggregates.Topic{}, err }
	return fromMongoToTopic(mongoTopic), nil
}

func fromCursorToTopics(cursor *mongo.Cursor) ([]aggregates.Topic, error) {
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
