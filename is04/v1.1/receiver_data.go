package is04v1_1

// Describes a data Receiver
type ReceiverData struct {
	ReceiverCore
	Format FormatURI        `json:"format"` // Type of Flow accepted by the Receiver as a URN
	Caps   ReceiverDataCaps `json:"caps"`   // Capabilities
}

// Capabilities
type ReceiverDataCaps struct {
	MediaTypes []IANAMediaType `json:"media_types"` // Subclassification of the formats accepted using IANA assigned media types
}
