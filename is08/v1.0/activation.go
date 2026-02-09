package is08v1_0

import "github.com/guregu/null/v6"

type ActivationMode string

const (
	ACTIVATION_MODE_IMMEDIATE           ActivationMode = "activate_immediate"
	ACTIVATION_MODE_SCHEDULED_ABSOLUTE  ActivationMode = "activate_scheduled_absolute"
	ACTIVATION_MODE_SCHEDULED_RELATEIVE ActivationMode = "activate_scheduled_relative"
)

// Parameters concerned with activation of the channel mapping
type Activation struct {
	Mode          ActivationMode `json:"mode"`           // Mode of activation: immediate (on message receipt), scheduled_absolute (when internal clock >= requested_time), or scheduled_relative (when internal clock >= time of message receipt + requested_time)
	RequestedTime null.String    `json:"requested_time"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating time (absolute or relative) for activation
}
