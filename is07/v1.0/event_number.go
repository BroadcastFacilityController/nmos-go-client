package is07v1_0

import "encoding/json"

// Describes an event with Number payload
type EventNumber struct {
	EventCore
	EventType string `json:"event_type"` // Event type
	Payload   Number `json:"payload"`    // Number payload
}

func (e *EventNumber) MarshalJSON() ([]byte, error) {
	type alias *EventNumber
	r := alias(e)
	r.MessageType = "state"
	return json.Marshal(r)
}
