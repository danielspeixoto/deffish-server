package aggregates

type Exam struct {
	Id
	Name string
	Components []Id
}
