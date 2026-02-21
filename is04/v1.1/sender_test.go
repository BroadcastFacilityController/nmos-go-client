package is04v1_1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPISenderIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Test Card\",\"label\":\"Test Card\",\"version\":\"1441704616:890020555\",\"manifest_href\":\"http://172.29.80.65/x-manufacturer/senders/d7aa5a30-681d-4e72-92fb-f0ba0f6f4c3e/stream.sdp\",\"flow_id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\",\"id\":\"d7aa5a30-681d-4e72-92fb-f0ba0f6f4c3e\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"tags\":{}}"
	data := []byte(dataStr)
	var test Sender
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

func TestNodeAPISendersGet(t *testing.T) {
	dataStr := "[{\"description\":\"Test Card\",\"label\":\"Test Card\",\"version\":\"1441704616:890020555\",\"manifest_href\":\"http://172.29.80.65/x-manufacturer/senders/d7aa5a30-681d-4e72-92fb-f0ba0f6f4c3e/stream.sdp\",\"flow_id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\",\"id\":\"d7aa5a30-681d-4e72-92fb-f0ba0f6f4c3e\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"tags\":{}}]"
	data := []byte(dataStr)
	var test []Sender
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

func TestQueryAPISenderIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Camera 1\",\"label\":\"Camera 1\",\"manifest_href\":\"http://172.29.176.146/x-manufacturer/senders/171d5c80-7fff-4c23-9383-46503eb1c63e/stream.sdp\",\"version\":\"1441723958:235623703\",\"flow_id\":\"b25d445a-20dc-4937-a8a1-5cb3d5c613ee\",\"id\":\"171d5c80-7fff-4c23-9383-46503eb1c63e\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"c501ae64-f525-48b7-9816-c5e8931bc017\",\"tags\":{}}"
	data := []byte(dataStr)
	var test Sender
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

func TestQueryAPISendersGet(t *testing.T) {
	dataStr := "[{\"description\":\"Camera 1\",\"label\":\"Camera 1\",\"manifest_href\":\"http://172.29.80.25/x-manufacturer/senders/4002d6b5-5775-4975-9859-5b330fcea288/stream.sdp\",\"version\":\"1441724086:828491206\",\"flow_id\":\"75bfddce-7142-4368-bb4a-8a82c837c6df\",\"id\":\"4002d6b5-5775-4975-9859-5b330fcea288\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"3a98e804-9871-4699-ba31-d608309d8933\",\"tags\":{}},{\"description\":\"Camera 2\",\"label\":\"Camera 2\",\"manifest_href\":\"http://172.29.176.146/x-manufacturer/senders/171d5c80-7fff-4c23-9383-46503eb1c63e/stream.sdp\",\"version\":\"1441723958:235623703\",\"flow_id\":\"b25d445a-20dc-4937-a8a1-5cb3d5c613ee\",\"id\":\"171d5c80-7fff-4c23-9383-46503eb1c63e\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"c501ae64-f525-48b7-9816-c5e8931bc017\",\"tags\":{}},{\"description\":\"Camera 2 Audio\",\"label\":\"Camera 2 Audio\",\"manifest_href\":\"http://172.29.176.81/x-manufacturer/senders/bb793530-8fd7-49f9-8514-314126bbc624/stream.sdp\",\"version\":\"1441724039:737277493\",\"flow_id\":\"4ffa5719-9ab9-4395-bedb-3534fa7ba438\",\"id\":\"bb793530-8fd7-49f9-8514-314126bbc624\",\"transport\":\"urn:x-nmos:transport:rtp.mcast\",\"device_id\":\"5f88d383-596c-409c-887e-a90e42ef3684\",\"tags\":{}}]"
	data := []byte(dataStr)
	var test []Sender
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
