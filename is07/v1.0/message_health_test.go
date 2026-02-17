package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageHealth(t *testing.T) {
	dataStr := "{\"timing\":{\"origin_timestamp\":\"1441974485:123000000\",\"creation_timestamp\":\"1441974485:234000000\"},\"message_type\":\"health\"}"
	data := []byte(dataStr)
	var test MessageHealth
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
