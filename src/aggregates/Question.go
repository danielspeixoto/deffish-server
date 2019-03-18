package aggregates

type Choice struct {
	Content string
}

type Question struct {
	Id
	PDF
	Answer  int
	Choices []Choice
	Tags    []Tag
}
