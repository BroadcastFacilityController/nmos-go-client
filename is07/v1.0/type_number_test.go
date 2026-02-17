package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeNumber(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"min\":{\"value\":1},\"max\":{\"value\":4294967295}}"
	data := []byte(dataStr)
	var test TypeNumber
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

func TestTypeNumberMeasurement(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"min\":{\"value\":-200,\"scale\":10},\"max\":{\"value\":1000,\"scale\":10},\"step\":{\"value\":1,\"scale\":10},\"unit\":\"C\"}"
	data := []byte(dataStr)
	var test TypeNumber
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
