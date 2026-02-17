package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeTypeBooleanEnum(t *testing.T) {
	dataStr := "{\"type\":\"boolean\",\"values\":[{\"value\":false,\"label\":\"OFF\",\"description\":\"The device is off\"},{\"value\":true,\"label\":\"ON\",\"description\":\"The device is on\"}]}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeBoolean(t *testing.T) {
	dataStr := "{\"type\":\"boolean\"}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeNumberEnum(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"values\":[{\"value\":0,\"label\":\"idle\",\"description\":\"Studio condition is idle\"},{\"value\":1,\"label\":\"reh\",\"description\":\"Studio condition is rehearsal\"},{\"value\":2,\"label\":\"tx\",\"description\":\"Studio condition is tx\"}]}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeNumber(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"min\":{\"value\":1},\"max\":{\"value\":4294967295}}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeNumberMeasurement(t *testing.T) {
	dataStr := "{\"type\":\"number\",\"min\":{\"value\":-200,\"scale\":10},\"max\":{\"value\":1000,\"scale\":10},\"step\":{\"value\":1,\"scale\":10},\"unit\":\"C\"}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeStringEnum(t *testing.T) {
	dataStr := "{\"type\":\"string\",\"values\":[{\"value\":\"unknown\",\"label\":\"Device state is unknown\",\"description\":\"Device state is unknown. Check extension card is plugged in correctly.\"},{\"value\":\"ok\",\"label\":\"Device state is ok\",\"description\":\"Device state is ok.\"},{\"value\":\"warn\",\"label\":\"Device state is warning\",\"description\":\"Device state is warning. PSU 1 shows signs of failure.\"},{\"value\":\"fail\",\"label\":\"Device state is fail\",\"description\":\"Device state is fail. No PTP reference found.\"}]}"
	data := []byte(dataStr)
	var test Type
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

func TestTypeTypeString(t *testing.T) {
	dataStr := "{\"type\":\"string\",\"min_length\":1,\"max_length\":30,\"pattern\":\"^[a-zA-Z]+$\"}"
	data := []byte(dataStr)
	var test Type
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
