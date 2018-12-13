package main

import (
	"deffish-server/src/data"
	"deffish-server/src/presentation"
	"log"
)

func main() {
	log.Printf("Main started")
	presentation.NewHandler(data.NewMongoQuestionRepository(
		"mongodb://192.168.0.11:27017",
		"deffishtest",
		"questions"), 5000)
}
