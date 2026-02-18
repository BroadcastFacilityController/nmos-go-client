package is05v1_1

import (
	"github.com/guregu/null/v6"
)

type SenderResponse struct {
	ReceiverID      null.String            `json:"receiver_id"`      // ID of the target Receiver of this Sender. This will be null if the sender is operating in multicast mode, or has not been assigned a receiver in unicast mode, or is sending to a non-NMOS receiver in unicast mode.
	MasterEnable    bool                   `json:"master_enable"`    // Master on/off control for sender
	Activation      ActivationResponse     `json:"activation"`       // Parameters concerned with activation of the transport parameters
	TransportParams []SenderTransportParam `json:"transport_params"` // Transport-specific parameters. If this parameter is included in a client request it must include the same number of array elements (or 'legs') as specified in the constraints. If no changes are required to a specific leg it must be included as an empty object ({}).
}
