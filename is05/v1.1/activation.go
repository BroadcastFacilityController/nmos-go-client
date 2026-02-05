package is05v1_1

import "github.com/guregu/null/v6"

const (
	ACTIVATION_MODE_IMMEDIATE          string = "activate_immediate"
	ACTIVATION_MODE_SCHEDULED_ABSOLUTE string = "activate_scheduled_absolute"
	ACTIVATION_MODE_SCHEDULED_RELATIVE string = "activate_scheduled_relative"
)

// Parameters concerned with activation of the transport parameters
type Activation struct {
	Mode          null.String `json:"mode"`           // Mode of activation: immediate (on message receipt), scheduled_absolute (when internal clock >= requested_time), scheduled_relative (when internal clock >= time of message receipt + requested_time), or null (no activation scheduled)
	RequestedTime null.String `json:"requested_time"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating time (absolute or relative) for activation. Should be null or not present if 'mode' is null.
}
