package is07v1_0

// Message that gets sent out when the device start a shutdown or reboot process
type MessageShutdownReboot struct {
	MessageType string                        `json:"message_type"` // A fixed string showing this is a reboot or a shutdown message
	Identity    MessageShutdownRebootIdentity `json:"identity"`     // Object describing event identity
	Timing      MessageShutdownRebootTiming   `json:"timing"`       // Object describing event timing
}

// Object describing event identity
type MessageShutdownRebootIdentity struct {
	SourceID string `json:"source_id"`         // ID of the source which identifies the emitter of the event
	FlowID   string `json:"flow_id,omitempty"` // ID of the flow carrying the message
}

// Object describing event timing
type MessageShutdownRebootTiming struct {
	CreationTimestamp string `json:"creation_timestamp"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating when the message was created
}
