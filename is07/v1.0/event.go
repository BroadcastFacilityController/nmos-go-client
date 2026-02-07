package is07v1_0

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type EventType string

const (
	EVENT_TYPE_BOOLEAN EventType = "boolean"
	EVENT_TYPE_NUMBER  EventType = "number"
	EVENT_TYPE_STRING  EventType = "string"
	EVENT_TYPE_OBJECT  EventType = "object"
)

type Event struct {
	Type        EventType
	Boolean     *EventBoolean
	Number      *EventNumber
	StringEvent *EventString
	Object      *EventObject
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	eventType, eventType_ok := dataTest["event_type"].(string)
	if !eventType_ok {
		return fmt.Errorf("unable to parse data %v", dataTest)
	}
	boolMatch, err := regexp.MatchString("^boolean(\\/[^\\s\\/]+)*$", eventType)
	if err != nil {
		return err
	}
	if boolMatch {
		var respParsed EventBoolean
		err := json.Unmarshal(data, &respParsed)
		if err != nil {
			return err
		}
		e.Boolean = &respParsed
		e.Type = EVENT_TYPE_BOOLEAN
		return nil
	}
	numberMatch, err := regexp.MatchString("^number(\\/[^\\s\\/]+)*$", eventType)
	if err != nil {
		return err
	}
	if numberMatch {
		var respParsed EventNumber
		err := json.Unmarshal(data, &respParsed)
		if err != nil {
			return err
		}
		e.Number = &respParsed
		e.Type = EVENT_TYPE_NUMBER
		return nil
	}
	stringMatch, err := regexp.MatchString("^string(\\/[^\\s\\/]+)*$", eventType)
	if err != nil {
		return err
	}
	if stringMatch {
		var respParsed EventString
		err := json.Unmarshal(data, &respParsed)
		if err != nil {
			return err
		}
		e.StringEvent = &respParsed
		e.Type = EVENT_TYPE_STRING
		return nil
	}
	objectMatch, err := regexp.MatchString("^object(\\/[^\\s\\/]+)*$", eventType)
	if err != nil {
		return err
	}
	if objectMatch {
		var respParsed EventObject
		err := json.Unmarshal(data, &respParsed)
		if err != nil {
			return err
		}
		e.Object = &respParsed
		e.Type = EVENT_TYPE_OBJECT
		return nil
	}
	return fmt.Errorf("unable to match event type %s", eventType)
}

func (e *Event) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case EVENT_TYPE_BOOLEAN:
		return json.Marshal(e.Boolean)
	case EVENT_TYPE_NUMBER:
		return json.Marshal(e.Number)
	case EVENT_TYPE_STRING:
		return json.Marshal(e.StringEvent)
	case EVENT_TYPE_OBJECT:
		return json.Marshal(e.Object)
	default:
		return nil, fmt.Errorf("unable to parse type %s", e.Type)
	}
}

func (e Event) String() string {
	switch e.Type {
	case EVENT_TYPE_BOOLEAN:
		return fmt.Sprint(e.Boolean)
	case EVENT_TYPE_NUMBER:
		return fmt.Sprint(e.Number)
	case EVENT_TYPE_STRING:
		return fmt.Sprint(e.StringEvent)
	case EVENT_TYPE_OBJECT:
		return fmt.Sprint(e.Object)
	default:
		return ""
	}
}
