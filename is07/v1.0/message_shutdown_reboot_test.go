package is07v1_0

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageReboot(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"6cbd0441-7882-44cd-9557-842243a0d618\",\"flow_id\":\"0febba81-d924-4bbf-b1d9-83a11fbabc61\"},\"timing\":{\"creation_timestamp\":\"1531680501:280709600\"},\"message_type\":\"reboot\"}"
	data := []byte(dataStr)
	var test MessageShutdownReboot
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}

func TestMessageShutdown(t *testing.T) {
	dataStr := "{\"identity\":{\"source_id\":\"6cbd0441-7882-44cd-9557-842243a0d618\",\"flow_id\":\"0febba81-d924-4bbf-b1d9-83a11fbabc61\"},\"timing\":{\"creation_timestamp\":\"1531680501:280709600\"},\"message_type\":\"shutdown\"}"
	data := []byte(dataStr)
	var test MessageShutdownReboot
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err.Error())
	}
	result, err := json.Marshal(test)
	if err != nil {
		t.Errorf("error marshalling json: %s", err.Error())
	}
	assert.JSONEq(t, string(data), string(result), "result does not match input, result: \n%s\n\nwanted: \n%s", string(data), string(result))
}
