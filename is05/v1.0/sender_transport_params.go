package is05v1_0

import (
	"encoding/json"
	"fmt"
)

type SenderTransportParamType string

const (
	SENDER_TRANSPORT_PARAM_TYPE_RTP  SenderTransportParamType = "transport_rtp"
	SENDER_TRANSPORT_PARAM_TYPE_DASH SenderTransportParamType = "transport_dash"
)

type SenderTransportParam struct {
	Type      SenderTransportParamType
	ParamRTP  *SenderTransportParamRTP
	ParamDash *SenderTransportParamDash
}

func (p *SenderTransportParam) UnmarshalJSON(data []byte) error {
	// Unmarshall to generic interface
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	// Check if RTP param
	_, rtpEnabled_ok := dataTest["rtp_enabled"]
	if rtpEnabled_ok {
		var paramTemp SenderTransportParamRTP
		p.Type = SENDER_TRANSPORT_PARAM_TYPE_RTP
		err = json.Unmarshal(data, &paramTemp)
		if err != nil {
			return err
		}
		p.ParamRTP = &paramTemp
	}
	// Assume dash (because dash seems to be broken in the schemas of this version)
	var paramTemp SenderTransportParamDash
	err = json.Unmarshal(data, &paramTemp)
	if err != nil {
		return err
	}
	p.Type = SENDER_TRANSPORT_PARAM_TYPE_DASH
	p.ParamDash = &paramTemp

	return nil
}

func (p *SenderTransportParam) MarshalJSON() ([]byte, error) {
	switch p.Type {
	case SENDER_TRANSPORT_PARAM_TYPE_RTP:
		return json.Marshal((p.ParamRTP))
	case SENDER_TRANSPORT_PARAM_TYPE_DASH:
		return json.Marshal((p.ParamRTP))
	default:
		return nil, fmt.Errorf("unable to parse type %s", p.Type)
	}
}
