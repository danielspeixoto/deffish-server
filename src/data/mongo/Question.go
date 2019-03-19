package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/helpers"
	"deffish-server/src/domain/question"
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
	Id primitive.ObjectID `bson:"_id,omitempty"`
	PDF []byte `bson:"pdf"`
	Answer int `bson:"answer"`
	Choices []string `bson:"choices"`
	Tags []string `bson:"tags"`
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
	log.Printf("question with id %s inserted", id.Value)
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

func (repo QuestionRepository) Random(amount int, tags []aggregates.Tag) ([]aggregates.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	agg := bson.D{
		{"", bson.M{
			"$match": bson.M{
				"tags": bson.M{
					"$all": helpers.TagsToStringArray(tags),
				},
			},
		}},
		{"", bson.M{ "$sample": bson.M{"size": amount} }},
	}

	if len(tags) == 0 {
		agg = bson.D{
			{"", bson.M{ "$sample": bson.M{"size": amount} }},
		}
	}

	cursor, err := repo.questions.Aggregate(ctx, agg)
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
	var choices []aggregates.Choice
	for _, element := range doc.Choices {
		choices = append(choices, aggregates.Choice{
			Content: element,
		})
	}

	var tags []aggregates.Tag
	for _, element := range doc.Tags {
		tags = append(tags, aggregates.Tag{
			Name: element,
		})
	}
	return aggregates.Question{
		Id: aggregates.Id {
			Value: doc.Id.Hex(),
		},
		PDF: aggregates.PDF{
			Content: doc.PDF,
		},
		Answer: doc.Answer,
		Tags:tags,
		Choices:choices,
	}
}

func toMongoQuestion(question aggregates.Question) Question {
	return Question{
		PDF: question.PDF.Content,
		Answer: question.Answer,
		Choices: helpers.ChoicesToStringArray(question.Choices),
		Tags: helpers.TagsToStringArray(question.Tags),
	}
}
