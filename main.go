package main

import (
	"deffish-server/src/data"
	"deffish-server/src/presentation"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("could no parse %s to port", os.Args[1])
		panic(err)
	}
	log.Printf("Main started")
	presentation.NewHandler(data.NewMongoQuestionRepository(
		"mongodb://192.168.0.11:27017",
		"deffishtest",
		"questions"), port)
}
