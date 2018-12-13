package data

import (
	"context"
	"deffish-server/src/domain"
	"deffish-server/src/domain/gateway"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type MongoQuestionRepository struct {
	questions *mongo.Collection
}

var _ gateway.IQuestionRepository = (*MongoQuestionRepository)(nil)

type MongoQuestion struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	PDF []byte `bson:"pdf"`
	Answer int `bson:"answer"`
	Choices []string `bson:"choices"`
	Tags []string `bson:"tags"`
}

func NewMongoQuestionRepository(
	uri string,
	database string,
	questionsCollection string) *MongoQuestionRepository {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	client, err := mongo.Connect(ctx, uri)
	if err != nil {
		panic(err)
	}
	db := client.Database(database)
	return &MongoQuestionRepository {
		db.Collection(questionsCollection),
	}
}

func (repo MongoQuestionRepository) Insert(question domain.Question) (domain.Id, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	res, err := repo.questions.InsertOne(
		ctx,
		toMongoQuestion(question))
	if err != nil { return domain.Id{}, err }
	id := domain.Id{
		Value: res.InsertedID.(primitive.ObjectID).String(),
	}
	return id,  nil
}

func (repo MongoQuestionRepository) Drop() error {
	return repo.questions.Drop(context.Background())
}

func (repo MongoQuestionRepository) Find() ([]domain.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.questions.Find(ctx, nil)
	if err != nil { return nil, err }
	return fromCursorToQuestions(cursor)
}

func (repo MongoQuestionRepository) Random(amount int, tags []domain.Tag) ([]domain.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err := repo.questions.Aggregate(ctx,
		bson.D{
			{"", bson.M{
				"$match": bson.M{
					"tags": bson.M{
						"$all": tagsToStringArray(tags),
					},
				},
			}},
			{"", bson.M{ "$sample": bson.M{"size": amount} }},
		},
	)
	if err != nil { return nil, err }
	return fromCursorToQuestions(cursor)
}

func fromCursorToQuestions(cursor mongo.Cursor) ([]domain.Question, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cursor.Close(ctx)
	var items []domain.Question
	for cursor.Next(ctx) {
		var doc MongoQuestion
		err := cursor.Decode(&doc)
		if err != nil { return nil, err }
		items = append(items, fromMongoToQuestion(doc))
	}
	return items, nil
}

func fromMongoToQuestion(doc MongoQuestion) domain.Question {
	var choices []domain.Choice
	for _, element := range doc.Choices {
		choices = append(choices, domain.Choice{
			Content: element,
		})
	}

	var tags []domain.Tag
	for _, element := range doc.Tags {
		tags = append(tags, domain.Tag{
			Name: element,
		})
	}
	return domain.Question{
		Id: domain.Id {
			Value: doc.Id.String(),
		},
		PDF: domain.PDF{
			Content: doc.PDF,
		},
		Answer: doc.Answer,
		Tags:tags,
		Choices:choices,
	}
}

func toMongoQuestion(question domain.Question) MongoQuestion{
	return MongoQuestion{
		PDF: question.PDF.Content,
		Answer: question.Answer,
		Choices: choicesToStringArray(question.Choices),
		Tags: tagsToStringArray(question.Tags),
	}
}

func tagsToStringArray(tags []domain.Tag) []string {
	var tagsStr []string
	for _, element := range tags {
		tagsStr = append(tagsStr, element.Name)
	}
	return tagsStr
}

func choicesToStringArray(choices []domain.Choice) []string {
	var choicesStr []string
	for _, element := range choices {
		choicesStr = append(choicesStr, element.Content)
	}
	return choicesStr
}

