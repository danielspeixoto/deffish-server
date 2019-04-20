package mongo

import "testing"

var testRepo = NewRepository(
"mongodb://localhost:27017",
"deffishtest")
var topicRepo = testRepo.Topics
var questionRepo = testRepo.Questions
var essayRepo = testRepo.Essays
var tagRepo = testRepo.Tags


func TestMain(m *testing.M) {
	dropAll()
	m.Run()
	dropAll()
}

func dropAll() {
	questionRepo.drop()
	topicRepo.drop()
	essayRepo.drop()
	tagRepo.drop()
}