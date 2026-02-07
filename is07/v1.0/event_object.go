package is07v1_0

import "encoding/json"

// Describes an event with object payload
type EventObject struct {
	EventCore
	EventType string `json:"event_type"` // Event type
	Payload   any    `json:"payload"`    // Object payload
}

func (e *EventObject) MarshalJSON() ([]byte, error) {
	type alias *EventObject
	r := alias(e)
	r.MessageType = "state"
	return json.Marshal(r)
}
