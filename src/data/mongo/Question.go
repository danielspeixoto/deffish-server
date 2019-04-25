package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

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
	ReferenceId string `bson:"referenceId"`
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
		ReferenceId: doc.ReferenceId,
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
		ReferenceId: question.ReferenceId,
	}
}

