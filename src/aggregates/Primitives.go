package aggregates

type Id struct {
	Value string
}

type Title struct {
	Value string
}

type Content struct {
	Value string
}

type Tag struct {
	Name string
}

type PDF struct {
	Content []byte
}