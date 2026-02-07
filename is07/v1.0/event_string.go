package is07v1_0

import "encoding/json"

// Describes an event with string payload
type EventString struct {
	EventCore
	EventType string             `json:"event_type"` // Event type
	Payload   EventStringPayload `json:"payload"`    // String payload
}

// String payload
type EventStringPayload struct {
	Value string `json:"value"` // String value
}

func (e *EventString) MarshalJSON() ([]byte, error) {
	type alias *EventString
	r := alias(e)
	r.MessageType = "state"
	return json.Marshal(r)
}
