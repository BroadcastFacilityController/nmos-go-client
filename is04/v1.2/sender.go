package is04v1_2

import "github.com/guregu/null/v6"

// Describes a sender
type Sender struct {
	ResourceCore
	Caps              any                `json:"caps"`               // Capabilities of this sender
	FlowID            string             `json:"flow_id"`            // ID of the Flow currently passing via this Sender
	Transport         TransportURI       `json:"transport"`          // Transport type used by the Sender in URN format
	DeviceID          string             `json:"device_id"`          // Device ID which this Sender forms part of. This attribute is used to ensure referential integrity by registry implementations.
	ManifestHRef      string             `json:"manifest_href"`      // HTTP URL to a file describing how to connect to the Sender (SDP for RTP). The Sender's 'version' attribute should be updated if the contents of this file are modified. This URL may return an HTTP 404 where the 'active' parameter in the 'subscription' object is present and set to false (v1.2+ only).
	InterfaceBindings []string           `json:"interface_bindings"` // Binding of Sender egress ports to interfaces on the parent Node. Should contain a single network interface unless a redundancy mechanism such as ST.2022-7 is in use, in which case each 'leg' should have its matching interface listed. Where the redundancy mechanism sends more than one copy of the stream via the same interface, that interface should be listed a corresponding number of times.
	Subscription      SenderSubscription `json:"subscription"`       // Object containing the 'receiver_id' currently subscribed to (unicast only). Receiver_id should be null on initialisation, or when connected to a non-NMOS unicast Receiver.
}

type SenderSubscription struct {
	ReceiverID null.String `json:"receiver_id"` // UUID of the Receiver that this Sender is currently subscribed to
	Active     bool        `json:"active"`      // Sender is enabled and configured to stream data to a single Receiver (unicast), or to the network via multicast or a pull-based mechanism
}
