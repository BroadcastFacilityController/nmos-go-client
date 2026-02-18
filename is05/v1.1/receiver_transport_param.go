package is05v1_1

import (
	"encoding/json"
	"fmt"
)

type ReceiverTransportParamType string

const (
	RECEIVER_TRANSPORT_PARAM_TYPE_RTP       ReceiverTransportParamType = "transport_rtp"
	RECEIVER_TRANSPORT_PARAM_TYPE_DASH      ReceiverTransportParamType = "transport_dash"
	RECEIVER_TRANSPORT_PARAM_TYPE_WEBSOCKET ReceiverTransportParamType = "transport_websocket"
	RECEIVER_TRANSPORT_PARAM_TYPE_MQTT      ReceiverTransportParamType = "transport_mqtt"
)

type ReceiverTransportParam struct {
	Type           ReceiverTransportParamType
	ParamRTP       *ReceiverTransportParamRTP
	ParamDash      *ReceiverTransportParamDash
	ParamWebsocket *ReceiverTransportParamWebsocket
	ParamMQTT      *ReceiverTransportParamMQTT
}

func (p *ReceiverTransportParam) UnmarshalJSON(data []byte) error {
	// Unmarshall to generic interface
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	// Check if RTP param
	_, rtpEnabled_ok := dataTest["rtp_enabled"]
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
	// Check if Websocket param
	_, connectionURI_ok := dataTest["connection_uri"]
	if connectionURI_ok {
		var paramTemp ReceiverTransportParamWebsocket
		err = json.Unmarshal(data, &paramTemp)
		if err != nil {
			return err
		}
		p.ParamWebsocket = &paramTemp
		p.Type = RECEIVER_TRANSPORT_PARAM_TYPE_WEBSOCKET
		return nil
	}
	// Check if MQTT param
	_, brokerTopic_ok := dataTest["broker_topic"]
	if brokerTopic_ok {
		var paramTemp ReceiverTransportParamMQTT
		err = json.Unmarshal(data, &paramTemp)
		if err != nil {
			return err
		}
		p.ParamMQTT = &paramTemp
		p.Type = RECEIVER_TRANSPORT_PARAM_TYPE_MQTT
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
	case RECEIVER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return json.Marshal(p.ParamWebsocket)
	case RECEIVER_TRANSPORT_PARAM_TYPE_MQTT:
		return json.Marshal(p.ParamMQTT)
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
	case RECEIVER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return fmt.Append(nil, p.ParamWebsocket), nil
	case RECEIVER_TRANSPORT_PARAM_TYPE_MQTT:
		return fmt.Append(nil, p.ParamMQTT), nil
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
	case RECEIVER_TRANSPORT_PARAM_TYPE_WEBSOCKET:
		return fmt.Sprint(*p.ParamWebsocket)
	case RECEIVER_TRANSPORT_PARAM_TYPE_MQTT:
		return fmt.Sprint(*p.ParamMQTT)
	case RECEIVER_TRANSPORT_PARAM_TYPE_DASH:
		return fmt.Sprint(*p.ParamDash)
	default:
		return ""
	}
}
