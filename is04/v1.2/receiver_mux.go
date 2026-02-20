package is04v1_2

// Describes a mux Receiver
type ReceiverMux struct {
	ReceiverCore
	Format FormatURI       `json:"format"` // Type of Flow accepted by the Receiver as a URN
	Caps   ReceiverMuxCaps `json:"caps"`   // Capabilities
}

// Capabilities
type ReceiverMuxCaps struct {
	MediaTypes []string `json:"media_types"` // Subclassification of the formats accepted using IANA assigned media types
}
