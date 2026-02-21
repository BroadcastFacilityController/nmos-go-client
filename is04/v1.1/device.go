package is04v1_1

// Describes a Device
type Device struct {
	ResourceCore
	Type      DeviceTypeURI `json:"type"`      // Device type URN
	NodeID    string        `json:"node_id"`   // Globally unique identifier for the Node which initially created the Device
	Senders   []string      `json:"senders"`   // UUIDs of Senders attached to the Device
	Receivers []string      `json:"receivers"` // UUIDs of Receivers attached to the Device
}

type DeviceTypeURI URI

const (
	DEVICE_GENERIC  DeviceTypeURI = "urn:x-nmos:device:generic"
	DEVICE_PIPELINE DeviceTypeURI = "urn:x-nmos:device:pipeline"
)
