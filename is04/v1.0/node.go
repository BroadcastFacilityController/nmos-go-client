package is04v1_0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type NodeAPI struct {
	hostname string
	port     uint16
}

func NewNodeAPI(hostname string, port uint16) *NodeAPI {
	v := NodeAPI{
		hostname: hostname,
		port:     port,
	}
	return &v
}

func (v *NodeAPI) GetBaseURL() string {
	return "http://" + v.hostname + ":" + strconv.FormatUint(uint64(v.port), 10) + "/x-nmos/node/v1.0"
}

func (v *NodeAPI) GetSelf() (Node, error) {
	url := v.GetBaseURL() + "/self"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Node{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Node{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Node{}, errors.New("invalid response code")
	}
	var respParsed Node
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Node{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetSources() ([]Source, error) {
	url := v.GetBaseURL() + "/sources"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Source{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Source{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []Source{}, errors.New("invalid response code")
	}
	var respParsed []Source
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return []Source{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetSource(sourceID string) (Source, error) {
	url := v.GetBaseURL() + "/sources/" + sourceID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Source{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Source{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Source{}, errors.New("invalid response code")
	}
	var respParsed Source
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Source{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetFlows() ([]Flow, error) {
	url := v.GetBaseURL() + "/flows"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Flow{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Flow{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []Flow{}, errors.New("invalid response code")
	}
	var respParsed []Flow
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return []Flow{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetFlow(flowID string) (Flow, error) {
	url := v.GetBaseURL() + "/flows/" + flowID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Flow{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Flow{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Flow{}, errors.New("invalid response code")
	}
	var respParsed Flow
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Flow{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetDevices() ([]Device, error) {
	url := v.GetBaseURL() + "/devices"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Device{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Device{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []Device{}, errors.New("invalid response code")
	}
	var respParsed []Device
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return []Device{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetDevice(deviceID string) (Device, error) {
	url := v.GetBaseURL() + "/devices/" + deviceID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Device{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Device{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Device{}, errors.New("invalid response code")
	}
	var respParsed Device
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Device{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetSenders() ([]Sender, error) {
	url := v.GetBaseURL() + "/senders"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Sender{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Sender{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []Sender{}, errors.New("invalid response code")
	}
	var respParsed []Sender
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return []Sender{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetSender(senderID string) (Sender, error) {
	url := v.GetBaseURL() + "/senders/" + senderID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Sender{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Sender{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Sender{}, errors.New("invalid response code")
	}
	var respParsed Sender
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Sender{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetReceivers() ([]Receiver, error) {
	url := v.GetBaseURL() + "/receivers"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Receiver{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Receiver{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []Receiver{}, errors.New("invalid response code")
	}
	var respParsed []Receiver
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return []Receiver{}, err
	}
	return respParsed, nil
}

func (v *NodeAPI) GetReceiver(receiverID string) (Receiver, error) {
	url := v.GetBaseURL() + "/receivers/" + receiverID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Receiver{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Receiver{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Receiver{}, errors.New("invalid response code")
	}
	var respParsed Receiver
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return Receiver{}, err
	}
	return respParsed, nil
}

// Patch a sender to a receiver. Sending a nil sender will unpatch.
func (v *NodeAPI) PutReceiverSubscription(receiverID string, sender *Sender) error {
	url := v.GetBaseURL() + "/receivers/" + receiverID + "/target"
	data, err := json.Marshal(sender)
	if err != nil {
		return err
	}
	if sender == nil {
		data = []byte("{}")
	}
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("received bad status response code. Expected 202 got %d. %s", response.StatusCode, response.Body)
	}
	return nil
}
