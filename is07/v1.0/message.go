package is07v1_0

import (
	"encoding/json"
	"fmt"
)

type MessageType string

const (
	MESSAGE_TYPE_EVENT             MessageType = "status"
	MESSAGE_TYPE_CONNECTION_STATUS MessageType = "connection_status"
	MESSAGE_TYPE_HEALTH            MessageType = "health"
	MESSAGE_TYPE_SHUTDOWN          MessageType = "shutdown"
	MESSAGE_TYPE_REBOOT            MessageType = "reboot"
)

// Allowed messages (communication from the sender to the receiver)
type Message struct {
	Type             MessageType
	Event            *Event
	ConnectionStatus *MessageConnectionStatus
	Health           *MessageHealth
	Shutdown         *MessageShutdownReboot
	Reboot           *MessageShutdownReboot
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	msgType, msgType_ok := dataTest["message_type"].(MessageType)
	if !msgType_ok {
		return fmt.Errorf("unable to parse message_type for data %v", data)
	}
	switch msgType {
	case MESSAGE_TYPE_EVENT:
		var parsed Event
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		m.Event = &parsed
		m.Type = MESSAGE_TYPE_EVENT
		return nil
	case MESSAGE_TYPE_CONNECTION_STATUS:
		var parsed MessageConnectionStatus
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		m.ConnectionStatus = &parsed
		m.Type = MESSAGE_TYPE_CONNECTION_STATUS
		return nil
	case MESSAGE_TYPE_HEALTH:
		var parsed MessageHealth
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		m.Health = &parsed
		m.Type = MESSAGE_TYPE_HEALTH
		return nil
	case MESSAGE_TYPE_SHUTDOWN:
		var parsed MessageShutdownReboot
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		m.Shutdown = &parsed
		m.Type = MESSAGE_TYPE_SHUTDOWN
		return nil
	case MESSAGE_TYPE_REBOOT:
		var parsed MessageShutdownReboot
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		m.Reboot = &parsed
		m.Type = MESSAGE_TYPE_REBOOT
		return nil
	default:
		return fmt.Errorf("unable to parse message type %s", msgType)
	}
}

func (m *Message) MarshalJSON() ([]byte, error) {
	switch m.Type {
	case MESSAGE_TYPE_EVENT:
		return json.Marshal(m.Event)
	case MESSAGE_TYPE_CONNECTION_STATUS:
		return json.Marshal(m.ConnectionStatus)
	case MESSAGE_TYPE_HEALTH:
		return json.Marshal(m.Health)
	case MESSAGE_TYPE_SHUTDOWN:
		return json.Marshal(m.Shutdown)
	case MESSAGE_TYPE_REBOOT:
		return json.Marshal(m.Reboot)
	default:
		return nil, fmt.Errorf("unable to parse type %s", m.Type)
	}
}

func (m Message) String() string {
	switch m.Type {
	case MESSAGE_TYPE_EVENT:
		return fmt.Sprint(m.Event)
	case MESSAGE_TYPE_CONNECTION_STATUS:
		return fmt.Sprint(m.ConnectionStatus)
	case MESSAGE_TYPE_HEALTH:
		return fmt.Sprint(m.Health)
	case MESSAGE_TYPE_SHUTDOWN:
		return fmt.Sprint(m.Shutdown)
	case MESSAGE_TYPE_REBOOT:
		return fmt.Sprint(m.Reboot)
	default:
		return ""
	}
}
