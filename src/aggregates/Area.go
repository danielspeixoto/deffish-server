package aggregates

type Area struct {
	Id
	Name string
	Tags []Id
}