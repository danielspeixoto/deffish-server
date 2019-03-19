package essay

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(essay aggregates.Essay) (aggregates.Id, error)
	Id(id aggregates.Id) (aggregates.Essay, error)
	FilterByTopic(aggregates.Id) ([]aggregates.Essay, error)
	Comment(essayId aggregates.Id, comment aggregates.Comment) error
}
