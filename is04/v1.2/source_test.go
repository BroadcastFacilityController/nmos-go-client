package is04v1_2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeAPISourceIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Capture Card Source Video\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"version\":\"1441703336:902850419\",\"parents\":[],\"label\":\"CaptureCardSourceVideo\",\"id\":\"4569cea2-ab63-4f97-8dd1-bad4669ea5e4\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk1\"}"
	data := []byte(dataStr)
	var test Source
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

func TestNodeAPISourcesGet(t *testing.T) {
	dataStr := "[{\"description\":\"Capture Card Source Video\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"version\":\"1441703336:902850419\",\"parents\":[],\"label\":\"CaptureCardSourceVideo\",\"id\":\"4569cea2-ab63-4f97-8dd1-bad4669ea5e4\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\"},{\"description\":\"Capture Card Source Audio\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:audio\",\"caps\":{},\"version\":\"1441703336:912670314\",\"parents\":[],\"label\":\"CaptureCardSourceAudio\",\"id\":\"fc97ab0f-b51b-4129-9385-dcaf30f9482b\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\",\"channels\":[{\"label\":\"Left Channel\",\"symbol\":\"L\"},{\"label\":\"Right Channel\",\"symbol\":\"R\"}]},{\"description\":\"Capture Card Source Audio\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:audio\",\"caps\":{},\"version\":\"1441704614:174935325\",\"parents\":[],\"label\":\"CaptureCardSourceAudio\",\"id\":\"9738780e-141f-4e19-8601-a157dc855aa2\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\",\"channels\":[{\"label\":\"Left Channel\",\"symbol\":\"L\"},{\"label\":\"Right Channel\",\"symbol\":\"R\"}]},{\"description\":\"Capture Card Source Video\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:video\",\"caps\":{},\"version\":\"1441704614:163285887\",\"parents\":[],\"label\":\"CaptureCardSourceVideo\",\"id\":\"02c46999-d532-4c52-905f-2e368a2af6cb\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\"},{\"description\":\"Capture Card Source VANC\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:data\",\"caps\":{},\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"Capture Card Source VANC\",\"id\":\"0e635152-e501-4d4e-bb87-9f3fe05eb79a\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\"},{\"description\":\"Capture Card Source TR-04/2022-6\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:mux\",\"caps\":{},\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"Capture Card Source TR-04/2022-6\",\"id\":\"782fac41-17f6-4a21-8186-57ba63a1a8d3\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\"},{\"description\":\"Capture Card Source 2022-6 (No Refclock)\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:mux\",\"caps\":{},\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"Capture Card Source 2022-6 (No Refclock)\",\"id\":\"3ca37fce-c0cf-42a6-86ad-43635a53b5bb\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":null}]"
	data := []byte(dataStr)
	var test []Source
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

func TestQueryAPISourceIDGet(t *testing.T) {
	dataStr := "{\"description\":\"Capture Card Source Video\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{\"SourceDeviceType\":[\"HD Camera\"],\"host\":[\"host1\"],\"Location\":[\"Location 1\"]},\"caps\":{},\"version\":\"1441722516:851371645\",\"parents\":[],\"label\":\"CaptureCardSourceVideo\",\"id\":\"c23c6a65-8e91-4f6c-a484-046363dbca29\",\"device_id\":\"65fa8c20-890e-4b86-87b2-cfd9df91b7f8\",\"clock_name\":\"clk1\"}"
	data := []byte(dataStr)
	var test Source
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

func TestQueryAPISourcesGet(t *testing.T) {
	dataStr := "[{\"description\":\"Camera 1\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{\"SourceDeviceType\":[\"UHD Camera\"],\"host\":[\"host1\"],\"location\":[\"Location 1\"]},\"caps\":{},\"version\":\"1441724551:288670563\",\"parents\":[],\"label\":\"Camera 1\",\"id\":\"042a4126-0208-443d-bda6-833ffc27ed51\",\"device_id\":\"21a28338-fb2e-4df5-9b55-d58e6124bc9f\",\"clock_name\":\"clk0\"},{\"description\":\"Camera 2\",\"format\":\"urn:x-nmos:format:video\",\"tags\":{\"SourceDeviceType\":[\"HD Camera\"],\"host\":[\"host2\"],\"Location\":[\"Location 2\"]},\"caps\":{},\"version\":\"1441722516:851371645\",\"parents\":[],\"label\":\"CaptureCardSourceVideo\",\"id\":\"c23c6a65-8e91-4f6c-a484-046363dbca29\",\"device_id\":\"65fa8c20-890e-4b86-87b2-cfd9df91b7f8\",\"clock_name\":\"clk0\"},{\"description\":\"Audio 1\",\"format\":\"urn:x-nmos:format:audio\",\"tags\":{\"host\":[\"host3\"]},\"caps\":{},\"version\":\"1441719058:3226205\",\"parents\":[],\"label\":\"Audio 1\",\"id\":\"62cf8dd3-015b-49e3-84c1-1d866a7540bc\",\"device_id\":\"e7aa15c4-f793-4d43-a0cb-c638db6215ac\",\"clock_name\":\"clk0\",\"channels\":[{\"label\":\"Left Channel\",\"symbol\":\"L\"},{\"label\":\"Right Channel\",\"symbol\":\"R\"}]},{\"description\":\"Capture Card Source TR-04/2022-6\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:mux\",\"caps\":{},\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"Capture Card Source TR-04/2022-6\",\"id\":\"782fac41-17f6-4a21-8186-57ba63a1a8d3\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":\"clk0\"},{\"description\":\"Capture Card Source 2022-6 (No Refclock)\",\"tags\":{\"host\":[\"host1\"]},\"format\":\"urn:x-nmos:format:mux\",\"caps\":{},\"version\":\"1453880605:374934072\",\"parents\":[],\"label\":\"Capture Card Source 2022-6 (No Refclock)\",\"id\":\"3ca37fce-c0cf-42a6-86ad-43635a53b5bb\",\"device_id\":\"9126cc2f-4c26-4c9b-a6cd-93c4381c9be5\",\"clock_name\":null}]"
	data := []byte(dataStr)
	var test []Source
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
