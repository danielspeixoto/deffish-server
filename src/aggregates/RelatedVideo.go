package aggregates

type Channel struct {
	Title string
	Id
}

type Thumbnail struct {
	HighResolution string
	DefaultResolution string
	MediumResolution string
}

type RelatedVideo struct {
	Id Id
	Title string
	Description string
	Thumbnail
	Channel
	VideoId Id
	QuestionId Id
}