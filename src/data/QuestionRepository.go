package data

import (
	"deffish-server/src/domain"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongoQuestionRepository struct {
	*mongo.Database
}

type MongoQuestion struct {
	PDF []byte
	Answer int
	Choices []string
	Tags []string
	Edition string
	Metadata map[string]string
}

func NewMongoQuestionRepository(uri string, database string) *MongoQuestionRepository {
	client, err := mongo.NewClient(uri)
	client.Database(database)
	if err != nil {
		panic(err)
	}
	db := client.Database(database)
	return &MongoQuestionRepository {
		db,
	}
}

func (repo MongoQuestionRepository) Insert(question domain.Question) (error) {
	
}

func (repo MongoQuestionRepository) Drop() (error) {
	
}

func (repo MongoQuestionRepository) Find() ([]domain.Question, error) {
	
}

func (repo MongoQuestionRepository) Random() (domain.Question, error) {
	
}

