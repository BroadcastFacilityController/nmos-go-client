package is05v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSenderGet(t *testing.T) {
	dataStr := "{\"receiver_id\":\"0a174530-e3cf-11e6-bf01-fe55135034f3\",\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_params\":[{\"source_ip\":\"172.29.82.23\",\"destination_ip\":\"172.29.82.95\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"172.29.82.24\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true},{\"source_ip\":\"172.30.82.23\",\"destination_ip\":\"172.30.82.95\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"172.30.82.24\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestSenderGetUninit(t *testing.T) {
	dataStr := "{\"receiver_id\":null,\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_params\":[{\"source_ip\":\"auto\",\"destination_ip\":\"auto\",\"source_port\":\"auto\",\"destination_port\":\"auto\",\"fec_enabled\":false,\"fec_destination_ip\":\"auto\",\"fec_type\":\"XOR\",\"fec_mode\":\"1D\",\"fec_block_width\":4,\"fec_block_height\":4,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":\"auto\",\"fec2D_source_port\":\"auto\",\"rtcp_enabled\":false,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":\"auto\",\"rtcp_source_port\":\"auto\",\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestSenderStagedGetAbsolute(t *testing.T) {
	dataStr := "{\"master_enable\":true,\"receiver_id\":null,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"1496759200:0\",\"activation_time\":\"1496759200:5000\"},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.22.4.59\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestSenderStagedGetRelative(t *testing.T) {
	dataStr := "{\"master_enable\":true,\"receiver_id\":null,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"2:0\",\"activation_time\":\"1496759200:0\"},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.79.192.98\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestSenderStagedGetRedundant(t *testing.T) {
	dataStr := "{\"master_enable\":true,\"receiver_id\":null,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null,\"activation_time\":null},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.22.4.59\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true},{\"source_ip\":\"192.168.210.10\",\"destination_ip\":\"232.3.28.144\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.79.192.98\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestSenderStagedGet(t *testing.T) {
	dataStr := "{\"master_enable\":true,\"receiver_id\":null,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null,\"activation_time\":null},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"fec_enabled\":true,\"fec_type\":\"XOR\",\"fec_mode\":\"2D\",\"fec_block_width\":50,\"fec_block_height\":50,\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"fec1D_source_port\":5010,\"fec2D_source_port\":5012,\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.22.4.59\",\"rtcp_destination_port\":5008,\"rtcp_source_port\":5008,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderResponse
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
		return
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
		return
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}
