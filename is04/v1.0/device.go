package is04v1_0

// Describes a Device
type Device struct {
	ID        string   `json:"id"`        // Globally unique identifier for the Device
	Version   string   `json:"version"`   // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label     string   `json:"label"`     // Freeform string label for the Device
	Type      string   `json:"type"`      // Device type URN
	NodeID    string   `json:"node_id"`   // Globally unique identifier for the Node which initially created the Device
	Senders   []string `json:"senders"`   // UUIDs of Senders attached to the Device
	Receivers []string `json:"receivers"` // UUIDs of Receivers attached to the Device
}
