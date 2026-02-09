package is08v1_0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/guregu/null/v6"
)

type IS08V1_0 struct {
	href string
}

func NewIS08V1_0(href string) *IS08V1_0 {
	v := IS08V1_0{
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

// List all inputs available
func (v *IS08V1_0) GetInputs() ([]string, error) {
	url := v.href + "/channelmapping/v1.0/inputs"
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
	// remove trailing slashes
	for i := range respParsed {
		respParsed[i] = strings.ReplaceAll(respParsed[i], "/", "")
	}
	return respParsed, nil
}

// Get information about the specified Input, such as a human-readable name and description
func (v *IS08V1_0) GetInputProperties(inputID string) (InputProperties, error) {
	url := v.href + "/channelmapping/v1.0/inputs/" + inputID + "/properties"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return InputProperties{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return InputProperties{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[InputProperties](response)
		if err != nil {
			return InputProperties{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return InputProperties{}, err
		}
		return InputProperties{}, respParsed.GetError()
	default:
		return InputProperties{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get information about the origin of the audio associated with the specified Input
func (v *IS08V1_0) GetInputParent(inputID string) (InputParentResponse, error) {
	url := v.href + "/channelmapping/v1.0/inputs/" + inputID + "/parent"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return InputParentResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return InputParentResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[InputParentResponse](response)
		if err != nil {
			return InputParentResponse{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return InputParentResponse{}, err
		}
		return InputParentResponse{}, respParsed.GetError()
	default:
		return InputParentResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get information about the specified Input's channels, such as human-readable labels
func (v *IS08V1_0) GetInputChannels(inputID string) ([]InputChannelsResponse, error) {
	url := v.href + "/channelmapping/v1.0/inputs/" + inputID + "/channels"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[[]InputChannelsResponse](response)
		if err != nil {
			return nil, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return nil, err
		}
		return nil, respParsed.GetError()
	default:
		return nil, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get information about the specified Input's capabilities, such as routing constraints
func (v *IS08V1_0) GetInputCaps(inputID string) (InputCapsResponse, error) {
	url := v.href + "/channelmapping/v1.0/inputs/" + inputID + "/caps"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return InputCapsResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return InputCapsResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[InputCapsResponse](response)
		if err != nil {
			return InputCapsResponse{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return InputCapsResponse{}, err
		}
		return InputCapsResponse{}, respParsed.GetError()
	default:
		return InputCapsResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// List all outputs available
func (v *IS08V1_0) GetOutputs() ([]string, error) {
	url := v.href + "/channelmapping/v1.0/outputs"
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
	// remove trailing slashes
	for i := range respParsed {
		respParsed[i] = strings.ReplaceAll(respParsed[i], "/", "")
	}
	return respParsed, nil
}

// Get information about the specified Output, such as a human-readable name and description
func (v *IS08V1_0) GetOutputProperties(outputID string) (OutputProperties, error) {
	url := v.href + "/channelmapping/v1.0/outputs/" + outputID + "/properties"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return OutputProperties{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return OutputProperties{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[OutputProperties](response)
		if err != nil {
			return OutputProperties{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return OutputProperties{}, err
		}
		return OutputProperties{}, respParsed.GetError()
	default:
		return OutputProperties{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get the ID of the Source associated with the specified Output
func (v *IS08V1_0) GetOutputSourceID(outputID string) (null.String, error) {
	url := v.href + "/channelmapping/v1.0/outputs/" + outputID + "/sourceid"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return null.String{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return null.String{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[null.String](response)
		if err != nil {
			return null.String{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return null.String{}, err
		}
		return null.String{}, respParsed.GetError()
	default:
		return null.String{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get information about the specified Output's channels, such as human-readable labels
func (v *IS08V1_0) GetOutputChannels(outputID string) ([]OutputChannelsResponse, error) {
	url := v.href + "/channelmapping/v1.0/outputs/" + outputID + "/channels"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[[]OutputChannelsResponse](response)
		if err != nil {
			return nil, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return nil, err
		}
		return nil, respParsed.GetError()
	default:
		return nil, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get information about the specified Output's capabilities, such as routing constraints
func (v *IS08V1_0) GetOutputCaps(outputID string) (OutputCapsResponse, error) {
	url := v.href + "/channelmapping/v1.0/outputs/" + outputID + "/caps"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return OutputCapsResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return OutputCapsResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[OutputCapsResponse](response)
		if err != nil {
			return OutputCapsResponse{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return OutputCapsResponse{}, err
		}
		return OutputCapsResponse{}, respParsed.GetError()
	default:
		return OutputCapsResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// List pending scheduled activations
func (v *IS08V1_0) GetMapActivations() (MapActivationsGetResponse, error) {
	url := v.href + "/channelmapping/v1.0/map/activations"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MapActivationsGetResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MapActivationsGetResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[MapActivationsGetResponse](response)
		if err != nil {
			return MapActivationsGetResponse{}, err
		}
		return respParsed, nil
	default:
		return MapActivationsGetResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Create a new activation
func (v *IS08V1_0) PostMapActivations(activation MapActivationsPostRequest) (MapActivationsPostResponse, error) {
	url := v.href + "/channelmapping/v1.0/map/activations"
	requestBody, err := json.Marshal(activation)
	if err != nil {
		return MapActivationsPostResponse{}, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return MapActivationsPostResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MapActivationsPostResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		// A 200 response is returned when a request for an immediate activation is accepted.
		respParsed, err := parseResponse[MapActivationsPostResponse](response)
		if err != nil {
			return MapActivationsPostResponse{}, err
		}
		return respParsed, nil
	case http.StatusAccepted:
		// A 202 response is returned when a request for a scheduled activation is accepted, to indicate that while the request itself was acceptable it has not yet been acted upon by the API.
		respParsed, err := parseResponse[MapActivationsPostResponse](response)
		if err != nil {
			return MapActivationsPostResponse{}, err
		}
		return respParsed, nil
	case http.StatusBadRequest:
		// A 400 response is returned, for example, when the request did not meet the schema.
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return MapActivationsPostResponse{}, err
		}
		return MapActivationsPostResponse{}, respParsed.GetError()
	case http.StatusLocked:
		// A 423 response is returned when a conflicting activation has been scheduled.
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return MapActivationsPostResponse{}, err
		}
		return MapActivationsPostResponse{}, respParsed.GetError()
	default:
		return MapActivationsPostResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get a signle pending scheduled activation
func (v *IS08V1_0) GetMapActivation(activationID string) (MapActivationsGetResponseElement, error) {
	url := v.href + "/channelmapping/v1.0/map/activations/" + activationID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MapActivationsGetResponseElement{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MapActivationsGetResponseElement{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[MapActivationsGetResponseElement](response)
		if err != nil {
			return MapActivationsGetResponseElement{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return MapActivationsGetResponseElement{}, err
		}
		return MapActivationsGetResponseElement{}, respParsed.GetError()
	default:
		return MapActivationsGetResponseElement{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Cancel a pending scheduled activation
func (v *IS08V1_0) DeleteMapActivation(activationID string) error {
	url := v.href + "/channelmapping/v1.0/map/activations/" + activationID
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return err
		}
		return respParsed.GetError()
	default:
		return fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get a view of the entire active map
func (v *IS08V1_0) GetMapActive() (MapActiveResponse, error) {
	url := v.href + "/channelmapping/v1.0/map/active"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MapActiveResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MapActiveResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[MapActiveResponse](response)
		if err != nil {
			return MapActiveResponse{}, err
		}
		return respParsed, nil
	default:
		return MapActiveResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get the active map for the specified output. This resource allows a controller to fetch only the section of the map related to the specified output, without the need to fetch the entire map.
func (v *IS08V1_0) GetMapActiveByOutput(outputID string) (MapActiveResponse, error) {
	url := v.href + "/channelmapping/v1.0/map/active/" + outputID
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MapActiveResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MapActiveResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[MapActiveResponse](response)
		if err != nil {
			return MapActiveResponse{}, err
		}
		return respParsed, nil
	case http.StatusNotFound:
		respParsed, err := parseResponse[Error](response)
		if err != nil {
			return MapActiveResponse{}, err
		}
		return MapActiveResponse{}, respParsed.GetError()
	default:
		return MapActiveResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}

// Get a view on all Inputs and Outputs
func (v *IS08V1_0) GetIO() (IOResponse, error) {
	url := v.href + "/channelmapping/v1.0/io"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return IOResponse{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return IOResponse{}, err
	}
	switch response.StatusCode {
	case http.StatusOK:
		respParsed, err := parseResponse[IOResponse](response)
		if err != nil {
			return IOResponse{}, err
		}
		return respParsed, nil
	default:
		return IOResponse{}, fmt.Errorf("unable to handle response code %d", response.StatusCode)
	}
}
