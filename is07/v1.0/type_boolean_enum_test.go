package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeBooleanEnum(t *testing.T) {
	dataStr := "{\"type\":\"boolean\",\"values\":[{\"value\":false,\"label\":\"OFF\",\"description\":\"The device is off\"},{\"value\":true,\"label\":\"ON\",\"description\":\"The device is on\"}]}"
	data := []byte(dataStr)
	var test TypeBooleanEnum
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
