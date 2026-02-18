package is05v1_1

import (
	"encoding/json"
	"fmt"
)

type ConstraintsType string

const (
	CONSTRAINTS_TYPE_RTP       ConstraintsType = "rtp"
	CONSTRAINTS_TYPE_WEBSOCKET ConstraintsType = "websocket"
	CONSTRAINTS_TYPE_MQTT      ConstraintsType = "mqtt"
)

// Generic entity to handle json marshalling / unmarshalling of different constraint types
type Constraints struct {
	Type                 ConstraintsType
	ConstraintsRTP       *ConstraintsRTP
	ConstraintsWebsocket *ConstraintsWebsocket
	ConstraintsMQTT      *ConstraintsMQTT
}

func (c *Constraints) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	// Check RTP
	_, rtpEnabled_ok := dataTest["rtp_enabled"]
	if rtpEnabled_ok {
		var dataTemp ConstraintsRTP
		err := json.Unmarshal(data, &dataTemp)
		if err != nil {
			return err
		}
		c.ConstraintsRTP = &dataTemp
		c.Type = CONSTRAINTS_TYPE_RTP
		return nil
	}
	// Check Websocket
	_, connectionURI_ok := dataTest["connection_uri"]
	if connectionURI_ok {
		var dataTemp ConstraintsWebsocket
		err := json.Unmarshal(data, &dataTemp)
		if err != nil {
			return err
		}
		c.ConstraintsWebsocket = &dataTemp
		c.Type = CONSTRAINTS_TYPE_WEBSOCKET
		return nil
	}
	// Check MQTT
	_, brokerTopic_ok := dataTest["broker_topic"]
	if brokerTopic_ok {
		var dataTemp ConstraintsMQTT
		err := json.Unmarshal(data, &dataTemp)
		if err != nil {
			return err
		}
		c.ConstraintsMQTT = &dataTemp
		c.Type = CONSTRAINTS_TYPE_MQTT
		return nil
	}
	return fmt.Errorf("unable to parse constraint type %v", dataTest)
}

func (c *Constraints) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case CONSTRAINTS_TYPE_RTP:
		return json.Marshal(c.ConstraintsRTP)
	case CONSTRAINTS_TYPE_WEBSOCKET:
		return json.Marshal(c.ConstraintsWebsocket)
	case CONSTRAINTS_TYPE_MQTT:
		return json.Marshal(c.ConstraintsMQTT)
	default:
		return nil, fmt.Errorf("unable to marshal type %v", c.Type)
	}
}

func (c Constraints) String() string {
	switch c.Type {
	case CONSTRAINTS_TYPE_RTP:
		return fmt.Sprint(*c.ConstraintsRTP)
	case CONSTRAINTS_TYPE_WEBSOCKET:
		return fmt.Sprint(*c.ConstraintsWebsocket)
	case CONSTRAINTS_TYPE_MQTT:
		return fmt.Sprint(*c.ConstraintsMQTT)
	default:
		return ""
	}
}
