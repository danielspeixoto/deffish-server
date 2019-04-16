package aggregates

type Question struct {
	Id
	Image
	Source string
	Variant string
	Edition int
	Number int
	Domain string
	Answer  int
	Tags []string
}
