package aggregates

type Question struct {
	Id
	PDF
	Source      string
	Variant     string
	Edition     int
	Number      int
	Domain      string
	Answer      int
	Tags        []string
	ReferenceId string
}

type QuestionData struct {
	PDF
	Source      string
	Variant     string
	Edition     int
	Number      int
	Domain      string
	Answer      int
	Tags        []string
	ReferenceId string
	CreatedAt	int64
}

type QuestionForm struct {
	PDF
	Source      string
	Variant     string
	Edition     int
	Number      int
	Domain      string
	Answer      int
	Tags        []string
	ReferenceId string
}



