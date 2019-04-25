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
	testRepo.drop()
	m.Run()
	testRepo.drop()
}