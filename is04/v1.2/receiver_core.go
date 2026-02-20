package is04v1_2

import "github.com/guregu/null/v6"

// Describes a receiver
type ReceiverCore struct {
	ResourceCore
	DeviceID          string               `json:"device_id"`          // Device ID which this Receiver forms part of. This attribute is used to ensure referential integrity by registry implementations.
	Transport         TransportURI         `json:"transport"`          // Transport type accepted by the Receiver in URN format
	InterfaceBindings []string             `json:"interface_bindings"` // Binding of Receiver ingress ports to interfaces on the parent Node. Should contain a single network interface unless a redundancy mechanism such as ST.2022-7 is in use, in which case each 'leg' should have its matching interface listed. Where the redundancy mechanism receives more than one copy of the stream via the same interface, that interface should be listed a corresponding number of times.
	Subscription      ReceiverSubscription `json:"subscription"`       // Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation, or when connected to a non-NMOS Sender.
}

// Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
type ReceiverSubscription struct {
	SenderID null.String `json:"sender_id"` // UUID of the Sender that this Receiver is currently subscribed to
	Active   bool        `json:"active"`    // Receiver is enabled and configured with a Sender's connection parameters
}
