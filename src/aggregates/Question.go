package aggregates

type Question struct {
	Id
	View
	Source      string
	Variant     string
	Edition     int
	Number      int
	Domain      string
	Answer      int
	Tags        []string
	ItemCode    string
	ReferenceId string
}



