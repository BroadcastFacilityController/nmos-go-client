package is05v1_1

import (
	"encoding/json"
	"fmt"
)

type SenderTransportParamType string

const (
	SENDER_TRANSPORT_PARAM_TYPE_RTP       SenderTransportParamType = "transport_rtp"
	SENDER_TRANSPORT_PARAM_TYPE_WEBSOCKET SenderTransportParamType = "transport_websocket"
	SENDER_TRANSPORT_PARAM_TYPE_MQTT      SenderTransportParamType = "transport_mqtt"
	SENDER_TRANSPORT_PARAM_TYPE_DASH      SenderTransportParamType = "transport_dash"
)

type SenderTransportParam struct {
	Type           SenderTransportParamType
	ParamRTP       *SenderTransportParamRTP
	ParamWebsocket *SenderTransportParamWebsocket
	ParamMQTT      *SenderTransportParamMQTT
	ParamDash      *SenderTransportParamDash
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
		var param SenderTransportParamRTP
		err = json.Unmarshal(data, &param)
		if err != nil {
			return err
		}
		p.ParamRTP = &param
		p.Type = SENDER_TRANSPORT_PARAM_TYPE_RTP
		return nil
	}
	// Check if Websocket param
	_, connectionURI_ok := dataTest["connection_uri"]
	if connectionURI_ok {
		var param SenderTransportParamWebsocket
		err = json.Unmarshal(data, &param)
		if err != nil {
			return err
		}
		p.ParamWebsocket = &param
		p.Type = SENDER_TRANSPORT_PARAM_TYPE_WEBSOCKET
		return nil
	}
	// Check if MQTT param
	_, brokerTopic_ok := dataTest["broker_topic"]
	if brokerTopic_ok {
		var param SenderTransportParamMQTT
		err = json.Unmarshal(data, &param)
		if err != nil {
			return err
		}
		p.ParamMQTT = &param
		p.Type = SENDER_TRANSPORT_PARAM_TYPE_MQTT
		return nil
	}
	// Assume dash (because dash seems to be broken in the schemas of this version)
	err = json.Unmarshal(data, p.ParamDash)
	if err != nil {
		return err
	}
	p.Type = SENDER_TRANSPORT_PARAM_TYPE_DASH

	return nil
}

func (p *SenderTransportParam) MarshalJSON() ([]byte, error) {
	switch p.Type {
	case SENDER_TRANSPORT_PARAM_TYPE_RTP:
		return json.Marshal(p.ParamRTP)
	case SENDER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return json.Marshal(p.ParamWebsocket)
	case SENDER_TRANSPORT_PARAM_TYPE_MQTT:
		return json.Marshal(p.ParamMQTT)
	case SENDER_TRANSPORT_PARAM_TYPE_DASH:
		return json.Marshal(p.ParamDash)
	default:
		return nil, fmt.Errorf("unable to marshal type %s", string(p.Type))
	}
}

func (p *SenderTransportParam) MarshalText() ([]byte, error) {
	switch p.Type {
	case SENDER_TRANSPORT_PARAM_TYPE_RTP:
		return fmt.Append(nil, p.ParamRTP), nil
	case SENDER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return fmt.Append(nil, p.ParamWebsocket), nil
	case SENDER_TRANSPORT_PARAM_TYPE_MQTT:
		return fmt.Append(nil, p.ParamMQTT), nil
	case SENDER_TRANSPORT_PARAM_TYPE_DASH:
		return fmt.Append(nil, p.ParamDash), nil
	default:
		return nil, fmt.Errorf("unable to marshal type %s", string(p.Type))
	}
}

func (p SenderTransportParam) String() string {
	switch p.Type {
	case SENDER_TRANSPORT_PARAM_TYPE_RTP:
		return fmt.Sprint(*p.ParamRTP)
	case SENDER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return fmt.Sprint(*p.ParamWebsocket)
	case SENDER_TRANSPORT_PARAM_TYPE_MQTT:
		return fmt.Sprint(*p.ParamMQTT)
	case SENDER_TRANSPORT_PARAM_TYPE_DASH:
		return fmt.Sprint(*p.ParamDash)
	default:
		return ""
	}
}
