package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoInt(t *testing.T) {
	dataStr := "[\"auto\",1234,null]"
	data := []byte(dataStr)
	var test []AutoInt
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
