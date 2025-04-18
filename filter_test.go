package greenleaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	got := Filter(
		Eq("field", "value"),
		Ne("field1", "value1"),
		Gt("field2", "value2"),
		Gte("field3", "value3"),
		Lt("field4", "value4"),
		Lte("field5", "value5"),
		Exists("field6", true),
		In("field7", []int{1, 2, 3}),
		Nin("field8", []string{"foo", "bar"}),
		GtLt("field9", "a", "b"),
		GteLt("field10", "a", "b"),
		GtLte("field11", "a", "b"),
		GteLte("field12", "a", "b"),
	)

	want := FilterDocument{
		"field":   M{"$eq": "value"},
		"field1":  M{"$ne": "value1"},
		"field2":  M{"$gt": "value2"},
		"field3":  M{"$gte": "value3"},
		"field4":  M{"$lt": "value4"},
		"field5":  M{"$lte": "value5"},
		"field6":  M{"$exists": true},
		"field7":  M{"$in": []int{1, 2, 3}},
		"field8":  M{"$nin": []string{"foo", "bar"}},
		"field9":  M{"$gt": "a", "$lt": "b"},
		"field10": M{"$gte": "a", "$lt": "b"},
		"field11": M{"$gt": "a", "$lte": "b"},
		"field12": M{"$gte": "a", "$lte": "b"},
	}

	assert.Equal(t, want, got)
}

func TestFilterSameFields(t *testing.T) {
	got := Filter(
		Eq("field", "value"),
		Eq("field", 10),
	)

	want := FilterDocument{
		"field": M{"$eq": 10},
	}

	assert.Equal(t, want, got)
}
