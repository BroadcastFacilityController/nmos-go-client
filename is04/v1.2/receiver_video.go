package is04v1_2

// Describes a video Receiver
type ReceiverVideo struct {
	ReceiverCore
	Format FormatURI         `json:"format"` // Type of Flow accepted by the Receiver as a URN
	Caps   ReceiverVideoCaps `json:"caps"`   // Capabilities
}

// Capabilities
type ReceiverVideoCaps struct {
	MediaTypes []string `json:"media_types"` // Subclassification of the formats accepted using IANA assigned media types
}
