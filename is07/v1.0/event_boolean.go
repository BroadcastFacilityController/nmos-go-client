package is07v1_0

import "encoding/json"

type EventBoolean struct {
	EventCore
	EventType string              `json:"event_type"` // Event type
	Payload   EventBooleanPayload `json:"payload"`    // Boolean payload
}

type EventBooleanPayload struct {
	Value bool `json:"value"` // Boolean value
}

func (e *EventBoolean) MarshalJSON() ([]byte, error) {
	type alias *EventBoolean
	r := alias(e)
	r.MessageType = "state"
	r.EventType = "boolean"
	return json.Marshal(r)
}
