package domain

type Choice struct {
	content string
}

type Tag struct {
	Name string
}

type Edition struct {
	Number int
}

type PDF struct {
	content []byte
}

type Question struct {
	PDF
	Answer  int
	Choices []Choice
	Tags    []Tag
	Edition
}
