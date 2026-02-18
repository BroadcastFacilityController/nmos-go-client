package is05v1_0

import (
	"encoding/json"
	"maps"
)

// The constraints for a single transport parameter
type Constraint struct {
	Max         int    `json:"maximum,omitempty"`     // The inclusive maximum value the parameter can be set to
	Min         int    `json:"minimum,omitempty"`     // The inclusive minimum value the parameter can be set to
	Enum        []any  `json:"enum,omitempty"`        // An array of allowed values
	Pattern     string `json:"pattern,omitempty"`     // A regex pattern that must be satisfied for this parameter
	Description string `json:"description,omitempty"` // A human readable string describing the constraint (optional)
}

type ConstraintType string

const (
	CONSTRAINT_MULTICAST_IP            ConstraintType = "multicast_ip"
	CONSTRAINT_DESTINATION_IP          ConstraintType = "destination_ip"
	CONSTRAINT_DESTINATION_PORT        ConstraintType = "destination_port"
	CONSTRAINT_SOURCE_IP               ConstraintType = "source_ip"
	CONSTRAINT_INTERFACE_IP            ConstraintType = "interface_ip"
	CONSTRAINT_SOURCE_PORT             ConstraintType = "source_port"
	CONSTRAINT_FEC_ENABLED             ConstraintType = "fec_enabled"
	CONSTRAINT_FEC_DESTINATION_IP      ConstraintType = "fec_destination_ip"
	CONSTRAINT_FEC_MODE                ConstraintType = "fec_mode"
	CONSTRAINT_FEC_BLOCK_WIDTH         ConstraintType = "fec_block_width"
	CONSTRAINT_FEC_BLOCK_HEIGHT        ConstraintType = "fec_block_height"
	CONSTRAINT_FEC_1D_DESTINATION_PORT ConstraintType = "fec1D_destination_port"
	CONSTRAINT_FEC_2D_DESTINATION_PORT ConstraintType = "fec2D_destination_port"
	CONSTRAINT_RTCP_ENABLED            ConstraintType = "rtcp_enabled"
	CONSTRAINT_RTCP_DESTINATION_IP     ConstraintType = "rtcp_destination_ip"
	CONSTRAINT_RTCP_DESTINATION_PORT   ConstraintType = "rtcp_destination_port"
	CONSTRAINT_RTCP_SOURCE_PORT        ConstraintType = "rtcp_source_port"
	CONSTRAINT_RTP_ENABLED             ConstraintType = "rtp_enabled"
)

type Constraints map[ConstraintType]Constraint

func (c *Constraints) UnmarshalJSON(data []byte) error {
	var dataTemp map[ConstraintType]Constraint
	err := json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if *c == nil {
		*c = make(Constraints)
	}
	maps.Copy((*c), dataTemp)
	return nil
}
