package is04v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPISelfGet(t *testing.T) {
	dataStr := "{\"version\":\"1441700172:318426300\",\"hostname\":\"host1\",\"caps\":{},\"href\":\"http://172.29.80.65:12345/\",\"services\":[{\"href\":\"http://172.29.80.65:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/tally/\",\"type\":\"urn:x-manufacturer:service:tally\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/storequery/\",\"type\":\"urn:x-manufacturer:service:storequery\"}],\"label\":\"host1\",\"id\":\"3b8be755-08ff-452b-b217-c9151eb21193\"}"
	data := []byte(dataStr)
	var test Node
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

func TestQueryAPINodeIDGet(t *testing.T) {
	dataStr := "{\"version\":\"1441716120:318744030\",\"hostname\":\"host1\",\"label\":\"host1\",\"href\":\"http://172.29.176.102:12345/\",\"services\":[{\"href\":\"http://172.29.176.102:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"c8ba20e9-e197-4ec5-8764-4da672128589\"}"
	data := []byte(dataStr)
	var test Node
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

func TestQueryAPINodesGet(t *testing.T) {
	dataStr := "[{\"version\":\"1441716120:318744030\",\"hostname\":\"host1\",\"label\":\"host1\",\"href\":\"http://172.29.176.102:12345/\",\"services\":[{\"href\":\"http://172.29.176.102:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"c8ba20e9-e197-4ec5-8764-4da672128589\"},{\"version\":\"1441716067:450064243\",\"hostname\":\"host2\",\"label\":\"host2\",\"href\":\"http://172.29.176.19:12345/\",\"services\":[{\"href\":\"http://172.29.176.19:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.19:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.19:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"cebc6305-e8db-4026-aeb5-eb7a5620839e\"}]"
	data := []byte(dataStr)
	var test []Node
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
