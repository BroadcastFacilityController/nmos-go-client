package is04v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistrationAPIHealthGet(t *testing.T) {
	dataStr := "{\"health\":\"1441974485\"}"
	data := []byte(dataStr)
	var test RegistrationHealth
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

func TestRegistrationAPIHealthPost(t *testing.T) {
	dataStr := "{\"health\":\"1441974485\"}"
	data := []byte(dataStr)
	var test RegistrationHealth
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
