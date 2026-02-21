package is04v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryAPISubscriptionIDGet(t *testing.T) {
	dataStr := "{\"max_update_rate_ms\":100,\"resource_path\":\"/nodes\",\"params\":{},\"persist\":false,\"secure\":false,\"ws_href\":\"ws://172.29.80.52:8870/ws/?uid=7c903667-7113-4a8f-8865-1c63f9070a9e\",\"id\":\"7c903667-7113-4a8f-8865-1c63f9070a9e\"}"
	data := []byte(dataStr)
	var test QueryAPISubscription
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

func TestQueryAPISubscriptionsGet(t *testing.T) {
	dataStr := "[{\"max_update_rate_ms\":100,\"resource_path\":\"/receivers\",\"params\":{},\"persist\":false,\"secure\":false,\"ws_href\":\"ws://172.29.80.52:8870/ws/?uid=6a52dbd5-a737-4c4e-823f-909ade8f8bf4\",\"id\":\"6a52dbd5-a737-4c4e-823f-909ade8f8bf4\"},{\"max_update_rate_ms\":100,\"resource_path\":\"/nodes\",\"params\":{},\"persist\":false,\"secure\":false,\"ws_href\":\"ws://172.29.80.52:8870/ws/?uid=7c903667-7113-4a8f-8865-1c63f9070a9e\",\"id\":\"7c903667-7113-4a8f-8865-1c63f9070a9e\"},{\"max_update_rate_ms\":100,\"resource_path\":\"/senders\",\"params\":{},\"persist\":false,\"secure\":true,\"ws_href\":\"wss://172.29.80.52:8870/ws/?uid=c5f5c1ad-f9e6-44c0-828f-da3f7cdef286\",\"id\":\"c5f5c1ad-f9e6-44c0-828f-da3f7cdef286\"},{\"max_update_rate_ms\":100,\"resource_path\":\"/flows\",\"params\":{},\"persist\":false,\"secure\":false,\"ws_href\":\"ws://172.29.80.52:8870/ws/?uid=b0416004-f825-4623-bf90-575b5dd32f93\",\"id\":\"b0416004-f825-4623-bf90-575b5dd32f93\"}]"
	data := []byte(dataStr)
	var test []QueryAPISubscription
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

func TestQueryAPISubscriptionPostResponse(t *testing.T) {
	dataStr := "{\"max_update_rate_ms\":100,\"resource_path\":\"/nodes\",\"params\":{\"label\":\"host1\"},\"persist\":false,\"secure\":false,\"ws_href\":\"ws://172.29.80.52:8870/ws/?uid=7c903667-7113-4a8f-8865-1c63f9070a9e\",\"id\":\"7c903667-7113-4a8f-8865-1c63f9070a9e\"}"
	data := []byte(dataStr)
	var test QueryAPISubscription
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
