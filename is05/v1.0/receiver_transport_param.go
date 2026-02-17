package is05v1_0

import (
	"encoding/json"
	"fmt"
)

type ReceiverTransportParamType string

const (
	RECEIVER_TRANSPORT_PARAM_TYPE_RTP  ReceiverTransportParamType = "transport_rtp"
	RECEIVER_TRANSPORT_PARAM_TYPE_DASH ReceiverTransportParamType = "transport_dash"
)

type ReceiverTransportParam struct {
	Type      ReceiverTransportParamType
	ParamRTP  *ReceiverTransportParamRTP
	ParamDash *ReceiverTransportParamDash
}

func (p *ReceiverTransportParam) UnmarshalJSON(data []byte) error {
	// Unmarshall to generic interface
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	// Check if RTP param
	_, rtpEnabled_ok := dataTest["rtp_enabled"].(bool)
	if rtpEnabled_ok {
		var paramTemp ReceiverTransportParamRTP
		err = json.Unmarshal(data, &paramTemp)
		if err != nil {
			return err
		}
		p.ParamRTP = &paramTemp
		p.Type = RECEIVER_TRANSPORT_PARAM_TYPE_RTP
		return nil
	}
	// Assume dash (because dash seems to be broken in the schemas of this version)
	var paramTempDash ReceiverTransportParamDash
	err = json.Unmarshal(data, &paramTempDash)
	if err != nil {
		return err
	}
	p.ParamDash = &paramTempDash
	p.Type = RECEIVER_TRANSPORT_PARAM_TYPE_DASH

	return nil
}

func (p *ReceiverTransportParam) MarshalJSON() ([]byte, error) {
	switch p.Type {
	case RECEIVER_TRANSPORT_PARAM_TYPE_RTP:
		return json.Marshal(p.ParamRTP)
	case RECEIVER_TRANSPORT_PARAM_TYPE_DASH:
		return json.Marshal(p.ParamDash)
	default:
		return nil, fmt.Errorf("unable to marshal type %s", string(p.Type))
	}
}

func (p *ReceiverTransportParam) MarshalText() ([]byte, error) {
	switch p.Type {
	case RECEIVER_TRANSPORT_PARAM_TYPE_RTP:
		return fmt.Append(nil, p.ParamRTP), nil
	case RECEIVER_TRANSPORT_PARAM_TYPE_DASH:
		return fmt.Append(nil, p.ParamDash), nil
	default:
		return nil, fmt.Errorf("unable to marshal type %s", string(p.Type))
	}
}

func (p ReceiverTransportParam) String() string {
	switch p.Type {
	case RECEIVER_TRANSPORT_PARAM_TYPE_RTP:
		return fmt.Sprint(*p.ParamRTP)
	case RECEIVER_TRANSPORT_PARAM_TYPE_DASH:
		return fmt.Sprint(*p.ParamDash)
	default:
		return ""
	}
}
