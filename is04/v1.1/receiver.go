package is04v1_1

import (
	"encoding/json"
	"fmt"
)

// Describes a receiver.
// Used for JSON unmarshalling.
// Type points to the one pointer which is non-nil.
type Receiver struct {
	Type          ReceiverType // Type of the receiver
	ReceiverVideo *ReceiverVideo
	ReceiverAudio *ReceiverAudio
	ReceiverData  *ReceiverData
	ReceiverMux   *ReceiverMux
}

type ReceiverType string

const (
	RECEIVER_TYPE_VIDEO ReceiverType = "receiver_video"
	RECEIVER_TYPE_AUDIO ReceiverType = "receiver_audio"
	RECEIVER_TYPE_DATA  ReceiverType = "receiver_data"
	RECEIVER_TYPE_MUX   ReceiverType = "receiver_mux"
)

func (r *Receiver) UnmarshalJSON(data []byte) error {
	// Unmarshall to an object to test again
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	format, format_ok := dataTest["format"].(string)
	if !format_ok {
		return fmt.Errorf("unable to parse format for data: %s", string(data))
	}
	switch FormatURI(format) {
	case FORMAT_VIDEO:
		r.Type = RECEIVER_TYPE_VIDEO
		var rxTemp ReceiverVideo
		err = json.Unmarshal(data, &rxTemp)
		if err != nil {
			return err
		}
		r.ReceiverVideo = &rxTemp
		return nil
	case FORMAT_AUDIO:
		r.Type = RECEIVER_TYPE_AUDIO
		var rxTemp ReceiverAudio
		err = json.Unmarshal(data, &rxTemp)
		if err != nil {
			return err
		}
		r.ReceiverAudio = &rxTemp
		return nil
	case FORMAT_DATA:
		r.Type = RECEIVER_TYPE_DATA
		var rxTemp ReceiverData
		err = json.Unmarshal(data, &rxTemp)
		if err != nil {
			return err
		}
		r.ReceiverData = &rxTemp
		return nil
	case FORMAT_MUX:
		r.Type = RECEIVER_TYPE_MUX
		var rxTemp ReceiverMux
		err = json.Unmarshal(data, &rxTemp)
		if err != nil {
			return err
		}
		r.ReceiverMux = &rxTemp
		return nil
	default:
		return fmt.Errorf("unable to parse format for data: %s", string(data))
	}
}

func (r *Receiver) MarshalJSON() ([]byte, error) {
	switch r.Type {
	case RECEIVER_TYPE_VIDEO:
		return json.Marshal(r.ReceiverVideo)
	case RECEIVER_TYPE_AUDIO:
		return json.Marshal(r.ReceiverAudio)
	case RECEIVER_TYPE_DATA:
		return json.Marshal(r.ReceiverData)
	case RECEIVER_TYPE_MUX:
		return json.Marshal(r.ReceiverMux)
	default:
		return nil, fmt.Errorf("unable to marshal receiver %v", *r)
	}
}

func (r Receiver) String() string {
	switch r.Type {
	case RECEIVER_TYPE_VIDEO:
		return fmt.Sprint(*r.ReceiverVideo)
	case RECEIVER_TYPE_AUDIO:
		return fmt.Sprint(*r.ReceiverAudio)
	case RECEIVER_TYPE_DATA:
		return fmt.Sprint(*r.ReceiverData)
	case RECEIVER_TYPE_MUX:
		return fmt.Sprint(*r.ReceiverMux)
	default:
		return ""
	}
}
