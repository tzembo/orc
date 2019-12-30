package orc

import (
	"testing"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func TestMapEntryMarshalBSON(t *testing.T) {
	t.Run("invalid MapEntry struct, non-string keys", func(t *testing.T) {
		entry := MapEntry{Key: 1, Value: "apple"}
		_, err := entry.MarshalBSON()
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("valid MapEntry struct", func(t *testing.T) {
		entry := MapEntry{Key: "b", Value: "banana"}
		byt, err := entry.MarshalBSON()
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		expected := "{\"b\": \"banana\"}"
		actual := bsoncore.Document(byt).String()
		if expected != actual {
			t.Errorf("expected %s, but got %s", expected, actual)
		}
	})
}

func TestUnionValueMarshalBSONValue(t *testing.T) {
	val := UnionValue{Tag: 0, Value: "apple"}
	btype, byt, err := val.MarshalBSONValue()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	expected := "\"apple\""
	actual := bsoncore.Value{Type: btype, Data: byt}.String()
	if expected != actual {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}
