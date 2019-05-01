package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"deffish-server/src/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type EssayRepository struct {
	collection *mongo.Collection
}

var _ essay.IRepository = (*EssayRepository)(nil)

type Essay struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Title    string           `bson:"title"`
	Text string `bson:"text"`
	Topic string `bson:"topicId"`
	Comments []string `bson:"comments"`
}

func (repo EssayRepository) Comment(essayId aggregates.Id, comment aggregates.Comment) error {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)

	objId, err := primitive.ObjectIDFromHex(essayId.Value)
	if err != nil { return err }

	res, err := repo.collection.UpdateOne(ctx, bson.M{"_id": objId},
		bson.M{"$push": bson.M{"comments": comment.Text.Value}})
	log.Print(res.MatchedCount)
	return err
}

func (repo EssayRepository) FilterByTopic(topicId aggregates.Id) ([]aggregates.Essay, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.collection.Find(ctx, bson.M{
		"topicId": topicId.Value,
	})
	if err != nil { return nil, err }
	return fromCursorToEssays(cursor)
}

func (repo EssayRepository) Random(amount int) ([]aggregates.Essay, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	cursor, err := repo.collection.Aggregate(ctx,  bson.D{
		{"", bson.M{ "$sample": bson.M{"size": amount} }},
	})
	if err != nil { return nil, err }
	return fromCursorToEssays(cursor)
}

func (repo EssayRepository) Insert(essay aggregates.Essay) (aggregates.Id, error) {
	id, err := insert(repo.collection, toMongoEssay(essay))
	return id,  err
}


func (repo EssayRepository) Find() ([]aggregates.Essay, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	cursor, err := repo.collection.Find(ctx, nil)
	if err != nil { return nil, err }
	return fromCursorToEssays(cursor)
}

func (repo EssayRepository) Id(id aggregates.Id) (aggregates.Essay, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	objId, err := primitive.ObjectIDFromHex(id.Value)
	if err != nil { return aggregates.Essay{}, err }

	res := repo.collection.FindOne(ctx,
		bson.M{"_id": objId},
	)
	var mongoEssay Essay
	err = res.Decode(&mongoEssay)
	if err != nil { return aggregates.Essay{}, err }
	return fromMongoToEssay(mongoEssay), nil
}

func fromCursorToEssays(cursor *mongo.Cursor) ([]aggregates.Essay, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.Essay
	for cursor.Next(ctx) {
		var doc Essay
		err := cursor.Decode(&doc)
		if err != nil { return nil, err }
		items = append(items, fromMongoToEssay(doc))
	}
	return items, nil
}

func fromMongoToEssay(doc Essay) aggregates.Essay {
	var contents []aggregates.Comment
	for _, element := range doc.Comments {
		contents = append(contents, aggregates.Comment{
			aggregates.Text{
				element,
			},
		})
	}

	return aggregates.Essay{
		Id: aggregates.Id {
			Value: doc.Id.Hex(),
		},
		Title: aggregates.Title{
			Value: doc.Title,
		},
		Text: aggregates.Text{
			Value: doc.Text,
		},
		Topic: aggregates.Id{Value: doc.Topic},
		Comments: contents,
	}
}

func toMongoEssay(essay aggregates.Essay) Essay {
	return Essay{
		Title: essay.Title.Value,
		Text: essay.Text.Value,
		Topic: essay.Topic.Value,
		Comments: helpers.CommentArrToStringArray(essay.Comments),
	}
}

func (repo EssayRepository) drop() {
	_ = repo.collection.Drop(context.Background())
}
