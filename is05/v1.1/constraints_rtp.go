package is05v1_1

import (
	"encoding/json"
	"maps"
)

type ConstraintTypeRTP string

const (
	CONSTRAINT_RTP_MULTICAST_IP            ConstraintTypeRTP = "multicast_ip"
	CONSTRAINT_RTP_DESTINATION_IP          ConstraintTypeRTP = "destination_ip"
	CONSTRAINT_RTP_DESTINATION_PORT        ConstraintTypeRTP = "destination_port"
	CONSTRAINT_RTP_SOURCE_IP               ConstraintTypeRTP = "source_ip"
	CONSTRAINT_RTP_INTERFACE_IP            ConstraintTypeRTP = "interface_ip"
	CONSTRAINT_RTP_SOURCE_PORT             ConstraintTypeRTP = "source_port"
	CONSTRAINT_RTP_FEC_ENABLED             ConstraintTypeRTP = "fec_enabled"
	CONSTRAINT_RTP_FEC_DESTINATION_IP      ConstraintTypeRTP = "fec_destination_ip"
	CONSTRAINT_RTP_FEC_MODE                ConstraintTypeRTP = "fec_mode"
	CONSTRAINT_RTP_FEC_TYPE                ConstraintTypeRTP = "fec_type"
	CONSTRAINT_RTP_FEC_BLOCK_WIDTH         ConstraintTypeRTP = "fec_block_width"
	CONSTRAINT_RTP_FEC_BLOCK_HEIGHT        ConstraintTypeRTP = "fec_block_height"
	CONSTRAINT_RTP_FEC_1D_DESTINATION_PORT ConstraintTypeRTP = "fec1D_destination_port"
	CONSTRAINT_RTP_FEC_2D_DESTINATION_PORT ConstraintTypeRTP = "fec2D_destination_port"
	CONSTRAINT_RTP_FEC_1D_SOURCE_PORT      ConstraintTypeRTP = "fec1D_source_port"
	CONSTRAINT_RTP_FEC_2D_SOURCE_PORT      ConstraintTypeRTP = "fec2D_source_port"
	CONSTRAINT_RTP_RTCP_ENABLED            ConstraintTypeRTP = "rtcp_enabled"
	CONSTRAINT_RTP_RTCP_DESTINATION_IP     ConstraintTypeRTP = "rtcp_destination_ip"
	CONSTRAINT_RTP_RTCP_DESTINATION_PORT   ConstraintTypeRTP = "rtcp_destination_port"
	CONSTRAINT_RTP_RTCP_SOURCE_PORT        ConstraintTypeRTP = "rtcp_source_port"
	CONSTRAINT_RTP_RTP_ENABLED             ConstraintTypeRTP = "rtp_enabled"
)

type ConstraintsRTP map[ConstraintTypeRTP]Constraint

func (c *ConstraintsRTP) UnmarshalJSON(data []byte) error {
	var dataTemp map[ConstraintTypeRTP]Constraint
	err := json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if *c == nil {
		*c = make(ConstraintsRTP)
	}
	maps.Copy((*c), dataTemp)
	return nil
}
