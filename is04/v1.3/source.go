package is04v1_3

import (
	"encoding/json"
	"fmt"
)

// Describes a Source
// Used for JSON marshalling
// Type indicates which pointer is non-nil
type Source struct {
	Type          SourceType
	SourceGeneric *SourceGeneric
	SourceData    *SourceData
	SourceAudio   *SourceAudio
}

type SourceType string

const (
	SOURCE_TYPE_GENERIC SourceType = "source_generic"
	SOURCE_TYPE_AUDIO   SourceType = "source_audio"
	SOURCE_TYPE_DATA    SourceType = "source_data"
)

func (s *Source) UnmarshalJSON(data []byte) error {
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
	case FORMAT_AUDIO:
		s.Type = SOURCE_TYPE_AUDIO
		var srcTemp SourceAudio
		err = json.Unmarshal(data, &srcTemp)
		if err != nil {
			return err
		}
		s.SourceAudio = &srcTemp
		return nil
	case FORMAT_VIDEO:
		s.Type = SOURCE_TYPE_GENERIC
		var srcTemp SourceGeneric
		err = json.Unmarshal(data, &srcTemp)
		if err != nil {
			return err
		}
		s.SourceGeneric = &srcTemp
		return nil
	case FORMAT_DATA:
		s.Type = SOURCE_TYPE_DATA
		var srcTemp SourceData
		err = json.Unmarshal(data, &srcTemp)
		if err != nil {
			return err
		}
		s.SourceData = &srcTemp
		return nil
	case FORMAT_MUX:
		s.Type = SOURCE_TYPE_GENERIC
		var srcTemp SourceGeneric
		err = json.Unmarshal(data, &srcTemp)
		if err != nil {
			return err
		}
		s.SourceGeneric = &srcTemp
		return nil
	default:
		return fmt.Errorf("unable to parse format for data: %s", string(data))
	}
}

func (s *Source) MarshalJSON() ([]byte, error) {
	switch s.Type {
	case SOURCE_TYPE_AUDIO:
		return json.Marshal(s.SourceAudio)
	case SOURCE_TYPE_DATA:
		return json.Marshal(s.SourceData)
	case SOURCE_TYPE_GENERIC:
		return json.Marshal(s.SourceGeneric)
	default:
		return nil, fmt.Errorf("could not marshal source %v", *s)
	}
}

func (s Source) String() string {
	switch s.Type {
	case SOURCE_TYPE_AUDIO:
		return fmt.Sprint(*s.SourceAudio)
	case SOURCE_TYPE_DATA:
		return fmt.Sprint(*s.SourceData)
	case SOURCE_TYPE_GENERIC:
		return fmt.Sprint(*s.SourceGeneric)
	default:
		return ""
	}
}
