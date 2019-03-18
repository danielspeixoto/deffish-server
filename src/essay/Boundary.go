package essay

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(essay aggregates.Essay) (aggregates.Id, error)
	Id(id aggregates.Id) (aggregates.Essay, error)
	Topic(aggregates.Id) ([]aggregates.Essay, error)
	Comment(comment aggregates.Comment) error
}
