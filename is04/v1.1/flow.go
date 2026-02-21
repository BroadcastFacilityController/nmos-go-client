package is04v1_1

import (
	"encoding/json"
	"fmt"
)

// Used purely for json unmarshalling and finding the listed type.
// Only one pointer will be non-nil.
// "Type" property will list the type for easy switch statement handling.
type Flow struct {
	Type           FlowType
	FlowVideoRaw   *FlowVideoRaw
	FlowVideoCoded *FlowVideoCoded
	FlowAudioRaw   *FlowAudioRaw
	FlowAudioCoded *FlowAudioCoded
	FlowData       *FlowData
	FlowSDIANCData *FlowSDIANCData
	FlowMux        *FlowMux
}

type FlowType string

const (
	FLOW_TYPE_VIDEO_RAW   FlowType = "flow_video_raw"
	FLOW_TYPE_VIDEO_CODED FlowType = "flow_video_coded"
	FLOW_TYPE_AUDIO_RAW   FlowType = "flow_audio_raw"
	FLOW_TYPE_AUDIO_CODED FlowType = "flow_audio_coded"
	FLOW_TYPE_DATA        FlowType = "flow_data"
	FLOW_TYPE_SDIANC_DATA FlowType = "flow_sdianc_data"
	FLOW_TYPE_MUX         FlowType = "flow_mux"
)

func (f *Flow) UnmarshalJSON(data []byte) error {
	// Unmarshal to a map[string]interface{} to test
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
		// Video
		mediaType, mediaType_ok := dataTest["media_type"].(IANAMediaType)
		if !mediaType_ok {
			return fmt.Errorf("unable to parse format for data: %s", string(data))
		}
		switch mediaType {
		case IANA_MEDIA_TYPE_VIDEO_RAW:
			// Raw
			err = json.Unmarshal(data, f.FlowVideoRaw)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_VIDEO_RAW
			return nil
		default:
			// Coded
			err = json.Unmarshal(data, f.FlowVideoCoded)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_VIDEO_CODED
			return nil
		}
	case FORMAT_AUDIO:
		mediaType, mediaType_ok := dataTest["media_type"].(IANAMediaType)
		if !mediaType_ok {
			return fmt.Errorf("unable to parse format for data: %s", string(data))
		}
		switch mediaType {
		case IANA_MEDIA_TYPE_AUDIO_L24:
			// Raw
			err = json.Unmarshal(data, f.FlowAudioRaw)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_AUDIO_RAW
			return nil
		case IANA_MEDIA_TYPE_AUDIO_L20:
			// Raw
			err = json.Unmarshal(data, f.FlowAudioRaw)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_AUDIO_RAW
			return nil
		case IANA_MEDIA_TYPE_AUDIO_L16:
			// Raw
			err = json.Unmarshal(data, f.FlowAudioRaw)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_AUDIO_RAW
			return nil
		case IANA_MEDIA_TYPE_AUDIO_L8:
			// Raw
			err = json.Unmarshal(data, f.FlowAudioRaw)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_AUDIO_RAW
			return nil
		default:
			// Coded
			err = json.Unmarshal(data, f.FlowAudioCoded)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_AUDIO_CODED
			return nil
		}
	case FORMAT_DATA:
		mediaType, mediaType_ok := dataTest["media_type"].(IANAMediaType)
		if !mediaType_ok {
			return fmt.Errorf("unable to parse format for data: %s", string(data))
		}
		switch mediaType {
		case IANA_MEDIA_TYPE_SMPTE291:
			// sdi_anc
			err = json.Unmarshal(data, f.FlowSDIANCData)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_SDIANC_DATA
			return nil
		default:
			// data
			err = json.Unmarshal(data, f.FlowData)
			if err != nil {
				return err
			}
			f.Type = FLOW_TYPE_DATA
			return nil
		}
	case FORMAT_MUX:
		err = json.Unmarshal(data, f.FlowMux)
		if err != nil {
			return err
		}
		f.Type = FLOW_TYPE_MUX
		return nil
	default:
		return fmt.Errorf("unable to parse format for data: %s", string(data))
	}
}

func (f *Flow) MarshalJSON() ([]byte, error) {
	switch f.Type {
	case FLOW_TYPE_VIDEO_RAW:
		return json.Marshal(f.FlowVideoRaw)
	case FLOW_TYPE_VIDEO_CODED:
		return json.Marshal(f.FlowVideoCoded)
	case FLOW_TYPE_AUDIO_RAW:
		return json.Marshal(f.FlowAudioRaw)
	case FLOW_TYPE_AUDIO_CODED:
		return json.Marshal(f.FlowAudioCoded)
	case FLOW_TYPE_DATA:
		return json.Marshal(f.FlowData)
	case FLOW_TYPE_SDIANC_DATA:
		return json.Marshal(f.FlowSDIANCData)
	case FLOW_TYPE_MUX:
		return json.Marshal(f.FlowMux)
	default:
		return nil, fmt.Errorf("could not marshal flow %v", *f)
	}
}
