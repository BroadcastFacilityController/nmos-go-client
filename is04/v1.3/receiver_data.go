package is04v1_3

// Describes a data Receiver
type ReceiverData struct {
	ReceiverCore
	Format FormatURI        `json:"format"` // Type of Flow accepted by the Receiver as a URN
	Caps   ReceiverDataCaps `json:"caps"`   // Capabilities
}

// Capabilities
type ReceiverDataCaps struct {
	MediaTypes []string `json:"media_types,omitempty"` // Subclassification of the formats accepted using IANA assigned media types
	EventTypes []string `json:"event_types,omitempty"` // Subclassification of the event types accepted defined by the AMWA IS-07 specification
}
