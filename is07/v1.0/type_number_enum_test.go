package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeNumberEnum(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"values\":[{\"value\":0,\"label\":\"idle\",\"description\":\"Studio condition is idle\"},{\"value\":1,\"label\":\"reh\",\"description\":\"Studio condition is rehearsal\"},{\"value\":2,\"label\":\"tx\",\"description\":\"Studio condition is tx\"}]}"
	data := []byte(dataStr)
	var test TypeNumberEnum
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}
