package is05v1_0

import "github.com/guregu/null/v6"

// Describes a receiver
type Receiver struct {
	SenderID        null.String              `json:"sender_id"`     // ID of the Sender subscribed to by this Receiver. This will be null if the receiver has not been configured to receive anything, or if it is receiving from a non-NMOS sender.
	MasterEnable    bool                     `json:"master_enable"` // Master on/off control for receiver
	Activation      Activation               `json:"activation"`
	TransportFile   ReceiverTransportFile    `json:"transport_file"`   // Transport file parameters. 'data' and 'type' must both be strings or both be null
	TransportParams []ReceiverTransportParam `json:"transport_params"` // Transport-specific parameters. If this parameter is included in a client request it must include the same number of array elements (or 'legs') as specified in the constraints. If no changes are required to a specific leg it must be included as an empty object ({}).
}

// Transport file parameters. 'data' and 'type' must both be strings or both be null
type ReceiverTransportFile struct {
	Data null.String `json:"data"` // Content of the transport file
	Type null.String `json:"type"` // IANA assigned media type for file (e.g application/sdp)
}
