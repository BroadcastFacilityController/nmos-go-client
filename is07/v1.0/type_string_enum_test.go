package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeStringEnum(t *testing.T) {
	dataStr := "{\"type\":\"string\",\"values\":[{\"value\":\"unknown\",\"label\":\"Device state is unknown\",\"description\":\"Device state is unknown. Check extension card is plugged in correctly.\"},{\"value\":\"ok\",\"label\":\"Device state is ok\",\"description\":\"Device state is ok.\"},{\"value\":\"warn\",\"label\":\"Device state is warning\",\"description\":\"Device state is warning. PSU 1 shows signs of failure.\"},{\"value\":\"fail\",\"label\":\"Device state is fail\",\"description\":\"Device state is fail. No PTP reference found.\"}]}"
	data := []byte(dataStr)
	var test TypeStringEnum
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
