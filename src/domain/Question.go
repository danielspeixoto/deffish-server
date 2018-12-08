package domain

type Choice struct {
	Content string
}

type Tag struct {
	Name string
}

type PDF struct {
	Content []byte
}

type Id struct {
	Value string
}

type Question struct {
	Id
	PDF
	Answer  int
	Choices []Choice
	Tags    []Tag
}
