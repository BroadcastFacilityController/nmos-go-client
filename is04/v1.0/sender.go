package is04v1_0

// Describes a sender
type Sender struct {
	ID           string              `json:"id"`             // Globally unique identifier for the Sender
	Version      string              `json:"version"`        // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label        string              `json:"label"`          // Freeform string label for the Sender
	Description  string              `json:"description"`    // Detailed description of the Sender
	FlowID       string              `json:"flow_id"`        // ID of the Flow currently passing via this Sender
	Transport    TransportURI        `json:"transport"`      // Transport type used by the Sender in URN format
	Tags         map[string][]string `json:"tags,omitempty"` // Key value set of freeform string tags to aid in filtering Senders. Values should be represented as an array of strings. Can be empty.
	DeviceID     string              `json:"device_id"`      // Device ID which this Sender forms part of
	ManifestHRef string              `json:"manifest_href"`  // HTTP URL to a file describing how to connect to the Sender (SDP for RTP)
}
