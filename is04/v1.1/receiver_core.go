package is04v1_1

import (
	"github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

// Describes a receiver
type ReceiverCore struct {
	ResourceCore
	DeviceID     string               `json:"device_id"`    // Device ID which this Receiver forms part of. This attribute is used to ensure referential integrity by registry implementations.
	Transport    TransportURI         `json:"transport"`    // Transport type accepted by the Receiver in URN format
	Subscription ReceiverSubscription `json:"subscription"` // Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
}

// Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
type ReceiverSubscription struct {
	SenderID null.String `json:"sender_id"` // UUID of the Sender that this Receiver is currently subscribed to
	Active   zero.Bool   `json:"active"`    // Not in spec, but usually implemented. True when sender_id is not null
}
