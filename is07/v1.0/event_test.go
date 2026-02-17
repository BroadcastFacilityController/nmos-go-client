package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventEventBoolean(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"1ea39324-a32b-4e1d-86e9-33f9956ebc60\"},\"event_type\":\"boolean\",\"timing\":{\"creation_timestamp\":\"1532504241:104000200\"},\"payload\":{\"value\": false},\"message_type\":\"state\"}"
	data := []byte(dataStr)
	var test Event
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(&test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestEventEventNumber(t *testing.T) {
	dataStr := "{\"identity\": { \"source_id\": \"1ea39324-a32b-4e1d-86e9-33f9956ebc60\" }, \"event_type\": \"number\", \"timing\": { \"creation_timestamp\": \"1532504241:104000200\" }, \"payload\": { \"value\": 16342 }, \"message_type\": \"state\" }"
	data := []byte(dataStr)
	var test Event
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(&test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestEventEventNumberMeasurement(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"f9c7b88b-1846-43d9-9e53-c230e77d91ac\"},\"event_type\":\"number/temperature/C\",\"timing\":{\"creation_timestamp\":\"1532504241:104000200\"},\"payload\":{\"value\":201,\"scale\":10},\"message_type\":\"state\"}"
	data := []byte(dataStr)
	var test Event
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(&test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestEventEventNumberRational(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"1ea39324-a32b-4e1d-86e9-33f9956ebc60\"},\"event_type\":\"number\",\"timing\":{\"creation_timestamp\":\"1532504241:104000200\"},\"payload\":{\"value\":201,\"scale\":10},\"message_type\":\"state\"}"
	data := []byte(dataStr)
	var test Event
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(&test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}
