package is04v1_3

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryAPISubscriptionPost(t *testing.T) {
	dataStr := "{\"max_update_rate_ms\":100,\"resource_path\":\"/nodes\",\"params\":{\"label\":\"host1\"},\"persist\":false,\"secure\":false}"
	data := []byte(dataStr)
	var test QueryAPISubscriptionPost
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
