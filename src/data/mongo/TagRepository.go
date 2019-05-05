package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type TagRepository struct {
	collection *mongo.Collection
}

func (repo TagRepository) IncrementCount(name string) error {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"name": name},
		bson.M{
			"$inc": bson.M{
				"amount": 1,
			},
		},
	)
	return err
}

func (repo TagRepository) Insert(tag aggregates.Tag) (aggregates.Id, error) {
	id, err := insert(repo.collection, toMongoTag(tag))
	return id, err
}

func (repo TagRepository) GetByName(name string) (aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	res := repo.collection.FindOne(ctx,
		bson.M{"name": name},
	)

	var mongoTag Tag
	err := res.Decode(&mongoTag)
	if err != nil {
		return aggregates.Tag{}, err
	}
	return fromMongoToTag(mongoTag), nil
}

func (repo TagRepository) SuggestionsBySubStr(name string, minCount int) ([]aggregates.Tag, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	log.Print(name)
	agg := []bson.M{
		{
			"$match": bson.M{
				"name": primitive.Regex{Pattern: name, Options: ""},
				"amount": bson.M{
					"$gt": minCount,
				},
			},
		},
	}

	res, err := repo.collection.Aggregate(ctx, agg)
	if err != nil {
		return []aggregates.Tag{}, err
	}
	return fromCursorToTags(res)
}

var _ tag.IRepository = (*TagRepository)(nil)
