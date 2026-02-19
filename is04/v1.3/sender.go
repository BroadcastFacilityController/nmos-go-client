package is04v1_3

import "github.com/guregu/null/v6"

// Describes a sender
type Sender struct {
	ResourceCore
	Caps              any                `json:"caps"`               // Capabilities of this sender
	FlowID            null.String        `json:"flow_id"`            // ID of the Flow currently passing via this Sender. Set to null when a Flow is not currently internally routed to the Sender.
	Transport         TransportURI       `json:"transport"`          // Transport type used by the Sender in URN format
	DeviceID          string             `json:"device_id"`          // Device ID which this Sender forms part of. This attribute is used to ensure referential integrity by registry implementations.
	ManifestHRef      null.String        `json:"manifest_href"`      // HTTP(S) accessible URL to a file describing how to connect to the Sender. Set to null when the transport type used by the Sender does not require a transport file.
	InterfaceBindings []string           `json:"interface_bindings"` // Binding of Sender egress ports to interfaces on the parent Node.
	Subscription      SenderSubscription `json:"subscription"`       // Object indicating how this Sender is currently configured to send data.
}

type SenderSubscription struct {
	ReceiverID null.String `json:"receiver_id"` // UUID of the Receiver to which this Sender is currently configured to send data. Only set if it is active, uses a unicast push-based transport and is sending to an NMOS Receiver; otherwise null.
	Active     bool        `json:"active"`      // Sender is enabled and configured to send data
}
