package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Tag struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Amount int `bson:"amount"`
}

func fromCursorToTags(cursor *mongo.Cursor) ([]aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.Tag
	for cursor.Next(ctx) {
		var doc Tag
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}
		items = append(items, fromMongoToTag(doc))
	}
	return items, nil
}

func toMongoTag(tag aggregates.Tag) Tag {
	return Tag{
		Name: tag.Name,
		Amount: tag.Amount,
	}
}

func fromMongoToTag(tag Tag) aggregates.Tag {
	return aggregates.Tag{
		Id: aggregates.Id{
			Value: tag.Id.Hex(),
		},
		Name: tag.Name,
		Amount: tag.Amount,
	}
}