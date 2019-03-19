package aggregates


type Essay struct {
	Id
	Title
	Text
	Topic Id
	Comments []Comment
}
