package is05v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstraints2022_7(t *testing.T) {
	dataStr := "[{\"source_ip\":{},\"multicast_ip\":{},\"interface_ip\":{\"enum\":[\"192.168.7.2\",\"192.168.8.5\",\"2001:0db8:85a3:0000:0000:8a2e:0370:7334\",\"2001:0db8:85a3:0000:0000:8a2e:0370:a24b\"]},\"destination_port\":{\"maximum\":50000,\"minimum\":49151},\"fec_enabled\":{},\"fec_destination_ip\":{},\"fec_mode\":{\"enum\":[\"1D\",\"2D\"]},\"fec1D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"fec2D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtcp_enabled\":{},\"rtcp_destination_ip\":{},\"rtcp_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtp_enabled\":{}},{\"source_ip\":{},\"multicast_ip\":{},\"interface_ip\":{\"enum\":[\"192.168.7.2\",\"192.168.8.5\",\"2001:0db8:85a3:0000:0000:8a2e:0370:7334\",\"2001:0db8:85a3:0000:0000:8a2e:0370:a24b\"]},\"destination_port\":{\"maximum\":50000,\"minimum\":49151},\"fec_enabled\":{},\"fec_destination_ip\":{},\"fec_mode\":{\"enum\":[\"1D\",\"2D\"]},\"fec_block_width\":{\"minimum\":4,\"maximum\":200},\"fec_block_height\":{\"minimum\":4,\"maximum\":200},\"fec1D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"fec2D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtcp_enabled\":{},\"rtcp_destination_ip\":{},\"rtcp_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtp_enabled\":{}}]"
	data := []byte(dataStr)
	var test []Constraints
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

func TestConstraints(t *testing.T) {
	dataStr := "[{\"source_ip\":{},\"multicast_ip\":{},\"interface_ip\":{\"enum\":[\"192.168.7.2\",\"192.168.8.5\",\"2001:0db8:85a3:0000:0000:8a2e:0370:7334\",\"2001:0db8:85a3:0000:0000:8a2e:0370:a24b\"]},\"destination_port\":{\"maximum\":50000,\"minimum\":49151},\"fec_enabled\":{},\"fec_destination_ip\":{},\"fec_mode\":{\"enum\":[\"1D\",\"2D\"]},\"fec1D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"fec2D_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtcp_enabled\":{},\"rtcp_destination_ip\":{},\"rtcp_destination_port\":{\"minimum\":5000,\"maximum\":49151},\"rtp_enabled\":{}}]"
	data := []byte(dataStr)
	var test []Constraints
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
