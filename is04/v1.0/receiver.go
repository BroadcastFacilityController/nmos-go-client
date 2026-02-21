package is04v1_0

import (
	"github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

// Describes a receiver
type Receiver struct {
	ID           string               `json:"id"`           // Globally unique identifier for the Receiver
	Version      string               `json:"version"`      // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label        string               `json:"label"`        // Freeform string label for the Receiver
	Description  string               `json:"description"`  // Detailed description of the Receiver
	Format       FormatURI            `json:"format"`       // Type of Flow accepted by the Receiver as a URN
	Caps         any                  `json:"caps"`         // Capabilities (not yet defined)
	Tags         map[string][]string  `json:"tags"`         // Key value set of freeform string tags to aid in filtering sources. Values should be represented as an array of strings. Can be empty.
	DeviceID     string               `json:"device_id"`    // Device ID which this Receiver forms part of
	Transport    TransportURI         `json:"transport"`    // Transport type accepted by the Receiver in URN format
	Subscription ReceiverSubscription `json:"subscription"` // Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
}

// Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
type ReceiverSubscription struct {
	SenderID null.String `json:"sender_id"`       // UUID of the Sender that this Receiver is currently subscribed to
	Active   zero.Bool   `json:"active,omitzero"` // Not in spec, but usually implemented. True when sender_id is not null
}
