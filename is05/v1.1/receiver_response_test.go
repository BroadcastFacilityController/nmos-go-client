package is05v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiverResponseActive(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":\"1496739062:0\",\"activation_time\":\"1496739062:263900\"},\"transport_file\":{\"data\":\"v=0\\r\\no=- 1496222842 1496222842 IN IP4 172.29.226.25\\r\\ns=IP Studio Stream\\r\\nt=0 0\\r\\nm=video 5010 RTP/AVP 103\\r\\nc=IN IP4 232.250.98.80/32\\r\\na=source-filter: incl IN IP4 232.250.98.80 172.29.226.25\\r\\na=rtpmap:103 raw/90000\\r\\na=fmtp:103 sampling=YCbCr-4:2:2; width=1920; height=1080; depth=10; interlace; exactframerate=25; colorimetry=BT709; PM=2110GPM; SSN=ST2110-20:2017; TP=2110TPW; \\r\\na=mediaclk:direct=1876655126 rate=90000\\r\\na=extmap:1 urn:x-nmos:rtp-hdrext:origin-timestamp\\r\\na=extmap:2 urn:ietf:params:rtp-hdrext:smpte-tc 3600@90000/25\\r\\na=extmap:3 urn:x-nmos:rtp-hdrext:flow-id\\r\\na=extmap:4 urn:x-nmos:rtp-hdrext:source-id\\r\\na=extmap:5 urn:x-nmos:rtp-hdrext:grain-flags\\r\\na=extmap:7 urn:x-nmos:rtp-hdrext:sync-timestamp\\r\\na=extmap:9 urn:x-nmos:rtp-hdrext:grain-duration\\r\\na=ts-refclk:ptp=IEEE1588-2008:08-00-11-FF-FE-21-E1-B0:0\\r\\n\",\"type\":\"application/sdp\"},\"transport_params\":[{\"source_ip\":\"172.29.226.25\",\"multicast_ip\":\"232.250.98.80\",\"interface_ip\":\"172.23.19.35\",\"destination_port\":5010,\"fec_enabled\":false,\"fec_mode\":\"2D\",\"fec1D_destination_port\":5012,\"fec2D_destination_port\":5014,\"rtcp_enabled\":false,\"rtcp_destination_ip\":\"232.250.98.80\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseActiveUninit(t *testing.T) {
	dataStr := "{\"sender_id\":null,\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"source_ip\":null,\"multicast_ip\":null,\"interface_ip\":\"172.23.19.35\",\"destination_port\":5004,\"fec_enabled\":false,\"fec_destination_ip\":\"172.23.19.35\",\"fec1D_destination_port\":5006,\"fec2D_destination_port\":5008,\"rtcp_enabled\":false,\"rtcp_destination_ip\":\"172.23.19.35\",\"rtcp_destination_port\":5005,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseActiveMQTT(t *testing.T) {
	dataStr := "{\"sender_id\":\"eae33112-b56c-4a36-8b68-05ab5b6268bc\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":\"1541508905:0\",\"activation_time\":\"1541508905:263900\"},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"source_host\":\"broker.example.com\",\"source_port\":1883,\"broker_topic\":\"my_topic_name\",\"broker_protocol\":\"mqtt\",\"broker_authorization\":false,\"connection_status_broker_topic\":\"my_status_topic_name\"}]}"
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

func TestReceiverResponseGetUninit(t *testing.T) {
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

func TestReceiverResponseGet(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"source_ip\":\"172.23.19.48\",\"multicast_ip\":\"232.250.98.80\",\"interface_ip\":\"172.23.19.110\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true},{\"source_ip\":\"172.23.20.48\",\"multicast_ip\":\"232.250.99.80\",\"interface_ip\":\"172.23.20.110\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_destination_ip\":\"auto\",\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseStagedAbsolute(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_absolute\",\"requested_time\":\"1496739064:0\",\"activation_time\":\"1496739064:2500\"},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseStagedRelative(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_scheduled_relative\",\"requested_time\":\"2:0\",\"activation_time\":\"1496739064:0\"},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseStagedTransportFile(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":null,\"requested_time\":null,\"activation_time\":null},\"transport_file\":{\"data\":\"v=0\\r\\no=- 1496222842 1496222842 IN IP4 172.29.226.25\\r\\ns=IP Studio Stream\\r\\nt=0 0\\r\\nm=video 5010 RTP/AVP 103\\r\\nc=IN IP4 232.250.98.80/32\\r\\na=source-filter: incl IN IP4 232.250.98.80 172.29.226.25\\r\\na=rtpmap:103 raw/90000\\r\\na=fmtp:103 sampling=YCbCr-4:2:2; width=1920; height=1080; depth=10; interlace; exactframerate=25; colorimetry=BT709; PM=2110GPM; SSN=ST2110-20:2017; TP=2110TPW; \\r\\na=mediaclk:direct=1876655126 rate=90000\\r\\na=extmap:1 urn:x-nmos:rtp-hdrext:origin-timestamp\\r\\na=extmap:2 urn:ietf:params:rtp-hdrext:smpte-tc 3600@90000/25\\r\\na=extmap:3 urn:x-nmos:rtp-hdrext:flow-id\\r\\na=extmap:4 urn:x-nmos:rtp-hdrext:source-id\\r\\na=extmap:5 urn:x-nmos:rtp-hdrext:grain-flags\\r\\na=extmap:7 urn:x-nmos:rtp-hdrext:sync-timestamp\\r\\na=extmap:9 urn:x-nmos:rtp-hdrext:grain-duration\\r\\na=ts-refclk:ptp=IEEE1588-2008:08-00-11-FF-FE-21-E1-B0:0\\r\\n\",\"type\":\"application/sdp\"},\"transport_params\":[{\"source_ip\":\"172.29.226.25\",\"multicast_ip\":\"232.250.98.80\",\"interface_ip\":\"172.23.19.35\",\"destination_port\":5010,\"fec_enabled\":false,\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":false,\"rtcp_destination_ip\":\"auto\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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

func TestReceiverResponseStagedRedundantStreams(t *testing.T) {
	dataStr := "{\"sender_id\":\"5709255c-c0ae-4e1e-99a0-e872e83e48e0\",\"master_enable\":true,\"activation\":{\"mode\":\"activate_immediate\",\"requested_time\":null,\"activation_time\":\"1496739062:0\"},\"transport_file\":{\"data\":null,\"type\":null},\"transport_params\":[{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.22.4.59\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true},{\"interface_ip\":\"192.168.200.15\",\"multicast_ip\":\"232.105.26.177\",\"source_ip\":\"192.168.200.10\",\"destination_port\":5000,\"fec_enabled\":true,\"fec_mode\":\"2D\",\"fec1D_destination_port\":\"auto\",\"fec2D_destination_port\":\"auto\",\"rtcp_enabled\":true,\"rtcp_destination_ip\":\"232.79.192.98\",\"rtcp_destination_port\":5001,\"rtp_enabled\":true}]}"
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
