package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"github.com/reactivex/rxgo/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type QuestionRepository struct {
	questionsCollection     *mongo.Collection
	relatedVideosCollection *mongo.Collection
}

func (repo QuestionRepository) Add(id aggregates.Id, tag string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil {
		return err
	}
	res, err := repo.questionsCollection.UpdateOne(ctx, bson.M{"_id": objId},
		bson.M{"$addToSet": bson.M{"tags": tag}})
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New(errors.ErrorCode(123), "already has tag")
	}
	return err
}

func (repo QuestionRepository) GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil {
		return []aggregates.RelatedVideo{}, err
	}

	agg := []bson.M{
		{
			"$match": bson.M{
				"questionId": objId,
			},
		},
		{
			"$sort": bson.M{
				"retrievalPosition": 1,
			},
		},
		{"$skip": start},
		{"$limit": count},
	}

	cursor, err := repo.relatedVideosCollection.Aggregate(ctx, agg)
	if err != nil {
		return []aggregates.RelatedVideo{}, err
	}
	return fromCursorToRelatedVideos(cursor)
}

func (repo QuestionRepository) Insert(question aggregates.Question) (aggregates.Id, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := repo.questionsCollection.InsertOne(
		ctx,
		toMongoQuestion(question))
	if err != nil {
		return aggregates.Id{}, err
	}
	id := aggregates.Id{
		Value: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	log.Printf("testQuestion with id %s inserted", id.Value)
	return id, nil
}

func (repo QuestionRepository) Find() ([]aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := repo.questionsCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	return fromCursorToQuestions(cursor)
}

func (repo QuestionRepository) Id(id aggregates.Id) (aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil {
		return aggregates.Question{}, err
	}

	res := repo.questionsCollection.FindOne(ctx,
		bson.M{"_id": objId},
	)
	var mongoQuestion Question
	err = res.Decode(&mongoQuestion)
	if err != nil {
		return aggregates.Question{}, err
	}
	return fromMongoToQuestion(mongoQuestion), nil
}

func (repo QuestionRepository) random(field string, value []string, amount int) (*mongo.Cursor, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	agg := []bson.M{
		{
			"$match": bson.M{
				field: bson.M{
					"$all": value,
				},
			},
		},
		{
			"$sample": bson.M{"size": amount},
		},
	}

	if len(value) == 0 {
		agg = []bson.M{
			{
				"$sample": bson.M{"size": amount},
			},
		}
	}
	return repo.questionsCollection.Aggregate(ctx, agg)
}

func (repo QuestionRepository) RandomByDomain(amount int, domain string) ([]aggregates.Question, error) {
	cursor, err := repo.random("domain", []string{domain}, amount)
	if err != nil {
		return nil, err
	}
	return fromCursorToQuestions(cursor)
}

func (repo QuestionRepository) RandomByTags(amount int, tags []string) ([]aggregates.Question, error) {
	cursor, err := repo.random("tags", tags, amount)
	if err != nil {
		return nil, err
	}
	return fromCursorToQuestions(cursor)
}

var _ question.IRepository = (*QuestionRepository)(nil)
