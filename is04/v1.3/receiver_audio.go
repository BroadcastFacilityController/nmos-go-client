package is04v1_3

// Describes an audio Receiver
type ReceiverAudio struct {
	ReceiverCore
	Format FormatURI         `json:"format"` // Type of Flow accepted by the Receiver as a URN
	Caps   ReceiverAudioCaps `json:"caps"`   // Capabilities
}

// Capabilities
type ReceiverAudioCaps struct {
	MediaTypes []string `json:"media_types"` // Subclassifications of the format accepted, using IANA assigned media types where available, or other values defined in the NMOS Parameter Registers
}
