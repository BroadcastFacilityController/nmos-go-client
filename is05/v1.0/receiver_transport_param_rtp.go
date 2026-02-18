package is05v1_0

import (
	"encoding/json"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

// Describes RTP Receiver transport parameters. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint. Receivers must support at least the `source_ip`, `interface_ip`, `rtp_enabled` and `destination_port` parameters, and must support the `multicast_ip` parameter if they are capable of multicast operation. Receivers supporting FEC and/or RTCP must support parameters prefixed with `fec` and `rtcp` respectively.
type ReceiverTransportParamRTP struct {
	SourceIP             null.String                           `json:"source_ip"`              // Source IP address of RTP packets in unicast mode, source filter for source specific multicast. A null value indicates that the receiver has not yet been configured, or in any-source multicast mode.
	MulticastIP          null.String                           `json:"multicast_ip"`           // IP multicast group address used in multicast operation only. Should be set to null during unicast operation. A null value indicates the parameter has not been configured, or the receiver is operating in unicast mode.
	InterfaceIP          string                                `json:"interface_ip"`           // IP address of the network interface the receiver should use. The receiver should provide an enum in the constraints endpoint, which should contain the available interface addresses. If set to auto in multicast mode the receiver should determine which interface to use for itself, for example by using the routing tables. The behaviour of auto is undefined in unicast mode, and controllers should supply a specific interface address.
	DestinationPort      common.AutoInt                        `json:"destination_port"`       // destination port for RTP packets (auto = 5004 by default).
	FECEnabled           bool                                  `json:"fec_enabled"`            // FEC on/off
	FECDestinationIP     string                                `json:"fec_destination_ip"`     // May be used if NAT is being used at the destination (auto = multicast_ip (multicast mode) or interface_ip (unicast mode) by default)
	FECMode              string                                `json:"fec_mode"`               // forward error correction mode to apply. (auto = highest available number of dimensions by default)
	FEC1DDestinationPort common.AutoInt                        `json:"fec1D_destination_port"` // destination port for RTP Column FEC packets (auto = RTP destination_port + 2 by default)
	FEC2DDestinationPort common.AutoInt                        `json:"fec2D_destination_port"` // destination port for RTP Row FEC packets (auto = RTP destination_port + 4 by default)
	RTCPDestinationIP    string                                `json:"rtcp_destination_ip"`    // Destination IP address of RTCP packets (auto = multicast_ip (multicast mode) or interface_ip (unicast mode) by default)
	RTCPEnabled          bool                                  `json:"rtcp_enabled"`           // RTCP on/off
	RTCPDestinationPort  common.AutoInt                        `json:"rtcp_destination_port"`  // destination port for RTCP packets (auto = RTP destination_port + 1 by default)
	RTPEnabled           bool                                  `json:"rtp_enabled"`            // RTP reception active/inactive
	IncludedKeys         map[ReceiverTransportParamRTPKey]bool `json:"-"`                      // JSON Keys which should be included in any json Marshalling. If empty, it will never be marshalled
}

const (
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_AUTO string = "auto"
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_1D   string = "1D"
	RECEIVER_TRANSPORT_PARAM_RTP_FEC_MODE_2D   string = "2D"
)

type ReceiverTransportParamRTPKey string

const (
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP              ReceiverTransportParamRTPKey = "source_ip"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_MULTICAST_IP           ReceiverTransportParamRTPKey = "multicast_ip"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_INTERFACE_IP           ReceiverTransportParamRTPKey = "interface_ip"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT       ReceiverTransportParamRTPKey = "destination_port"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED            ReceiverTransportParamRTPKey = "fec_enabled"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP     ReceiverTransportParamRTPKey = "fec_destination_ip"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE               ReceiverTransportParamRTPKey = "fec_mode"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT ReceiverTransportParamRTPKey = "fec1D_destination_port"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT ReceiverTransportParamRTPKey = "fec2D_destination_port"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED           ReceiverTransportParamRTPKey = "rtcp_enabled"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP    ReceiverTransportParamRTPKey = "rtcp_destination_ip"
	RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT  ReceiverTransportParamRTPKey = "rtcp_destination_port"
	RECIEVER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED            ReceiverTransportParamRTPKey = "rtp_enabled"
)

func (p *ReceiverTransportParamRTP) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias ReceiverTransportParamRTP
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if p.IncludedKeys == nil {
		p.IncludedKeys = make(map[ReceiverTransportParamRTPKey]bool)
	}
	for key := range dataTest {
		switch ReceiverTransportParamRTPKey(key) {
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP] = true
			p.SourceIP = dataTemp.SourceIP
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_MULTICAST_IP:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_MULTICAST_IP] = true
			p.MulticastIP = dataTemp.MulticastIP
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_INTERFACE_IP:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_INTERFACE_IP] = true
			p.InterfaceIP = dataTemp.InterfaceIP
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT] = true
			p.DestinationPort = dataTemp.DestinationPort
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED] = true
			p.FECEnabled = dataTemp.FECEnabled
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP] = true
			p.FECDestinationIP = dataTemp.FECDestinationIP
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE] = true
			p.FECMode = dataTemp.FECMode
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT] = true
			p.FEC1DDestinationPort = dataTemp.FEC1DDestinationPort
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT] = true
			p.FEC2DDestinationPort = dataTemp.FEC2DDestinationPort
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED] = true
			p.RTCPEnabled = dataTemp.RTCPEnabled
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP] = true
			p.RTCPDestinationIP = dataTemp.RTCPDestinationIP
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT:
			p.IncludedKeys[RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT] = true
			p.RTCPDestinationPort = dataTemp.RTCPDestinationPort
		case RECIEVER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED:
			p.IncludedKeys[RECIEVER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED] = true
			p.RTPEnabled = dataTemp.RTPEnabled
		}
	}
	return nil
}

func (p *ReceiverTransportParamRTP) MarshalJSON() ([]byte, error) {
	resultMap := make(map[string][]byte)

	for key, val := range p.IncludedKeys {
		if !val {
			// Ignore false keys
			continue
		}
		switch ReceiverTransportParamRTPKey(key) {
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP:
			valData, err := json.Marshal(&p.SourceIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_MULTICAST_IP:
			valData, err := json.Marshal(&p.MulticastIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_INTERFACE_IP:
			valData, err := json.Marshal(p.InterfaceIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT:
			valData, err := json.Marshal(&p.DestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED:
			valData, err := json.Marshal(p.FECEnabled)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP:
			valData, err := json.Marshal(p.FECDestinationIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE:
			valData, err := json.Marshal(p.FECMode)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT:
			valData, err := json.Marshal(&p.FEC1DDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT:
			valData, err := json.Marshal(&p.FEC2DDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED:
			valData, err := json.Marshal(p.RTCPEnabled)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP:
			valData, err := json.Marshal(p.RTCPDestinationIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECEIVER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT:
			valData, err := json.Marshal(&p.RTCPDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case RECIEVER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED:
			valData, err := json.Marshal(p.RTPEnabled)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		}

	}
	result := make([]byte, 0)
	for key, val := range resultMap {
		if len(result) == 0 {
			result = append(result, []byte{'{'}...)
		} else {
			result = append(result, []byte{','}...)
		}
		lineStr := make([]byte, 0)
		lineStr = append(lineStr, []byte{'"'}...)
		lineStr = append(lineStr, []byte(key)...)
		lineStr = append(lineStr, []byte{'"', ':'}...)
		lineStr = append(lineStr, val...)
		result = append(result, lineStr...)
	}
	result = append(result, '}')
	return result, nil
}
