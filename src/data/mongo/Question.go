package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

type QuestionRepository struct {
	questions *mongo.Collection
}

var _ question.IRepository = (*QuestionRepository)(nil)

type Question struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	PDF     []byte             `bson:"pdf"`
	Source  string             `bson:"source"`
	Variant string             `bson:"variant"`
	Edition int                `bson:"edition"`
	Number  int                `bson:"number"`
	Answer  int                `bson:"answer"`
	Domain  string             `bson:"domain"`
	Tags    []string           `bson:"tags"`
}

func NewQuestionRepository(questions *mongo.Collection) *QuestionRepository {
	return &QuestionRepository{questions}
}

func (repo QuestionRepository) Insert(question aggregates.Question) (aggregates.Id, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	res, err := repo.questions.InsertOne(
		ctx,
		toMongoQuestion(question))
	if err != nil { return aggregates.Id{}, err }
	id := aggregates.Id{
		Value: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	log.Printf("testQuestion with id %s inserted", id.Value)
	return id,  nil
}

func (repo QuestionRepository) drop() {
	_ = repo.questions.Drop(context.Background())
}

func (repo QuestionRepository) Find() ([]aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.questions.Find(ctx, nil)
	if err != nil { return nil, err }
	return fromCursorToQuestions(cursor)
}

func (repo QuestionRepository) Id(id aggregates.Id) (aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil { return aggregates.Question{}, err }

	res := repo.questions.FindOne(ctx,
		bson.M{"_id": objId},
	)
	var mongoQuestion Question
	err = res.Decode(&mongoQuestion)
	if err != nil { return aggregates.Question{}, err }
	return fromMongoToQuestion(mongoQuestion), nil
}

func (repo QuestionRepository) random(field string, value []string, amount int) (mongo.Cursor, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	agg := bson.D{
		{"", bson.M{
			"$match": bson.M{
				field: bson.M{
					"$all": value,
				},
			},
		}},
		{"", bson.M{ "$sample": bson.M{"size": amount} }},
	}


	if len(value) == 0 {
		agg = bson.D{
			{"", bson.M{"$sample": bson.M{"size": amount}}},
		}
	}



	return repo.questions.Aggregate(ctx, agg)
}

func (repo QuestionRepository) RandomByDomain(amount int, domain string) ([]aggregates.Question, error) {
	cursor, err := repo.random("domain", []string{domain}, amount)
	if err != nil { return nil, err }
	return fromCursorToQuestions(cursor)
}

func (repo QuestionRepository) RandomByTags(amount int, tags []string) ([]aggregates.Question, error) {
	cursor, err := repo.random("tags", tags, amount)
	if err != nil { return nil, err }
	return fromCursorToQuestions(cursor)
}

func fromCursorToQuestions(cursor mongo.Cursor) ([]aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.Question
	for cursor.Next(ctx) {
		var doc Question
		err := cursor.Decode(&doc)
		if err != nil { return nil, err }
		items = append(items, fromMongoToQuestion(doc))
	}
	return items, nil
}

func fromMongoToQuestion(doc Question) aggregates.Question {
	return aggregates.Question{
		Id: aggregates.Id{
			Value: doc.Id.Hex(),
		},
		PDF:   aggregates.PDF{
			doc.PDF,
		},
		Source:  doc.Source,
		Variant: doc.Variant,
		Edition: doc.Edition,
		Number:  doc.Number,
		Domain:  doc.Domain,
		Answer:  doc.Answer,
		Tags:    doc.Tags,
	}
}

func toMongoQuestion(question aggregates.Question) Question {
	return Question{
		PDF:     question.PDF.Contents,
		Source:  question.Source,
		Variant: question.Variant,
		Edition: question.Edition,
		Number:  question.Number,
		Answer:  question.Answer,
		Domain:  question.Domain,
		Tags:    question.Tags,
	}
}