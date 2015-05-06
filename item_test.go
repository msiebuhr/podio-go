package podio

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalNestedItems(t *testing.T) {
	data := []byte(`
		{
			"field_id": 123,
			"type": "app",
			"values": [{
				"value": {
					"item_id": 42,
					"title": "nested item"
				}
			}]
		}
	`)

	out := Field{}
	err := json.Unmarshal(data, &out)

	if err != nil {
		t.Fatalf("Unexpected error in json.Unmarshal: %s", err)
	}

	// Check it has a nested item
	if len(out.Values) != 1 {
		t.Fatalf("Expected to parse out one value, got %+v", out.Values)
	}

	// Try to decode Value as an Item and see what happens
	value, ok := out.Values[0].Value.(Item)

	if !ok || value.Id != 42 || value.Title != "nested item" {
		t.Errorf("Didn't get nested item with id=42: %+v", out.Values[0].Value)
	}
}
