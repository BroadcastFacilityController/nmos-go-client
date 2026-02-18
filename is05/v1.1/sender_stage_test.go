package is05v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSenderStagedPatch(t *testing.T) {
	dataStr := "{\"receiver_id\":null,\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderStage
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

func TestSenderStagedPatchAbsolute(t *testing.T) {
	dataStr := "{\"receiver_id\":null,\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_absolute\",\"requested_time\":\"1496759200:0\"},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderStage
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

func TestSenderStagedPatchRelative(t *testing.T) {
	dataStr := "{\"receiver_id\":null,\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"2:0\"},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderStage
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

func TestSenderStagedPatchRedundant(t *testing.T) {
	dataStr := "{\"receiver_id\":null,\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"232.105.26.177\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true},{\"source_ip\":\"192.168.210.10\",\"destination_ip\":\"232.3.28.144\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderStage
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

func TestSenderStagedPatchUnicast(t *testing.T) {
	dataStr := "{\"receiver_id\":\"0b174530-e3cf-41e6-bf01-fe55135034f3\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"2:0\"},\"transport_params\":[{\"source_ip\":\"192.168.200.10\",\"destination_ip\":\"192.168.200.15\",\"source_port\":5000,\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test SenderStage
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
