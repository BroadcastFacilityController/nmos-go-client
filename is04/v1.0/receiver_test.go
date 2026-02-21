package is04v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPIReceiverIDGet(t *testing.T) {
	dataStr := "{\"description\":\"\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"caps\":{},\"subscription\":{\"sender_id\":\"2683ad14-642f-459d-a169-ef91c76cec6b\"},\"version\":\"1441704532:450093308\",\"label\":\"RTPRx\",\"id\":\"1eb53d65-ac83-441c-86f6-9b27df30ef0c\",\"transport\":\"urn:x-nmos:transport:rtp\",\"device_id\":\"05017e08-b329-45f9-a566-a3f99cc11e4d\"}"
	data := []byte(dataStr)
	var test Receiver
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

func TestNodeAPIReceiversGet(t *testing.T) {
	dataStr := "[{\"description\":\"\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"caps\":{},\"subscription\":{\"sender_id\":\"2683ad14-642f-459d-a169-ef91c76cec6b\"},\"version\":\"1441704532:450093308\",\"label\":\"RTPRx\",\"id\":\"1eb53d65-ac83-441c-86f6-9b27df30ef0c\",\"transport\":\"urn:x-nmos:transport:rtp\",\"device_id\":\"05017e08-b329-45f9-a566-a3f99cc11e4d\"}]"
	data := []byte(dataStr)
	var test []Receiver
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

func TestQueryAPIReceiverIDGet(t *testing.T) {
	dataStr := "{\"description\":\"RTP Receiver\",\"tags\":{\"Location\":[\"Location 1\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"device_id\":\"0d0cb97e-b96a-4a39-887f-d491492d9081\",\"version\":\"1441895693:480000000\",\"label\":\"Viewer\",\"id\":\"3350d113-1593-4271-a7f5-f4974415bb8e\",\"transport\":\"urn:x-nmos:transport:rtp\",\"subscription\":{\"sender_id\":\"55311762-8003-48fa-a645-0a0c7621ce45\"}}"
	data := []byte(dataStr)
	var test Receiver
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

func TestQueryAPIReceiversGet(t *testing.T) {
	dataStr := "[{\"description\":\"RTP receiver 1\",\"tags\":{\"Location\":[\"Location 1\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"device_id\":\"0d0cb97e-b96a-4a39-887f-d491492d9081\",\"version\":\"1441895693:480000000\",\"label\":\"Viewer 1\",\"id\":\"3350d113-1593-4271-a7f5-f4974415bb8e\",\"transport\":\"urn:x-nmos:transport:rtp\",\"subscription\":{\"sender_id\":\"55311762-8003-48fa-a645-0a0c7621ce45\"}},{\"description\":\"RTP receiver 2\",\"tags\":{\"Location\":[\"Location 2\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"device_id\":\"76aa31e5-c3a4-4639-8dbb-98a86dc32942\",\"version\":\"1441726973:353560292\",\"label\":\"Viewer 2\",\"id\":\"3a1be8bd-6077-4933-aa7a-b42b13bdbe8e\",\"transport\":\"urn:x-nmos:transport:rtp\",\"subscription\":{\"sender_id\":\"a325cfe1-47ee-4a23-9a5c-7b5fcb9c2bb2\"}},{\"description\":\"Audio RX\",\"tags\":{},\"format\":\"urn:x-nmos:format:audio\",\"caps\":{},\"device_id\":\"e19ef82c-5f0a-48da-a86c-bb2377ab09a4\",\"version\":\"1441722334:801293520\",\"label\":\"Audio RX\",\"id\":\"a383178a-76cc-4894-9121-dc390c7847d3\",\"transport\":\"urn:x-nmos:transport:rtp\",\"subscription\":{\"sender_id\":null}}]"
	data := []byte(dataStr)
	var test []Receiver
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
