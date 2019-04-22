package mongo

import (
	"deffish-server/src/aggregates"
	"reflect"
	"sort"
	"testing"
)

func TestTagManyItems(t *testing.T) {
	tag := aggregates.Tag{
		Name: "mytag",
	}
	tag1 := aggregates.Tag{
		Name: "mytag2",
	}
	tag2 := aggregates.Tag{
		Name: "mybear",
	}
	tag3 := aggregates.Tag{
		Name: "not",
	}
	tag4 := aggregates.Tag{
		Name: "never",
	}

	t.Run("Insert Items", func(t *testing.T) {
		_, err := testRepo.Tags.Insert(tag1)
		if err != nil {
			t.Fatal(err)
		}
		id, err := testRepo.Tags.Insert(tag)
		if err != nil {
			t.Fatal(err)
		}
		_, err = testRepo.Tags.Insert(tag2)
		if err != nil {
			t.Fatal(err)
		}
		_, err = testRepo.Tags.Insert(tag3)
		if err != nil {
			t.Fatal(err)
		}
		_, err = testRepo.Tags.Insert(tag4)
		if err != nil {
			t.Fatal(err)
		}
		t.Run("Retrieve using Name", func(t *testing.T) {
			result, err := tagRepo.GetByName("mytag")
			if err != nil { t.Fatal(err) }
			if result.Id != id {
				t.Errorf("Differs")
			}
		})

		t.Run("Suggestion using beginning substr", func(t *testing.T) {
			result, err := tagRepo.SuggestionsBySubStr("my")
			if err != nil { t.Fatal(err) }
			expected := []string{"mytag", "mytag2", "mybear"}
			sort.Strings(expected)
			retrieved := []string{}
			for _, element := range result {
				retrieved = append(retrieved, element.Name)
			}
			sort.Strings(retrieved)
			if reflect.DeepEqual(expected, retrieved) {
				t.Errorf("Differs")
			}

		})

		t.Run("Suggestion using middle substr", func(t *testing.T) {
			result, err := tagRepo.SuggestionsBySubStr("tag")
			if err != nil { t.Fatal(err) }
			expected := []string{"mytag", "mytag2"}
			sort.Strings(expected)
			retrieved := []string{}
			for _, element := range result {
				retrieved = append(retrieved, element.Name)
			}
			sort.Strings(retrieved)
			if reflect.DeepEqual(expected, retrieved) {
				t.Errorf("Differs")
			}

		})
	})
}