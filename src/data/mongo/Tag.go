package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type TagRepository struct {
	collection *mongo.Collection
}

func (repo TagRepository) Insert(tag aggregates.Tag) (aggregates.Id, error) {
	id, err := insert(repo.collection, toMongoTag(tag))
	return id,  err
}

func (repo TagRepository) GetByName(name string) (aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	res := repo.collection.FindOne(ctx,
		bson.M{"name": name},
	)

	var mongoTag Tag
	err := res.Decode(&mongoTag)
	if err != nil { return aggregates.Tag{}, err }
	return fromMongoToTag(mongoTag), nil
}

func (repo TagRepository) SuggestionsBySubStr(name string) ([]aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	res, err := repo.collection.Find(ctx,
		bson.M{"name": name},
	)
	if err != nil {
		return []aggregates.Tag{}, err
	}
	return fromCursorToTags(res)
}

var _ tag.IRepository = (*TagRepository)(nil)

type Tag struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string
}

func fromCursorToTags(cursor mongo.Cursor) ([]aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.Tag
	for cursor.Next(ctx) {
		var doc Tag
		err := cursor.Decode(&doc)
		if err != nil { return nil, err }
		items = append(items, fromMongoToTag(doc))
	}
	return items, nil
}

func toMongoTag(tag aggregates.Tag) Tag {
	return Tag {
		Name: tag.Name,
	}
}

func fromMongoToTag(doc Tag) aggregates.Tag {
	return aggregates.Tag{
		Id: aggregates.Id {
			Value: doc.Id.Hex(),
		},
		Name: doc.Name,
	}
}

func (repo TagRepository) drop() {
	_ = repo.collection.Drop(context.Background())
}
