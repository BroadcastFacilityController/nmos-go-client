package is05v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiverResponse(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_file\":{\"data\":null,\"type\":\"application/sdp\"},\"transport_params\":[{\"source_ip\":\"172.23.19.48\",\"multicast_ip\":\"232.250.98.80\",\"interface_ip\":\"172.23.19.110\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true},{\"source_ip\":\"172.23.20.48\",\"multicast_ip\":\"232.250.99.80\",\"interface_ip\":\"172.23.20.110\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverResponse
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

func TestReceiverResponseUnunit(t *testing.T) {
	dataStr := "{\"sender_id\":null,\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"source_ip\":null,\"multicast_ip\":null,\"interface_ip\":\"auto\",\"destination_port\":\"auto\",\"fec_enabled\":false,\"fec_destination_ip\":\"auto\",\"fec_mode\":\"1D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":false,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":\"auto\",\"rtp_enabled\":true}]}"
	data := []byte(dataStr)
	var test ReceiverResponse
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
