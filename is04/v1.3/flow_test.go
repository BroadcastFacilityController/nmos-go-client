package is04v1_3

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPIFlowIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Test Card\",\"tags\":{},\"format\":\"urn:x-nmos:format:video\",\"label\":\"Test Card\",\"version\":\"1441704616:587121295\",\"parents\":[],\"source_id\":\"02c46999-d532-4c52-905f-2e368a2af6cb\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\",\"media_type\":\"video/raw\",\"frame_width\":1920,\"frame_height\":1080,\"interlace_mode\":\"interlaced_tff\",\"colorspace\":\"BT709\",\"components\":[{\"name\":\"Y\",\"width\":1920,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cb\",\"width\":960,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cr\",\"width\":960,\"height\":1080,\"bit_depth\":10}]}"
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
	dataStr := "[{\"description\":\"Test Card\",\"tags\":{},\"format\":\"urn:x-nmos:format:video\",\"label\":\"Test Card\",\"version\":\"1441704616:587121295\",\"parents\":[],\"source_id\":\"02c46999-d532-4c52-905f-2e368a2af6cb\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"id\":\"5fbec3b1-1b0f-417d-9059-8b94a47197ed\",\"media_type\":\"video/raw\",\"frame_width\":1920,\"frame_height\":1080,\"interlace_mode\":\"interlaced_tff\",\"colorspace\":\"BT709\",\"components\":[{\"name\":\"Y\",\"width\":1920,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cb\",\"width\":960,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cr\",\"width\":960,\"height\":1080,\"bit_depth\":10}]},{\"description\":\"VANC Data\",\"tags\":{},\"format\":\"urn:x-nmos:format:data\",\"label\":\"VANC Data\",\"version\":\"1453880607:123995943\",\"parents\":[],\"source_id\":\"0e635152-e501-4d4e-bb87-9f3fe05eb79a\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"id\":\"db3bd465-2772-484f-8fac-830b0471258b\",\"media_type\":\"video/smpte291\"},{\"description\":\"TR-04 Video\",\"tags\":{},\"format\":\"urn:x-nmos:format:mux\",\"label\":\"TR-04 Video\",\"version\":\"1453880607:123995943\",\"parents\":[],\"source_id\":\"782fac41-17f6-4a21-8186-57ba63a1a8d3\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"id\":\"4857f747-96cf-4ed7-8f4b-9497199f1f25\",\"media_type\":\"video/SMPTE2022-6\"},{\"description\":\"IS-07 Temperature Readings (C)\",\"tags\":{},\"format\":\"urn:x-nmos:format:data\",\"event_type\":\"number/temperature/C\",\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"IS-07 Temperature Readings (C)\",\"id\":\"6327c381-1239-41d1-b314-efc719600e26\",\"source_id\":\"33e28c6f-d5ab-4ae5-b00d-f1cccab29af4\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"media_type\":\"application/json\"},{\"description\":\"IS-07 Temperature Readings (F)\",\"tags\":{},\"format\":\"urn:x-nmos:format:data\",\"event_type\":\"number/temperature/F\",\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"IS-07 Temperature Readings (F)\",\"id\":\"6327c381-1239-41d1-b315-efc719600e26\",\"source_id\":\"33e28c6f-d5ab-4ae5-b00d-f1cccab29af4\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"media_type\":\"application/json\"},{\"description\":\"IS-07 Button\",\"tags\":{},\"format\":\"urn:x-nmos:format:data\",\"event_type\":\"boolean\",\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"IS-07 Button\",\"id\":\"fa6258b9-2826-4a0d-81d0-7da9edbc405f\",\"source_id\":\"c8d27a1d-d124-4d06-bc43-312fd36f7db1\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"media_type\":\"application/json\"}]"
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
	dataStr := "{\"description\":\"Capture Audio Proxy\",\"format\":\"urn:x-nmos:format:audio\",\"tags\":{\"host\":[\"host1\"]},\"label\":\"Capture Audio Proxy\",\"version\":\"1441812152:154331951\",\"parents\":[],\"source_id\":\"2aa143ac-0ab7-4d75-bc32-5c00c13d186f\",\"device_id\":\"169feb2c-3fae-42a5-ae2e-f6f8cbce29cf\",\"id\":\"b3bb5be7-9fe9-4324-a5bb-4c70e1084449\",\"media_type\":\"audio/L16\",\"sample_rate\":{\"numerator\":48000},\"bit_depth\":16}"
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
	dataStr := "[{\"description\":\"Off-air proxy\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"label\":\"Off-air proxy\",\"version\":\"1441724130:194944510\",\"parents\":[],\"source_id\":\"c7b1c809-84d4-427d-b6bb-c8397c66ce2b\",\"device_id\":\"a711ddd3-376d-4d6d-934d-834e54e045b1\",\"id\":\"0c1f03d7-7e94-4b21-94d1-3ffbee8a0606\",\"media_type\":\"video/H264\",\"frame_width\":960,\"frame_height\":540,\"interlace_mode\":\"interlaced_tff\",\"colorspace\":\"BT709\"},{\"description\":\"Off-air\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{},\"label\":\"Off-air\",\"version\":\"1441724130:186085940\",\"parents\":[],\"source_id\":\"c7b1c809-84d4-427d-b6bb-c8397c66ce2b\",\"device_id\":\"a711ddd3-376d-4d6d-934d-834e54e045b1\",\"id\":\"0e85d87b-4b19-4452-aea3-984c9f94bbc9\",\"media_type\":\"video/raw\",\"frame_width\":1920,\"frame_height\":1080,\"interlace_mode\":\"interlaced_tff\",\"colorspace\":\"BT709\",\"components\":[{\"name\":\"Y\",\"width\":1920,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cb\",\"width\":960,\"height\":1080,\"bit_depth\":10},{\"name\":\"Cr\",\"width\":960,\"height\":1080,\"bit_depth\":10}]},{\"description\":\"Capture Audio Proxy\",\"format\":\"urn:x-nmos:format:audio\",\"tags\":{\"host\":[\"host1\"]},\"label\":\"Capture Audio Proxy\",\"version\":\"1441812152:154331951\",\"parents\":[],\"source_id\":\"2aa143ac-0ab7-4d75-bc32-5c00c13d186f\",\"device_id\":\"169feb2c-3fae-42a5-ae2e-f6f8cbce29cf\",\"id\":\"b3bb5be7-9fe9-4324-a5bb-4c70e1084449\",\"media_type\":\"audio/L16\",\"sample_rate\":{\"numerator\":48000},\"bit_depth\":16},{\"description\":\"TR-04 Video\",\"tags\":{},\"format\":\"urn:x-nmos:format:mux\",\"label\":\"TR-04 Video\",\"version\":\"1453880607:123995943\",\"parents\":[],\"source_id\":\"782fac41-17f6-4a21-8186-57ba63a1a8d3\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"id\":\"4857f747-96cf-4ed7-8f4b-9497199f1f25\",\"media_type\":\"video/SMPTE2022-6\"}]"
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
