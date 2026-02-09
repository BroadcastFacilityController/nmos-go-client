package is08v1_0

import "github.com/guregu/null/v6"

type ActivationResponse struct {
	Mode           null.String `json:"mode"`            // Mode of activation: immediate (on message receipt), scheduled_absolute (when internal clock >= requested_time), scheduled_relative (when internal clock >= time of message receipt + requested_time), or null (no activation scheduled).
	RequestedTime  null.String `json:"requested_time"`  // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating time (absolute or relative) for activation requested. For an immediate activation this field will always be null.
	ActivationTime null.String `json:"activation_time"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating the absolute time the channel mapping will or did actually activate for scheduled activations, or the time activation occurred for immediate activations. For immediate activations this property will be the time the activation actually occurred in the response to the POST request.
}
