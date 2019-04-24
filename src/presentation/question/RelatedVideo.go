package question

import (
	"deffish-server/src/aggregates"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Channel struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

type Thumbnails struct {
	HighResolution    string `json:"high"`
	DefaultResolution string `json:"default"`
	MediumResolution  string `json:"medium"`
}

type RelatedVideo struct {
	Id                string `json:"id"`
	Title             string             `json:"title"`
	Description       string `json:"description"`
	Thumbnails        `json:"thumbnails"`
	Channel           `json:"channel"`
	VideoId string `json:"videoId"`
	QuestionId        string `json:"questionId"`
}

func fromRelatedVideosToJsonArray(relatedVideos []aggregates.RelatedVideo) []RelatedVideo {
	jsons := make([]RelatedVideo, len(relatedVideos))
	for q := range relatedVideos {
		jsons = append(jsons, fromRelatedVideoToJson(q))
	}
	return jsons
}

func fromRelatedVideoToJson(relatedVideo aggregates.RelatedVideo) RelatedVideo {
	return RelatedVideo{
		Id:          relatedVideo.Id.Value,
		Title:       relatedVideo.Title,
		Description: relatedVideo.Description,
		Thumbnails: Thumbnails{
			HighResolution:    relatedVideo.Thumbnail.HighResolution,
			MediumResolution:  relatedVideo.Thumbnail.MediumResolution,
			DefaultResolution: relatedVideo.Thumbnail.DefaultResolution,
		},
		Channel: Channel{
			Title: relatedVideo.Channel.Title,
			Id:    relatedVideo.Channel.Id.Value,
		},
		VideoId:    relatedVideo.VideoId.Value,
		QuestionId: relatedVideo.QuestionId.Value,
	}
}