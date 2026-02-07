package is07v1_0

// Message that gets sent out when the source state changes
type EventCore struct {
	MessageType string            `json:"message_type"` // A fixed string showing this is a state message
	Identity    EventCoreIdentity `json:"identity"`     // Object describing event identity
	Timing      EventCoreTiming   `json:"timing"`       // Object describing event timing
}

// Object describing event identity
type EventCoreIdentity struct {
	SourceID string `json:"source_id"`         // ID of the source that holds the state that has changed
	FlowID   string `json:"flow_id,omitempty"` // ID of the flow carrying the message
}

// Object describing event timing
type EventCoreTiming struct {
	CreationTimestamp string `json:"creation_timestamp"`         // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating when the message was created
	OriginTimestamp   string `json:"origin_timestamp,omitempty"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating when the message that triggered this message (if any) was created
	ActionTimestamp   string `json:"action_timestamp,omitempty"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating when the state change described in the event took or will take place
}
