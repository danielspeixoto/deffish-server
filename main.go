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
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		log.Fatal("$PORT must be set")
	}
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		fmt.Printf("could no parse %s to port", os.Args[1])
		panic(err)
	}
	mongoUri := os.Getenv("MONGO_CONNECTION")
	log.Printf("Main started")
	presentation.NewHandler(data.NewMongoQuestionRepository(
		mongoUri,
		"deffish",
		"questions"), port)
}
