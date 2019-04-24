package question

import "deffish-server/src/aggregates"

type Channel struct {
	Title string
	Id string
}

type Thumbnails struct {
	HighResolution string
	DefaultResolution string
	MediumResolution string
}

type RelatedVideo struct {
	Id string
	Title string
	Description string
	Thumbnails
	Channel
	VideoId string
	QuestionId string
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