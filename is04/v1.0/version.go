package is04v1_0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IS04V1_0 struct {
	href string
}

func NewIS04V1_0(href string) *IS04V1_0 {
	v := IS04V1_0{
		href: href,
	}
	return &v
}

func parseResponse[T any](response *http.Response) (T, error) {
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return *new(T), err
	}
	var responseParsed T
	err = json.Unmarshal(responseBody, &responseParsed)
	if err != nil {
		return *new(T), err
	}
	return responseParsed, nil
}

// --------------------------------------------------
// -------------------- NODE API --------------------
// --------------------------------------------------

func (v *IS04V1_0) NodeGetSelf() (Node, error) {
	url := v.href + "/node/v1.0/self"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Node{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Node{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Node](response)
		if err != nil {
			return Node{}, err
		}
		return respParsed, nil
	default:
		return Node{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetSources() ([]Source, error) {
	url := v.href + "/node/v1.0/sources"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Source{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Source{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Source](response)
		if err != nil {
			return []Source{}, err
		}
		return respParsed, nil
	default:
		return []Source{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetSource(sourceID string) (Source, error) {
	url := v.href + "/node/v1.0/sources/" + sourceID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Source{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Source{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Source](response)
		if err != nil {
			return Source{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Source ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Source{}, err
		}
		return Source{}, respParsed.GetError()
	default:
		return Source{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetFlows() ([]Flow, error) {
	url := v.href + "/node/v1.0/flows"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Flow{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Flow{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Flow](response)
		if err != nil {
			return []Flow{}, err
		}
		return respParsed, nil
	default:
		return []Flow{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetFlow(flowID string) (Flow, error) {
	url := v.href + "/node/v1.0/flows/" + flowID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Flow{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Flow{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Flow](response)
		if err != nil {
			return Flow{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Flow{}, err
		}
		return Flow{}, respParsed.GetError()
	default:
		return Flow{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetDevices() ([]Device, error) {
	url := v.href + "/node/v1.0/devices"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Device{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Device{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Device](response)
		if err != nil {
			return []Device{}, err
		}
		return respParsed, nil
	default:
		return []Device{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetDevice(deviceID string) (Device, error) {
	url := v.href + "/node/v1.0/devices/" + deviceID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Device{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Device{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Device](response)
		if err != nil {
			return Device{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Device{}, err
		}
		return Device{}, respParsed.GetError()
	default:
		return Device{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetSenders() ([]Sender, error) {
	url := v.href + "/node/v1.0/senders"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Sender{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Sender{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Sender](response)
		if err != nil {
			return []Sender{}, err
		}
		return respParsed, nil
	default:
		return []Sender{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetSender(senderID string) (Sender, error) {
	url := v.href + "/node/v1.0/senders/" + senderID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Sender{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Sender{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Sender](response)
		if err != nil {
			return Sender{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Sender{}, err
		}
		return Sender{}, respParsed.GetError()
	default:
		return Sender{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetReceivers() ([]Receiver, error) {
	url := v.href + "/node/v1.0/receivers"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Receiver{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Receiver{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Receiver](response)
		if err != nil {
			return []Receiver{}, err
		}
		return respParsed, nil
	default:
		return []Receiver{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) NodeGetReceiver(receiverID string) (Receiver, error) {
	url := v.href + "/node/v1.0/receivers/" + receiverID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Receiver{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Receiver{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Receiver](response)
		if err != nil {
			return Receiver{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Receiver{}, err
		}
		return Receiver{}, respParsed.GetError()
	default:
		return Receiver{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

// Patches a sender to a receiver. Passing nil as sender agument unpatches
func (v *IS04V1_0) NodePutReceiver(receiverID string, sender *Sender) error {
	url := v.href + "/node/v1.0/sources/" + receiverID + "/target"
	requestBody, err := json.Marshal(sender)
	if err != nil {
		return err
	}
	if sender == nil {
		requestBody = []byte("{}")
	}
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case http.StatusAccepted:
		// Default
		return nil
	case http.StatusBadRequest:
		// Returned when the PUT request is incorrectly formatted or missing mandatory attributes
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return err
		}
		return respParsed.GetError()
	case http.StatusNotFound:
		// Returned when the requested Receiver ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return err
		}
		return respParsed.GetError()
	default:
		return fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

// --------------------------------------------------
// -------------------- QUERY API -------------------
// --------------------------------------------------

func (v *IS04V1_0) QueryGetNodes() ([]Node, error) {
	url := v.href + "/query/v1.0/nodes"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Node{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Node{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Node](response)
		if err != nil {
			return []Node{}, err
		}
		return respParsed, nil
	default:
		return []Node{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetNode(nodeID string) (Node, error) {
	url := v.href + "/query/v1.0/nodes/" + nodeID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Node{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Node{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Node](response)
		if err != nil {
			return Node{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Node{}, err
		}
		return Node{}, respParsed.GetError()
	default:
		return Node{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetSources() ([]Source, error) {
	url := v.href + "/query/v1.0/sources"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Source{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Source{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Source](response)
		if err != nil {
			return []Source{}, err
		}
		return respParsed, nil
	default:
		return []Source{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetSource(sourceID string) (Source, error) {
	url := v.href + "/query/v1.0/sources/" + sourceID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Source{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Source{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Source](response)
		if err != nil {
			return Source{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Source{}, err
		}
		return Source{}, respParsed.GetError()
	default:
		return Source{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetFlows() ([]Flow, error) {
	url := v.href + "/query/v1.0/flows"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Flow{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Flow{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Flow](response)
		if err != nil {
			return []Flow{}, err
		}
		return respParsed, nil
	default:
		return []Flow{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetFlow(flowID string) (Flow, error) {
	url := v.href + "/query/v1.0/flows/" + flowID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Flow{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Flow{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Flow](response)
		if err != nil {
			return Flow{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Flow{}, err
		}
		return Flow{}, respParsed.GetError()
	default:
		return Flow{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetDevices() ([]Device, error) {
	url := v.href + "/query/v1.0/devices"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Device{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Device{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Device](response)
		if err != nil {
			return []Device{}, err
		}
		return respParsed, nil
	default:
		return []Device{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetDevice(deviceID string) (Device, error) {
	url := v.href + "/query/v1.0/devices/" + deviceID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Device{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Device{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Device](response)
		if err != nil {
			return Device{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Device{}, err
		}
		return Device{}, respParsed.GetError()
	default:
		return Device{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetSenders() ([]Sender, error) {
	url := v.href + "/query/v1.0/senders"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Sender{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Sender{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Sender](response)
		if err != nil {
			return []Sender{}, err
		}
		return respParsed, nil
	default:
		return []Sender{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetSender(senderID string) (Sender, error) {
	url := v.href + "/query/v1.0/senders/" + senderID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Sender{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Sender{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Sender](response)
		if err != nil {
			return Sender{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Sender{}, err
		}
		return Sender{}, respParsed.GetError()
	default:
		return Sender{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetReceivers() ([]Receiver, error) {
	url := v.href + "/query/v1.0/receivers"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Receiver{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []Receiver{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]Receiver](response)
		if err != nil {
			return []Receiver{}, err
		}
		return respParsed, nil
	default:
		return []Receiver{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetReceiver(receiverID string) (Receiver, error) {
	url := v.href + "/query/v1.0/receivers/" + receiverID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Receiver{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Receiver{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[Receiver](response)
		if err != nil {
			return Receiver{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the requested Flow ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Receiver{}, err
		}
		return Receiver{}, respParsed.GetError()
	default:
		return Receiver{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryGetSubscriptions() ([]QueryAPISubscription, error) {
	url := v.href + "/query/v1.0/subscriptions"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []QueryAPISubscription{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return []QueryAPISubscription{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[[]QueryAPISubscription](response)
		if err != nil {
			return []QueryAPISubscription{}, err
		}
		return respParsed, nil
	default:
		return []QueryAPISubscription{}, fmt.Errorf("bad response code. got %d, wanted %d", response.StatusCode, http.StatusOK)
	}
}

func (v *IS04V1_0) QueryGetSubscription(subscriptionID string) (QueryAPISubscription, error) {
	url := v.href + "/query/v1.0/subscriptions/" + subscriptionID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return QueryAPISubscription{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return QueryAPISubscription{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[QueryAPISubscription](response)
		if err != nil {
			return QueryAPISubscription{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the POST request is incorrectly formatted or missing mandatory attributes
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return QueryAPISubscription{}, err
		}
		return QueryAPISubscription{}, respParsed.GetError()
	default:
		return QueryAPISubscription{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

// Open a new websocket connection
func (v *IS04V1_0) QueryPostSubscription(subscription QueryAPISubscriptionPost) (QueryAPISubscription, error) {
	url := v.href + "/query/v1.0/subscriptions"
	requestBody, err := json.Marshal(subscription)
	if err != nil {
		return QueryAPISubscription{}, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return QueryAPISubscription{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return QueryAPISubscription{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[QueryAPISubscription](response)
		if err != nil {
			return QueryAPISubscription{}, err
		}
		return respParsed, nil
	case http.StatusCreated:
		// Default
		respParsed, err := parseResponse[QueryAPISubscription](response)
		if err != nil {
			return QueryAPISubscription{}, err
		}
		return respParsed, nil
	case http.StatusBadRequest:
		// Returned when the POST request is incorrectly formatted or missing mandatory attributes
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return QueryAPISubscription{}, err
		}
		return QueryAPISubscription{}, respParsed.GetError()
	default:
		return QueryAPISubscription{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_0) QueryDeleteSubscription(subscriptionID string) error {
	url := v.href + "/query/v1.0/subscriptions/" + subscriptionID
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case http.StatusNoContent:
		// Default
		return nil
	case http.StatusForbidden:
		// Returned when a DELETE request is made against a non-persistent subscription which is managed by the Query API and cannot be deleted
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return err
		}
		return respParsed.GetError()
	case http.StatusNotFound:
		// Returned when the requested Subscription ID does not exist
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return err
		}
		return respParsed.GetError()
	default:
		return fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}
