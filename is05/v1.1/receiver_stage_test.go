package is05v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiverPatch(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverStage
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

func TestReceiverPatchAbsolute(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_absolute\",\"requested_time\":\"1496677550:20\"},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverStage
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

func TestReceiverPatchRedundant(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"rtp_enabled\":true},{\"interface_ip\":\"192.168.210.15\",\"multicast_ip\":\"232.3.28.144\",\"source_ip\":\"192.168.210.10\",\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverStage
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

func TestReceiverPatchRelative(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"2:0\"},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverStage
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

func TestReceiverPatchUnicast(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":null,\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverStage
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
