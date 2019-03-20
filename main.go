package main

import (
	"deffish-server/src/data/mongo"
	"deffish-server/src/injection/essay"
	"deffish-server/src/presentation"
	"deffish-server/src/presentation/question"
	"deffish-server/src/presentation/status"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
	MongoConnection string
	DBName string
}

func main() {
	config := setup()
	log.Printf("Main started")

	repo := mongo.NewRepository(
		config.MongoConnection,
		config.DBName,
	)

	presentation.NewRouter(
		essay.GetRouter(repo.Essays),
		status.Router{},
		question.Router{Repo: repo.Questions},
		config.Port)
}

func setup() Config {
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
	if mongoUri == "" {
		log.Fatal("$MONGO_CONNECTION must be set")
	}
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		log.Fatal("$MONGO_DB_NAME must be set")
	}
	return Config{
		Port:port,
		MongoConnection:mongoUri,
		DBName:mongoDBName,
	}
}