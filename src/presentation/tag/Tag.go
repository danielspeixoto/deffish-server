package tag

import "deffish-server/src/aggregates"

type Tag struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func fromRequestToTag(t Tag) aggregates.Tag {
	return aggregates.Tag{
		Id: aggregates.Id{
			t.Id,
		},
		Name: t.Name,
	}
}

func fromTagToJson(t aggregates.Tag) Tag {
	return Tag{
		Id: t.Id.Value,
		Name: t.Name,
	}
}