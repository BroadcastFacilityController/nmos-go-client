package is05v1_1

import "github.com/BroadcastFacilityController/nmos-go-client/common"

// Describes RTP Sender transport parameters. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint. As a minimum all senders must support `source_ip`, `destination_ip`, `source_port`, `rtp_enabled` and `destination_port`. Senders supporting FEC and/or RTCP must support parameters prefixed with `fec` and `rtcp` respectively.
type SenderTransportParamRTP struct {
	SourceIP             string                         `json:"source_ip"`              // IP address from which RTP packets will be sent (IP address of interface bound to this output). The sender should provide an enum in the constraints endpoint, which should contain the available interface addresses. If the parameter is set to auto the sender should establish for itself which interface it should use, based on routing rules or its own internal configuration.
	DestinationIP        string                         `json:"destination_ip"`         // IP address to which RTP packets will be sent. If auto is set the sender should select a multicast address to send to itself. For example it may implement MADCAP (RFC 2730), ZMAAP, or be allocated address by some other system responsible for co-ordination multicast address use.
	SourcePort           common.AutoInt                 `json:"source_port"`            // source port for RTP packets (auto = 5004 by default)
	DestinationPort      common.AutoInt                 `json:"destination_port"`       // destination port for RTP packets (auto = 5004 by default)
	FECEnabled           bool                           `json:"fec_enabled"`            // FEC on/off
	FECDestinationIP     string                         `json:"fec_destination_ip"`     // May be used if NAT is being used at the destination (auto = destination_ip by default)
	FECType              SenderTransportParamRTPFECType `json:"fec_type"`               // forward error correction mode to apply
	FECMode              SenderTransportParamRTPFECMode `json:"fec_mode"`               // forward error correction mode to apply
	FECBlockWidth        int                            `json:"fec_block_width"`        // width of block over which FEC is calculated in packets
	FECBlockHeight       int                            `json:"fec_block_height"`       // height of block over which FEC is calculated in packets
	FEC1DDestinationPort common.AutoInt                 `json:"fec1D_destination_port"` // destination port for RTP Column FEC packets (auto = RTP destination_port + 2 by default)
	FEC2DDestinationPort common.AutoInt                 `json:"fec2D_destination_port"` // destination port for RTP Row FEC packets (auto = RTP destination_port + 4 by default)
	FEC1DSourcePort      common.AutoInt                 `json:"fec1D_source_port"`      // source port for RTP FEC packets (auto = RTP source_port + 2 by default)
	FEC2DSourcePort      common.AutoInt                 `json:"fec2D_source_port"`      // source port for RTP FEC packets (auto = RTP source_port + 4 by default)
	RTCPEnabled          bool                           `json:"rtcp_enabled"`           // rtcp on/off
	RTCPDestinationIP    string                         `json:"rtcp_destination_ip"`    // IP address to which RTCP packets will be sent (auto = same as RTP destination_ip by default)
	RTCPDestinationPort  common.AutoInt                 `json:"rtcp_destination_port"`  // destination port for RTCP packets (auto = RTP destination_port + 1 by default)
	RTCPSourcePort       common.AutoInt                 `json:"rtcp_source_port"`       // source port for RTCP packets (auto = RTP source_port + 1 by default)
	RTPEnabled           bool                           `json:"rtp_enabled"`            // RTP transmission active/inactive
}

type SenderTransportParamRTPFECType string

const (
	SENDER_TRANSPORT_PARAM_RTP_FEC_TYPE_XOR SenderTransportParamRTPFECType = "XOR"
	SENDER_TRANSPORT_PARAM_RTP_FEC_TYPE_RS  SenderTransportParamRTPFECType = "Reed-Solomon"
)

type SenderTransportParamRTPFECMode string

const (
	SENDER_TRANSPORT_PARAM_RTP_FEC_MODE_1D SenderTransportParamRTPFECMode = "1D"
	SENDER_TRANSPORT_PARAM_RTP_FEC_MODE_2D SenderTransportParamRTPFECMode = "2D"
)
