package is07v1_0

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IS07V1_0 struct {
	href string
}

func NewIS07V1_0(href string) *IS07V1_0 {
	v := IS07V1_0{
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

// List of the Event & Tally compatible sources
func (v *IS07V1_0) GetSources() ([]string, error) {
	url := v.href + "/events/v1.0/sources"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	respParsed, err := parseResponse[[]string](response)
	if err != nil {
		return nil, err
	}
	return respParsed, nil
}

// JSON containing detailed type information
func (v *IS07V1_0) GetSourceType(sourceID string) (Type, error) {
	url := v.href + "/events/v1.0/sources/" + sourceID + "/type"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Type{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Type{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[Type](response)
		if err != nil {
			return Type{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Type{}, err
		}
		return Type{}, respParsed.GetError()
	default:
		return Type{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// JSON containing the last state change notification sent over the transport protocols
func (v *IS07V1_0) GetSourceState(sourceID string) (Event, error) {
	url := v.href + "/events/v1.0/sources/" + sourceID + "/state"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Event{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Event{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[Event](response)
		if err != nil {
			return Event{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return Event{}, err
		}
		return Event{}, respParsed.GetError()
	default:
		return Event{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}
