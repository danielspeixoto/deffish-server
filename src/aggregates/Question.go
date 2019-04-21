package aggregates

type Question struct {
	Id
	PDF
	Source string
	Variant string
	Edition int
	Number int
	Domain string
	Answer  int
	Tags []string
}
