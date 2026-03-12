package is04v1_1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IS04V1_1 struct {
	href string
}

func NewIS04V1_1(href string) *IS04V1_1 {
	v := IS04V1_1{
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

func getPagingHRefs(response *http.Response) map[string]string {
	result := make(map[string]string)
	link := response.Header.Get("Link")
	linkParts := strings.Split(link, ",")
	for _, part := range linkParts {
		subparts := strings.Split(part, ";")
		key := subparts[1]
		key = strings.ReplaceAll(key, " ", "")
		key = key[5 : len(key)-1]
		val := subparts[0]
		val = strings.ReplaceAll(val, "<", "")
		val = strings.ReplaceAll(val, ">", "")
		val = strings.ReplaceAll(val, " ", "")
		result[key] = val
	}
	return result
}

// --------------------------------------------------
// -------------------- NODE API --------------------
// --------------------------------------------------

func (v *IS04V1_1) NodeGetSelf() (Node, error) {
	url := v.href + "/node/v1.1/self"
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

func (v *IS04V1_1) NodeGetSources() ([]Source, error) {
	url := v.href + "/node/v1.1/sources"
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

func (v *IS04V1_1) NodeGetSource(sourceID string) (Source, error) {
	url := v.href + "/node/v1.1/sources/" + sourceID
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

func (v *IS04V1_1) NodeGetFlows() ([]Flow, error) {
	url := v.href + "/node/v1.1/flows"
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

func (v *IS04V1_1) NodeGetFlow(flowID string) (Flow, error) {
	url := v.href + "/node/v1.1/flows/" + flowID
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

func (v *IS04V1_1) NodeGetDevices() ([]Device, error) {
	url := v.href + "/node/v1.1/devices"
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

func (v *IS04V1_1) NodeGetDevice(deviceID string) (Device, error) {
	url := v.href + "/node/v1.1/devices/" + deviceID
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

func (v *IS04V1_1) NodeGetSenders() ([]Sender, error) {
	url := v.href + "/node/v1.1/senders"
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

func (v *IS04V1_1) NodeGetSender(senderID string) (Sender, error) {
	url := v.href + "/node/v1.1/senders/" + senderID
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

func (v *IS04V1_1) NodeGetReceivers() ([]Receiver, error) {
	url := v.href + "/node/v1.1/receivers"
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

func (v *IS04V1_1) NodeGetReceiver(receiverID string) (Receiver, error) {
	url := v.href + "/node/v1.1/receivers/" + receiverID
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
func (v *IS04V1_1) NodePutReceiver(receiverID string, sender *Sender) error {
	url := v.href + "/node/v1.1/sources/" + receiverID + "/target"
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

func (v *IS04V1_1) QueryGetNodes() ([]Node, error) {
	url := v.href + "/query/v1.1/nodes"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Node, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Node](response)
			if err != nil {
				return []Node{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Node{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Node{}, err
			}
		}
		return result, nil
	default:
		return []Node{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetNode(nodeID string) (Node, error) {
	url := v.href + "/query/v1.1/nodes/" + nodeID
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

func (v *IS04V1_1) QueryGetSources() ([]Source, error) {
	url := v.href + "/query/v1.1/sources"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Source, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Source](response)
			if err != nil {
				return []Source{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Source{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Source{}, err
			}
		}
		return result, nil
	default:
		return []Source{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetSource(sourceID string) (Source, error) {
	url := v.href + "/query/v1.1/sources/" + sourceID
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

func (v *IS04V1_1) QueryGetFlows() ([]Flow, error) {
	url := v.href + "/query/v1.1/flows"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Flow, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Flow](response)
			if err != nil {
				return []Flow{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Flow{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Flow{}, err
			}
		}
		return result, nil
	default:
		return []Flow{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetFlow(flowID string) (Flow, error) {
	url := v.href + "/query/v1.1/flows/" + flowID
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

func (v *IS04V1_1) QueryGetDevices() ([]Device, error) {
	url := v.href + "/query/v1.1/devices"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Device, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Device](response)
			if err != nil {
				return []Device{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Device{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Device{}, err
			}
		}
		return result, nil
	default:
		return []Device{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetDevice(deviceID string) (Device, error) {
	url := v.href + "/query/v1.1/devices/" + deviceID
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

func (v *IS04V1_1) QueryGetSenders() ([]Sender, error) {
	url := v.href + "/query/v1.1/senders"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Sender, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Sender](response)
			if err != nil {
				return []Sender{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Sender{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Sender{}, err
			}
		}
		return result, nil
	default:
		return []Sender{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetSender(senderID string) (Sender, error) {
	url := v.href + "/query/v1.1/senders/" + senderID
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

func (v *IS04V1_1) QueryGetReceivers() ([]Receiver, error) {
	url := v.href + "/query/v1.1/receivers"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]Receiver, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]Receiver](response)
			if err != nil {
				return []Receiver{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []Receiver{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []Receiver{}, err
			}
		}
		return result, nil
	default:
		return []Receiver{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}

func (v *IS04V1_1) QueryGetReceiver(receiverID string) (Receiver, error) {
	url := v.href + "/query/v1.1/receivers/" + receiverID
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

func (v *IS04V1_1) QueryGetSubscriptions() ([]QueryAPISubscription, error) {
	url := v.href + "/query/v1.1/subscriptions"
	urlInitial := fmt.Sprintf("%s?paging.order=create&paging.since=0:0", url)
	request, err := http.NewRequest(http.MethodGet, urlInitial, nil)
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
		result := make([]QueryAPISubscription, 0)
		lastSince := "0:0"
		for true {
			respParsed, err := parseResponse[[]QueryAPISubscription](response)
			if err != nil {
				return []QueryAPISubscription{}, err
			}
			result = append(result, respParsed...)
			// Check paging
			limit := response.Header.Get("X-Paging-Until")
			if limit == "" || limit == lastSince {
				break
			}
			lastSince = limit
			next := fmt.Sprintf("%s?paging.order=create&paging.since=%s", url, limit)
			reqNew, err := http.NewRequest(http.MethodGet, next, nil)
			if err != nil {
				return []QueryAPISubscription{}, err
			}
			reqNew.Header.Add("Accept", "application/json")
			response, err = http.DefaultClient.Do(reqNew)
			if err != nil {
				return []QueryAPISubscription{}, err
			}
		}
		return result, nil
	default:
		return []QueryAPISubscription{}, fmt.Errorf("bad response code. got %d, wanted %d", response.StatusCode, http.StatusOK)
	}
}

func (v *IS04V1_1) QueryGetSubscription(subscriptionID string) (QueryAPISubscription, error) {
	url := v.href + "/query/v1.1/subscriptions/" + subscriptionID
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
func (v *IS04V1_1) QueryPostSubscription(subscription QueryAPISubscriptionPost) (QueryAPISubscription, error) {
	url := v.href + "/query/v1.1/subscriptions"
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

func (v *IS04V1_1) QueryDeleteSubscription(subscriptionID string) error {
	url := v.href + "/query/v1.1/subscriptions/" + subscriptionID
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

// --------------------------------------------------
// ---------------- REGISTRATION API ----------------
// --------------------------------------------------

// Show a Node's health (for debug use only)
func (v *IS04V1_1) RegistrationGetNodeHealth(nodeID string) (RegistrationHealth, error) {
	url := v.href + "/registration/v1.1/health/nodes/" + nodeID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RegistrationHealth{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return RegistrationHealth{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// Default
		respParsed, err := parseResponse[RegistrationHealth](response)
		if err != nil {
			return RegistrationHealth{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		// Returned when the POST request is incorrectly formatted or missing mandatory attributes
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return RegistrationHealth{}, err
		}
		return RegistrationHealth{}, respParsed.GetError()
	default:
		return RegistrationHealth{}, fmt.Errorf("bad response code. got %d", response.StatusCode)
	}
}
