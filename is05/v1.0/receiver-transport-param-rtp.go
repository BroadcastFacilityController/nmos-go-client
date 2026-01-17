package is05v1_0

import (
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

// Describes RTP Receiver transport parameters. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint. Receivers must support at least the `source_ip`, `interface_ip`, `rtp_enabled` and `destination_port` parameters, and must support the `multicast_ip` parameter if they are capable of multicast operation. Receivers supporting FEC and/or RTCP must support parameters prefixed with `fec` and `rtcp` respectively.
type ReceiverTransportParamRTP struct {
	SourceIP             null.String    `json:"source_ip"`              // Source IP address of RTP packets in unicast mode, source filter for source specific multicast. A null value indicates that the receiver has not yet been configured, or in any-source multicast mode.
	MulticastIP          null.String    `json:"multicast_ip"`           // IP multicast group address used in multicast operation only. Should be set to null during unicast operation. A null value indicates the parameter has not been configured, or the receiver is operating in unicast mode.
	InterfaceIP          string         `json:"interface_ip"`           // IP address of the network interface the receiver should use. The receiver should provide an enum in the constraints endpoint, which should contain the available interface addresses. If set to auto in multicast mode the receiver should determine which interface to use for itself, for example by using the routing tables. The behaviour of auto is undefined in unicast mode, and controllers should supply a specific interface address.
	DestinationPort      common.AutoInt `json:"destination_port"`       // destination port for RTP packets (auto = 5004 by default).
	FECEnabled           bool           `json:"fec_enabled"`            // FEC on/off
	FECDestinationIP     string         `json:"fec_destination_ip"`     // May be used if NAT is being used at the destination (auto = multicast_ip (multicast mode) or interface_ip (unicast mode) by default)
	FECMode              string         `json:"fec_mode"`               // forward error correction mode to apply. (auto = highest available number of dimensions by default)
	FEC1DDestinationPort common.AutoInt `json:"fec1D_destination_port"` // destination port for RTP Column FEC packets (auto = RTP destination_port + 2 by default)
	FEC2DDestinationPort common.AutoInt `json:"fec2D_destination_port"` // destination port for RTP Row FEC packets (auto = RTP destination_port + 4 by default)
	RTCPDestinationIP    string         `json:"rtcp_destination_ip"`    // Destination IP address of RTCP packets (auto = multicast_ip (multicast mode) or interface_ip (unicast mode) by default)
	RTCPEnabled          bool           `json:"rtcp_enabled"`           // RTCP on/off
	RTCPDestinationPort  common.AutoInt `json:"rtcp_destination_port"`  // destination port for RTCP packets (auto = RTP destination_port + 1 by default)
	RTPEnabled           bool           `json:"rtp_enabled"`            // RTP reception active/inactive
}

const (
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_AUTO string = "auto"
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_1D   string = "1D"
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_2D   string = "2D"
)
