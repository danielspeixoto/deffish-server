package topic

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(aggregates.Topic) (aggregates.Id, error)
	Id(id aggregates.Id) (aggregates.Topic, error)
}
