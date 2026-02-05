package is05v1_1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IS05V1_1 struct {
	href string
}

func NewIS05V1_1(href string) *IS05V1_1 {
	v := IS05V1_1{
		href: href,
	}
	return &v
}

func (v *IS05V1_1) ConnectionPostBulkSenders(senders []BulkSenderPost) ([]BulkResponse, Error, error) {
	url := v.href + "/connection/v1.1/bulk/senders"
	requestBody, err := json.Marshal(senders)
	if err != nil {
		return nil, Error{}, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []BulkResponse
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusBadRequest:
		var errResponse Error
		errResponseParse := json.NewDecoder(response.Body).Decode(&errResponse)
		if errResponseParse != nil {
			return nil, Error{}, errResponseParse
		}
		return nil, errResponse, nil
	default:
		return nil, Error{}, fmt.Errorf("could not parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionPostBulkReceivers(receivers []BulkReceiverPost) ([]BulkResponse, Error, error) {
	url := v.href + "/connection/v1.1/bulk/receivers"
	requestBody, err := json.Marshal(receivers)
	if err != nil {
		return nil, Error{}, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []BulkResponse
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusBadRequest:
		var errResponse Error
		errResponseParse := json.NewDecoder(response.Body).Decode(&errResponse)
		if errResponseParse != nil {
			return nil, Error{}, errResponseParse
		}
		return nil, errResponse, nil
	default:
		return nil, Error{}, fmt.Errorf("could not parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenders() ([]string, error) {
	url := v.href + "/connection/v1.1/single/senders"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []string
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, err
		}
		// trim trailing slashes
		for i, str := range respParsed {
			respParsed[i] = strings.ReplaceAll(str, "/", "")
		}
		return respParsed, nil
	default:
		return nil, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenderConstraints(senderID string) ([]Constraints, Error, error) {
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/constraints"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []Constraints
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return nil, respParsed, nil
	default:
		return nil, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenderStaged(senderID string) (Sender, Error, error) {
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/staged"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Sender{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Sender{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed Sender
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	default:
		return Sender{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionPatchSenderStaged(senderID string, senderStaged Sender) (Sender, Error, error) {
	requestBody, err := json.Marshal(senderStaged)
	if err != nil {
		return Sender{}, Error{}, err
	}
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/staged"
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(requestBody))
	if err != nil {
		return Sender{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Sender{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK, http.StatusAccepted:
		var respParsed Sender
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusBadRequest:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	case http.StatusLocked:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	case http.StatusInternalServerError:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	default:
		return Sender{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenderActive(senderID string) (Sender, Error, error) {
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/active"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Sender{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Sender{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed Sender
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Sender{}, Error{}, err
		}
		return Sender{}, respParsed, nil
	default:
		return Sender{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenderTransportFile(senderID string) (string, Error, error) {
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/transportfile"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", Error{}, err
	}
	request.Header.Add("Accept", "application/json,application/sdp")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			return "", Error{}, err
		}
		return string(respBody), Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return "", respParsed, nil
	default:
		return "", Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetSenderTransportType(senderID string) (TransportType, Error, error) {
	url := v.href + "/connection/v1.1/single/senders/" + senderID + "/transporttype"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", Error{}, err
	}
	request.Header.Add("Accept", "application/json,application/sdp")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed TransportType
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return "", respParsed, nil
	case http.StatusConflict:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return "", respParsed, nil
	default:
		return "", Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetReceivers() ([]string, error) {
	url := v.href + "/connection/v1.1/single/receivers"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []string
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, err
		}
		return respParsed, nil
	default:
		return nil, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetReceiverConstraints(receiverID string) ([]Constraints, Error, error) {
	url := v.href + "/connection/v1.1/single/receivers/" + receiverID + "/constraints"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed []Constraints
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return nil, Error{}, err
		}
		return nil, respParsed, nil
	default:
		return nil, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetReceiverStaged(receiverID string) (Receiver, Error, error) {
	url := v.href + "/connection/v1.1/single/receivers/" + receiverID + "/staged"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed Receiver
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	default:
		return Receiver{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionPatchReceiverStaged(receiverID string, receiverStaged Receiver) (Receiver, Error, error) {
	requestBody, err := json.Marshal(receiverStaged)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	url := v.href + "/connection/v1.1/single/receivers/" + receiverID + "/staged"
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(requestBody))
	if err != nil {
		return Receiver{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK, http.StatusAccepted:
		var respParsed Receiver
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusBadRequest:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	case http.StatusLocked:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	case http.StatusInternalServerError:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	default:
		return Receiver{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetReceiverActive(receiverID string) (Receiver, Error, error) {
	url := v.href + "/connection/v1.1/single/receivers/" + receiverID + "/active"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Receiver{}, Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed Receiver
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return Receiver{}, Error{}, err
		}
		return Receiver{}, respParsed, nil
	default:
		return Receiver{}, Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}

func (v *IS05V1_1) ConnectionGetReceiverTransportType(receiverID string) (TransportType, Error, error) {
	url := v.href + "/connection/v1.1/single/receivers/" + receiverID + "/transporttype"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", Error{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", Error{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		var respParsed TransportType
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return respParsed, Error{}, nil
	case http.StatusNotFound:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return "", respParsed, nil
	case http.StatusConflict:
		var respParsed Error
		err = json.NewDecoder(response.Body).Decode(&respParsed)
		if err != nil {
			return "", Error{}, err
		}
		return "", respParsed, nil
	default:
		return "", Error{}, fmt.Errorf("unable to parse response code %d", response.StatusCode)
	}
}
