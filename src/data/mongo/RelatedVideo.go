package mongo

import (
	"context"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/relatedVideo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type RelatedVideoRepository struct {
	collection *mongo.Collection
}

var _ relatedVideo.IRepository = (*RelatedVideoRepository)(nil)

type Channel struct {
	Title string `bson:"title"`
	Id    string `bson:"id"`
}

type Thumbnails struct {
	High    string `bson:"high"`
	Default string `bson:"default"`
	Medium  string `bson:"medium"`
}

type RelatedVideo struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	QuestionId        primitive.ObjectID `bson:"questionId,omitempty"`
	RetrievalPosition int                `bson:"retrievalPosition"`
	Title             string             `bson:"title"`
	Channel           `bson:"channel"`
	Thumbnails        `bson:"thumbnails"`
	Description       string `bson:"description"`
	VideoId string `bson:"videoId"`
	//PublishedAt int64 `bson"/"`
}

func (repo RelatedVideoRepository) FilterByQuestion(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error) {
	panic("implement me")
}

func (repo RelatedVideoRepository) drop() {
	_ = repo.collection.Drop(context.Background())
}

func fromCursorToRelatedVideos(cursor mongo.Cursor) ([]aggregates.RelatedVideo, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	defer cursor.Close(ctx)
	var items []aggregates.RelatedVideo
	for cursor.Next(ctx) {
		var doc RelatedVideo
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}
		items = append(items, fromMongoToRelatedVideo(doc))
	}
	return items, nil
}

func fromMongoToRelatedVideo(doc RelatedVideo) aggregates.RelatedVideo {
	return aggregates.RelatedVideo{
		Id: aggregates.Id{
			Value: doc.Id.Hex(),
		},
		Title:       doc.Title,
		Description: doc.Description,
		Thumbnail: aggregates.Thumbnail{
			HighResolution:    doc.Thumbnails.High,
			MediumResolution:  doc.Thumbnails.Medium,
			DefaultResolution: doc.Thumbnails.Default,
		},
		VideoId: aggregates.Id{
			doc.VideoId,
		},
		Channel: aggregates.Channel{
			Title: doc.Channel.Title,
			Id: aggregates.Id{
				doc.Channel.Id,
			},
		},
	}
}

func toMongoRelatedVideo(relatedVideo aggregates.RelatedVideo) RelatedVideo {
	objId, err := primitive.ObjectIDFromHex(relatedVideo.QuestionId.Value)
	if err != nil {
		panic(err)
	}
	return RelatedVideo{
		QuestionId: objId,
		Title: relatedVideo.Title,
		Channel: Channel{
			Title: relatedVideo.Channel.Title,
			Id:  relatedVideo.Channel.Id.Value,
		},
		Thumbnails: Thumbnails{
			High: relatedVideo.Thumbnail.HighResolution,
			Medium: relatedVideo.Thumbnail.MediumResolution,
			Default: relatedVideo.Thumbnail.DefaultResolution,
		},
		Description: relatedVideo.Description,
		VideoId: relatedVideo.VideoId.Value,
	}
}
