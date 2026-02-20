package is04v1_2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPISelfGet(t *testing.T) {
	dataStr := "{\"version\":\"1441700172:318426300\",\"hostname\":\"host1\",\"caps\":{},\"href\":\"http://172.29.80.65:12345/\",\"api\":{\"versions\":[\"v1.0\",\"v1.1\"],\"endpoints\":[{\"host\":\"172.29.80.65\",\"port\":12345,\"protocol\":\"http\"},{\"host\":\"172.29.80.65\",\"port\":443,\"protocol\":\"https\"}]},\"services\":[{\"href\":\"http://172.29.80.65:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/tally/\",\"type\":\"urn:x-manufacturer:service:tally\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"},{\"href\":\"http://172.29.80.65:12345/x-manufacturer/storequery/\",\"type\":\"urn:x-manufacturer:service:storequery\"}],\"label\":\"host1\",\"description\":\"host1\",\"tags\":{},\"id\":\"3b8be755-08ff-452b-b217-c9151eb21193\",\"clocks\":[{\"name\":\"clk0\",\"ref_type\":\"internal\"},{\"name\":\"clk1\",\"ref_type\":\"ptp\",\"traceable\":true,\"version\":\"IEEE1588-2008\",\"gmid\":\"08-00-11-ff-fe-21-e1-b0\",\"locked\":true}],\"interfaces\":[{\"name\":\"eth0\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-31\"},{\"name\":\"eth1\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-32\"}]}"
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
	dataStr := "{\"version\":\"1441716120:318744030\",\"hostname\":\"host1\",\"label\":\"host1\",\"description\":\"host1\",\"tags\":{},\"href\":\"http://172.29.176.102:12345/\",\"api\":{\"versions\":[\"v1.1\"],\"endpoints\":[{\"host\":\"172.29.176.102\",\"port\":12345,\"protocol\":\"http\"}]},\"services\":[{\"href\":\"http://172.29.176.102:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"c8ba20e9-e197-4ec5-8764-4da672128589\",\"clocks\":[{\"name\":\"clk0\",\"ref_type\":\"internal\"},{\"name\":\"clk1\",\"ref_type\":\"ptp\",\"traceable\":true,\"version\":\"IEEE1588-2008\",\"gmid\":\"08-00-11-ff-fe-21-e1-b0\",\"locked\":true}],\"interfaces\":[{\"name\":\"eth0\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-31\"},{\"name\":\"eth1\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-32\"}]}"
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
	dataStr := "[{\"version\":\"1441716120:318744030\",\"hostname\":\"host1\",\"label\":\"host1\",\"description\":\"host1\",\"tags\":{},\"href\":\"http://172.29.176.102:12345/\",\"api\":{\"versions\":[\"v1.1\"],\"endpoints\":[{\"host\":\"172.29.176.102\",\"port\":12345,\"protocol\":\"http\"}]},\"services\":[{\"href\":\"http://172.29.176.102:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.102:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"c8ba20e9-e197-4ec5-8764-4da672128589\",\"clocks\":[{\"name\":\"clk0\",\"ref_type\":\"ptp\",\"traceable\":false,\"version\":\"IEEE1588-2008\",\"gmid\":\"08-00-11-ff-fe-21-e1-b0\",\"locked\":true}],\"interfaces\":[{\"name\":\"eth0\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-31\"},{\"name\":\"eth1\",\"chassis_id\":\"74-26-96-db-87-31\",\"port_id\":\"74-26-96-db-87-32\"}]},{\"version\":\"1441716067:450064243\",\"hostname\":\"host2\",\"label\":\"host2\",\"description\":\"host2\",\"tags\":{},\"href\":\"http://172.29.176.19:12345/\",\"api\":{\"versions\":[\"v1.1\"],\"endpoints\":[{\"host\":\"172.29.176.19\",\"port\":12345,\"protocol\":\"http\"}]},\"services\":[{\"href\":\"http://172.29.176.19:12345/x-manufacturer/pipelinemanager/\",\"type\":\"urn:x-manufacturer:service:pipelinemanager\"},{\"href\":\"http://172.29.176.19:12345/x-manufacturer/status/\",\"type\":\"urn:x-manufacturer:service:status\"},{\"href\":\"http://172.29.176.19:12345/x-manufacturer/mdnsbridge/\",\"type\":\"urn:x-manufacturer:service:mdnsbridge\"}],\"caps\":{},\"id\":\"cebc6305-e8db-4026-aeb5-eb7a5620839e\",\"clocks\":[{\"name\":\"clk0\",\"ref_type\":\"ptp\",\"traceable\":false,\"version\":\"IEEE1588-2008\",\"gmid\":\"08-00-11-ff-fe-21-e1-b0\",\"locked\":true}],\"interfaces\":[{\"name\":\"en0\",\"chassis_id\":\"a4-26-84-db-58-31\",\"port_id\":\"a4-26-84-db-58-31\"},{\"name\":\"en1\",\"chassis_id\":\"a4-26-84-db-58-31\",\"port_id\":\"a4-26-84-db-58-32\"}]}]"
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
