package is04v1_1

// Describes a sender
type Sender struct {
	ResourceCore
	FlowID       string       `json:"flow_id"`       // ID of the Flow currently passing via this Sender
	Transport    TransportURI `json:"transport"`     // Transport type used by the Sender in URN format
	DeviceID     string       `json:"device_id"`     // Device ID which this Sender forms part of
	ManifestHRef string       `json:"manifest_href"` // HTTP URL to a file describing how to connect to the Sender (SDP for RTP)
}
