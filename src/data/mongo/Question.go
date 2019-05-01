package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Question struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	View        []byte             `bson:"view"`
	Source      string             `bson:"source"`
	Variant     string             `bson:"variant"`
	Edition     int                `bson:"edition"`
	Number      int                `bson:"number"`
	Answer      int                `bson:"answer"`
	Domain      string             `bson:"domain"`
	Tags        []string           `bson:"tags"`
	ItemCode    string             `bson:"itemCode"`
	ReferenceId string             `bson:"referenceId"`
	Stage int `bson:"stage"`
}

func fromCursorToQuestions(cursor *mongo.Cursor) ([]aggregates.Question, error) {
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
		View:   aggregates.View{
			doc.View,
		},
		Source:      doc.Source,
		Variant:     doc.Variant,
		Edition:     doc.Edition,
		Number:      doc.Number,
		Domain:      doc.Domain,
		Answer:      doc.Answer,
		Tags:        doc.Tags,
		ItemCode:    doc.ItemCode,
		ReferenceId: doc.ReferenceId,
		Stage: doc.Stage,
	}
}

func toMongoQuestion(question aggregates.Question) Question {
	return Question{
		View:        question.View.Contents,
		Source:      question.Source,
		Variant:     question.Variant,
		Edition:     question.Edition,
		Number:      question.Number,
		Answer:      question.Answer,
		Domain:      question.Domain,
		Tags:        question.Tags,
		ItemCode:    question.ItemCode,
		ReferenceId: question.ReferenceId,
		Stage: question.Stage,
	}
}
