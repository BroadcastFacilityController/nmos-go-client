package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandSubscription(t *testing.T) {
	dataStr := "{\"command\":\"subscription\",\"sources\":[\"772116e0-b4ba-43b1-9ffc-70287c17cb9e\",\"674e32cb-84b5-475e-b7db-7821530c4375\"]}"
	data := []byte(dataStr)
	var test CommandSubscription
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

func TestCommandSubscriptionUnsubscribe(t *testing.T) {
	dataStr := "{\"command\":\"subscription\",\"sources\":[]}"
	data := []byte(dataStr)
	var test CommandSubscription
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
