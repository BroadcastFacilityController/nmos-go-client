package is05v1_1

import (
	"encoding/json"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
)

// Describes RTP Sender transport parameters. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint. As a minimum all senders must support `source_ip`, `destination_ip`, `source_port`, `rtp_enabled` and `destination_port`. Senders supporting FEC and/or RTCP must support parameters prefixed with `fec` and `rtcp` respectively.
type SenderTransportParamRTP struct {
	SourceIP             string                              `json:"source_ip"`              // IP address from which RTP packets will be sent (IP address of interface bound to this output). The sender should provide an enum in the constraints endpoint, which should contain the available interface addresses. If the parameter is set to auto the sender should establish for itself which interface it should use, based on routing rules or its own internal configuration.
	DestinationIP        string                              `json:"destination_ip"`         // IP address to which RTP packets will be sent. If auto is set the sender should select a multicast address to send to itself. For example it may implement MADCAP (RFC 2730), ZMAAP, or be allocated address by some other system responsible for co-ordination multicast address use.
	SourcePort           common.AutoInt                      `json:"source_port"`            // source port for RTP packets (auto = 5004 by default)
	DestinationPort      common.AutoInt                      `json:"destination_port"`       // destination port for RTP packets (auto = 5004 by default)
	FECEnabled           bool                                `json:"fec_enabled"`            // FEC on/off
	FECDestinationIP     string                              `json:"fec_destination_ip"`     // May be used if NAT is being used at the destination (auto = destination_ip by default)
	FECType              SenderTransportParamRTPFECType      `json:"fec_type"`               // forward error correction mode to apply
	FECMode              SenderTransportParamRTPFECMode      `json:"fec_mode"`               // forward error correction mode to apply
	FECBlockWidth        int                                 `json:"fec_block_width"`        // width of block over which FEC is calculated in packets
	FECBlockHeight       int                                 `json:"fec_block_height"`       // height of block over which FEC is calculated in packets
	FEC1DDestinationPort common.AutoInt                      `json:"fec1D_destination_port"` // destination port for RTP Column FEC packets (auto = RTP destination_port + 2 by default)
	FEC2DDestinationPort common.AutoInt                      `json:"fec2D_destination_port"` // destination port for RTP Row FEC packets (auto = RTP destination_port + 4 by default)
	FEC1DSourcePort      common.AutoInt                      `json:"fec1D_source_port"`      // source port for RTP FEC packets (auto = RTP source_port + 2 by default)
	FEC2DSourcePort      common.AutoInt                      `json:"fec2D_source_port"`      // source port for RTP FEC packets (auto = RTP source_port + 4 by default)
	RTCPEnabled          bool                                `json:"rtcp_enabled"`           // rtcp on/off
	RTCPDestinationIP    string                              `json:"rtcp_destination_ip"`    // IP address to which RTCP packets will be sent (auto = same as RTP destination_ip by default)
	RTCPDestinationPort  common.AutoInt                      `json:"rtcp_destination_port"`  // destination port for RTCP packets (auto = RTP destination_port + 1 by default)
	RTCPSourcePort       common.AutoInt                      `json:"rtcp_source_port"`       // source port for RTCP packets (auto = RTP source_port + 1 by default)
	RTPEnabled           bool                                `json:"rtp_enabled"`            // RTP transmission active/inactive
	IncludedKeys         map[SenderTransportParamRTPKey]bool `json:"-"`                      // Keys to be included in JSON
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

type SenderTransportParamRTPKey string

const (
	SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP              SenderTransportParamRTPKey = "source_ip"
	SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_IP         SenderTransportParamRTPKey = "destination_ip"
	SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_PORT            SenderTransportParamRTPKey = "source_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT       SenderTransportParamRTPKey = "destination_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED            SenderTransportParamRTPKey = "fec_enabled"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP     SenderTransportParamRTPKey = "fec_destination_ip"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_TYPE               SenderTransportParamRTPKey = "fec_type"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE               SenderTransportParamRTPKey = "fec_mode"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_WIDTH        SenderTransportParamRTPKey = "fec_block_width"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_HEIGHT       SenderTransportParamRTPKey = "fec_block_height"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT SenderTransportParamRTPKey = "fec1D_destination_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT SenderTransportParamRTPKey = "fec2D_destination_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_SOURCE_PORT      SenderTransportParamRTPKey = "fec1D_source_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_SOURCE_PORT      SenderTransportParamRTPKey = "fec2D_source_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED           SenderTransportParamRTPKey = "rtcp_enabled"
	SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP    SenderTransportParamRTPKey = "rtcp_destination_ip"
	SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT  SenderTransportParamRTPKey = "rtcp_destination_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_SOURCE_PORT       SenderTransportParamRTPKey = "rtcp_source_port"
	SENDER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED            SenderTransportParamRTPKey = "rtp_enabled"
)

func (p *SenderTransportParamRTP) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias SenderTransportParamRTP
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if p.IncludedKeys == nil {
		p.IncludedKeys = make(map[SenderTransportParamRTPKey]bool)
	}
	for key := range dataTest {
		switch SenderTransportParamRTPKey(key) {
		case SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP] = true
			p.SourceIP = dataTemp.SourceIP
		case SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_IP:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_IP] = true
			p.DestinationIP = dataTemp.DestinationIP
		case SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_PORT] = true
			p.SourcePort = dataTemp.SourcePort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT] = true
			p.DestinationPort = dataTemp.DestinationPort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED] = true
			p.FECEnabled = dataTemp.FECEnabled
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP] = true
			p.FECDestinationIP = dataTemp.FECDestinationIP
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_TYPE:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_TYPE] = true
			p.FECType = dataTemp.FECType
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE] = true
			p.FECMode = dataTemp.FECMode
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_WIDTH:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_WIDTH] = true
			p.FECBlockWidth = dataTemp.FECBlockWidth
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_HEIGHT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_HEIGHT] = true
			p.FECBlockHeight = dataTemp.FECBlockHeight
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT] = true
			p.FEC1DDestinationPort = dataTemp.FEC1DDestinationPort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT] = true
			p.FEC2DDestinationPort = dataTemp.FEC2DDestinationPort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_SOURCE_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_SOURCE_PORT] = true
			p.FEC1DSourcePort = dataTemp.FEC1DSourcePort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_SOURCE_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_SOURCE_PORT] = true
			p.FEC2DSourcePort = dataTemp.FEC2DSourcePort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED] = true
			p.RTCPEnabled = dataTemp.RTCPEnabled
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP] = true
			p.RTCPDestinationIP = dataTemp.RTCPDestinationIP
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT] = true
			p.RTCPDestinationPort = dataTemp.RTCPDestinationPort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_SOURCE_PORT:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_SOURCE_PORT] = true
			p.RTCPSourcePort = dataTemp.RTCPSourcePort
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED:
			p.IncludedKeys[SENDER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED] = true
			p.RTPEnabled = dataTemp.RTPEnabled
		}
	}
	return nil
}

func (p *SenderTransportParamRTP) MarshalJSON() ([]byte, error) {
	resultMap := make(map[string][]byte)

	for key, val := range p.IncludedKeys {
		if !val {
			// Ignore false keys
			continue
		}
		switch SenderTransportParamRTPKey(key) {
		case SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_IP:
			valData, err := json.Marshal(p.SourceIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_IP:
			valData, err := json.Marshal(p.DestinationIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_SOURCE_PORT:
			valData, err := json.Marshal(&p.SourcePort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_DESTINATION_PORT:
			valData, err := json.Marshal(&p.DestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_ENABLED:
			valData, err := json.Marshal(p.FECEnabled)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_DESTINATION_IP:
			valData, err := json.Marshal(p.FECDestinationIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_TYPE:
			valData, err := json.Marshal(p.FECType)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_MODE:
			valData, err := json.Marshal(p.FECMode)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_WIDTH:
			valData, err := json.Marshal(p.FECBlockWidth)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC_BLOCK_HEIGHT:
			valData, err := json.Marshal(p.FECBlockHeight)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_DESTINATION_PORT:
			valData, err := json.Marshal(&p.FEC1DDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_DESTINATION_PORT:
			valData, err := json.Marshal(&p.FEC2DDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC1D_SOURCE_PORT:
			valData, err := json.Marshal(&p.FEC1DSourcePort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_FEC2D_SOURCE_PORT:
			valData, err := json.Marshal(&p.FEC2DSourcePort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_ENABLED:
			valData, err := json.Marshal(p.RTCPEnabled)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_IP:
			valData, err := json.Marshal(p.RTCPDestinationIP)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_DESTINATION_PORT:
			valData, err := json.Marshal(&p.RTCPDestinationPort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTCP_SOURCE_PORT:
			valData, err := json.Marshal(&p.RTCPSourcePort)
			if err != nil {
				return nil, err
			}
			resultMap[string(key)] = valData
		case SENDER_TRANSPORT_PARAM_RTP_KEY_RTP_ENABLED:
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
