package tag

import (
	"deffish-server/src/aggregates"
	"testing"
)

func TestConversion(t *testing.T) {
	t.Run("From request to Name", func(t *testing.T) {
		aggTag := fromRequestToTag(Tag{
			Id: "1",
			Name: "a",
		})

		if aggTag.Id.Value != "1" {
			t.Errorf("Id is wrong")
		}

		if aggTag.Name != "a" {
			t.Errorf("Name is wrong")
		}
	})

	t.Run("From Name to Json", func(t *testing.T) {
		jsonTag := fromTagToJson(aggregates.Tag{
			Id: aggregates.Id{
				"1",
			},
			Name:"a",
		})

		if jsonTag.Id != "1" {
			t.Errorf("Id is Wrong")
		}
		if jsonTag.Name != "a" {
			t.Errorf("Name is wrong")
		}
	})
}