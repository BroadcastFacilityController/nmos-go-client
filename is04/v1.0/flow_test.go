package is04v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPIFlowIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Test Card\",\"tags\":{},\"format\":\"urn:x-nmos:format:video\",\"label\":\"Test Card\",\"version\":\"1441704616:587121295\",\"parents\":[],\"source_id\":\"02c46999-d532-4c52-905f-2e368a2af6cb\",\"id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\"}"
	data := []byte(dataStr)
	var test Flow
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

func TestNodeAPIFlowsGet(t *testing.T) {
	dataStr := "[{\"description\":\"Test Card\",\"tags\":{},\"format\":\"urn:x-nmos:format:video\",\"label\":\"Test Card\",\"version\":\"1441704616:587121295\",\"parents\":[],\"source_id\":\"02c46999-d532-4c52-905f-2e368a2af6cb\",\"id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\"},{\"description\":\"VANC Data\",\"tags\":{},\"format\":\"urn:x-nmos:format:data\",\"label\":\"VANC Data\",\"version\":\"1453880607:123995943\",\"parents\":[],\"source_id\":\"0e635152-e501-4d4e-bb87-9f3fe05eb79a\",\"id\":\"db3bd465-2772-484f-8fac-830b0471258b\"}]"
	data := []byte(dataStr)
	var test []Flow
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

func TestQueryAPIFlowIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Capture Audio Proxy\",\"format\":\"urn:x-nmos:format:audio\",\"tags\":{\"host\":[\"host1\"]},\"label\":\"Capture Audio Proxy\",\"version\":\"1441812152:154331951\",\"parents\":[],\"source_id\":\"2aa143ac-0ab7-4d75-bc32-5c00c13d186f\",\"id\":\"b3bb5be7-9fe9-4324-a5bb-4c70e1084449\"}"
	data := []byte(dataStr)
	var test Flow
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

func TestQueryAPIFlowsGet(t *testing.T) {
	dataStr := "[{\"description\":\"Off-air proxy\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"label\":\"Off-air proxy\",\"version\":\"1441724130:194944510\",\"parents\":[],\"source_id\":\"c7b1c809-84d4-427d-b6bb-c8397c66ce2b\",\"id\":\"0c1f03d7-7e94-4b21-94d1-3ffbee8a0606\"},{\"description\":\"Off-air\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"label\":\"Off-air\",\"version\":\"1441724130:186085940\",\"parents\":[],\"source_id\":\"c7b1c809-84d4-427d-b6bb-c8397c66ce2b\",\"id\":\"0e85d87b-4b19-4452-aea3-984c9f94bbc9\"},{\"description\":\"Capture Audio Proxy\",\"format\":\"urn:x-nmos:format:audio\",\"tags\":{\"host\":[\"host1\"]},\"label\":\"Capture Audio Proxy\",\"version\":\"1441812152:154331951\",\"parents\":[],\"source_id\":\"2aa143ac-0ab7-4d75-bc32-5c00c13d186f\",\"id\":\"b3bb5be7-9fe9-4324-a5bb-4c70e1084449\"}]"
	data := []byte(dataStr)
	var test []Flow
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
