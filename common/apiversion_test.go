package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIVersion(t *testing.T) {
	dataStr := "[\"v1.0\",\"v1.1\",\"v2.0\",\"v2.12\"]"
	data := []byte(dataStr)
	var test []APIVersion
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(&test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result))
}
