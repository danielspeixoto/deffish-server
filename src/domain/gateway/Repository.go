package gateway

import "deffish-server/src/domain"

type IQuestionRepository interface {
	Insert(domain.Question) (domain.Id, error)
	Random(amount int, tags []domain.Tag) ([]domain.Question, error)
	Id(id domain.Id) (domain.Question, error)
}