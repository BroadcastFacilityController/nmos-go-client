package is05v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulkResponse(t *testing.T) {
	dataStr := "[{\"id\":\"0a174530-e3cf-11e6-bf01-fe55135034f3\",\"code\":200},{\"id\":\"7a3ebebe-0405-11e7-93ae-92361f002671\",\"code\":200},{\"id\":\"7a3ec21a-0405-11e7-93ae-92361f002671\",\"code\":200},{\"id\":\"7a3ec3b4-0405-11e7-93ae-92361f002671\",\"code\":400,\"error\":\"Un-recognised parameter 'frc_enabled'\",\"debug\":\"<stack trace>\"}]"
	data := []byte(dataStr)
	var test []BulkResponse
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
