package is07v1_0

import "encoding/json"

// Message that gets sent out as the response to a health (heartbeat) command
type MessageHealth struct {
	MessageType string              `json:"message_type"` // A fixed string showing this is a health message
	Timing      MessageHealthTiming `json:"timing"`       // Object describing event timing
}

// Object describing event timing
type MessageHealthTiming struct {
	CreationTimestamp string `json:"creation_timestamp"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating when the message was created
	OriginTimestamp   string `json:"origin_timestamp"`   // String formatted TAI timestamp (<seconds>:<nanoseconds>) copied from the incoming health command
}

func (m *MessageHealth) MarshalJSON() ([]byte, error) {
	type alias *MessageHealth
	r := alias(m)
	r.MessageType = "health"
	return json.Marshal(r)
}
