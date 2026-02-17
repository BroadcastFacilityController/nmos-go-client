package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventString(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"1ea39324-a32b-4e1d-86e9-33f9956ebc60\"},\"event_type\":\"string\",\"timing\":{\"creation_timestamp\":\"1532504241:104000200\"},\"payload\":{\"value\":\"ok\"},\"message_type\":\"state\"}"
	data := []byte(dataStr)
	var test EventString
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
