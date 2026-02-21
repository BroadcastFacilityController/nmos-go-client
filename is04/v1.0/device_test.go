package is04v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPIDeviceIDGet(t *testing.T) {
	dataStr := "{\"receivers\":[],\"label\":\"pipeline 1 default device\",\"version\":\"1441703338:962976113\",\"id\":\"67c25159-ce25-4000-a66c-f31fff890265\",\"type\":\"urn:x-nmos:device:pipeline\",\"senders\":[],\"node_id\":\"3b8be755-08ff-452b-b217-c9151eb21193\"}"
	data := []byte(dataStr)
	var test Device
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

func TestNodeAPIDevicesGet(t *testing.T) {
	dataStr := "[{\"receivers\":[],\"label\":\"pipeline 3 default device\",\"version\":\"1441704616:592733242\",\"id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"type\":\"urn:x-nmos:device:pipeline\",\"senders\":[\"d7aa5a30-681d-4e72-92fb-f0ba0f6f4c3e\"],\"node_id\":\"3b8be755-08ff-452b-b217-c9151eb21193\"},{\"receivers\":[],\"label\":\"pipeline 1 default device\",\"version\":\"1441703338:962976113\",\"id\":\"67c25159-ce25-4000-a66c-f31fff890265\",\"type\":\"urn:x-nmos:device:pipeline\",\"senders\":[],\"node_id\":\"3b8be755-08ff-452b-b217-c9151eb21193\"},{\"receivers\":[\"1eb53d65-ac83-441c-86f6-9b27df30ef0c\"],\"label\":\"pipeline 2 default device\",\"version\":\"1441704514:993221361\",\"id\":\"05017e08-b329-45f9-a566-a3f99cc11e4d\",\"type\":\"urn:x-nmos:device:pipeline\",\"senders\":[],\"node_id\":\"3b8be755-08ff-452b-b217-c9151eb21193\"}]"
	data := []byte(dataStr)
	var test []Device
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

func TestQueryAPIDeviceIDGet(t *testing.T) {
	dataStr := "{\"receivers\":[\"863532de-a97d-4597-989a-e79688f2d5f9\",\"632d7e6d-7357-44de-a425-a94fbe94974e\",\"95ef711b-564d-4655-a98b-5b9ccfb419d7\",\"9ee74607-f831-42f5-af08-a614ce0706df\",\"1311bf13-869c-45b4-915e-8b4e8b8e26fd\",\"debcd758-129d-4e0b-a0e3-f1d9ce5edfbc\",\"0fee9741-e266-4c27-b480-1f897463ea4b\",\"ef9e58bd-431a-466d-a67f-0318858b981c\",\"68ca0867-ec4b-4eca-92e0-4c4c668a72b0\"],\"label\":\"pipeline 1 default device\",\"version\":\"1441723676:366608283\",\"senders\":[],\"id\":\"a370d258-69de-4422-860a-ee4cf32ee9f4\",\"node_id\":\"3a25a674-e6eb-4987-84ad-ef479fe4d527\",\"type\":\"urn:x-nmos:device:pipeline\"}"
	data := []byte(dataStr)
	var test Device
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

func TestQueryAPIDevicesGet(t *testing.T) {
	dataStr := "[{\"receivers\":[],\"label\":\"pipeline 1 default device\",\"version\":\"1441723957:582701772\",\"senders\":[\"c72cca5b-01db-47aa-bb00-03893defbfae\",\"171d5c80-7fff-4c23-9383-46503eb1c63e\",\"a2655c48-8a46-4c82-b9bc-98760d59d7f8\"],\"id\":\"c501ae64-f525-48b7-9816-c5e8931bc017\",\"node_id\":\"1d452562-e1d5-4e84-b057-5c24de5f6b48\",\"type\":\"urn:x-nmos:device:pipeline\"},{\"receivers\":[],\"label\":\"pipeline 1 default device\",\"version\":\"1441976012:727999141\",\"senders\":[],\"id\":\"a30e4fba-254a-4e97-8bf7-daec80b8e57f\",\"node_id\":\"3b8be755-08ff-452b-b217-c9151eb21193\",\"type\":\"urn:x-nmos:device:pipeline\"},{\"receivers\":[\"a383178a-76cc-4894-9121-dc390c7847d3\"],\"label\":\"pipeline 1 default device\",\"version\":\"1441722334:834709519\",\"senders\":[\"24da6ba7-b0aa-4883-a89a-2b867328dbf9\"],\"id\":\"e19ef82c-5f0a-48da-a86c-bb2377ab09a4\",\"node_id\":\"4cf38bb4-d6c4-48d6-a086-6eac45d73ae5\",\"type\":\"urn:x-nmos:device:pipeline\"},{\"receivers\":[\"863532de-a97d-4597-989a-e79688f2d5f9\",\"632d7e6d-7357-44de-a425-a94fbe94974e\",\"95ef711b-564d-4655-a98b-5b9ccfb419d7\",\"9ee74607-f831-42f5-af08-a614ce0706df\",\"1311bf13-869c-45b4-915e-8b4e8b8e26fd\",\"debcd758-129d-4e0b-a0e3-f1d9ce5edfbc\",\"0fee9741-e266-4c27-b480-1f897463ea4b\",\"ef9e58bd-431a-466d-a67f-0318858b981c\",\"68ca0867-ec4b-4eca-92e0-4c4c668a72b0\"],\"label\":\"pipeline 1 default device\",\"version\":\"1441723676:366608283\",\"senders\":[],\"id\":\"a370d258-69de-4422-860a-ee4cf32ee9f4\",\"node_id\":\"3a25a674-e6eb-4987-84ad-ef479fe4d527\",\"type\":\"urn:x-nmos:device:pipeline\"}]"
	data := []byte(dataStr)
	var test []Device
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
