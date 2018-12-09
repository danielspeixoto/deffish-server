package main

import (
	"deffish-server/src/data"
	"deffish-server/src/presentation"
)

func main() {
	presentation.NewHandler(data.NewMongoQuestionRepository(
		"mongodb://localhost:27017",
		"deffishtest",
		"questions"), 5000)
}
