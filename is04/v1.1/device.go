package is04v1_1

import "encoding/json"

// Describes a Device
type Device struct {
	ResourceCore
	Type      DeviceTypeURI   `json:"type"`      // Device type URN
	NodeID    string          `json:"node_id"`   // Globally unique identifier for the Node which initially created the Device
	Senders   []string        `json:"senders"`   // UUIDs of Senders attached to the Device
	Receivers []string        `json:"receivers"` // UUIDs of Receivers attached to the Device
	Controls  []DeviceControl `json:"controls"`  // Control endpoints exposed for the Device
}

type DeviceTypeURI URI

const (
	DEVICE_GENERIC  DeviceTypeURI = "urn:x-nmos:device:generic"
	DEVICE_PIPELINE DeviceTypeURI = "urn:x-nmos:device:pipeline"
)

type DeviceControl struct {
	HRef             string `json:"href,omitempty"`          // URL to reach a control endpoint, whether http or otherwise
	Type             string `json:"type"`                    // URN identifying the control format
	Authorization    bool   `json:"authorization,omitempty"` // This endpoint requires authorization, optional
	UseAuthorization bool   `json:"-"`                       // Whether or not to force the "authorization" key in JSON responses
}

func (d *DeviceControl) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias DeviceControl
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	d.HRef = dataTemp.HRef
	d.Type = dataTemp.Type
	d.Authorization = dataTemp.Authorization
	_, auth_ok := dataTest["authorization"]
	if auth_ok {
		d.UseAuthorization = true
	}
	return nil
}

func (d *DeviceControl) MarshalJSON() ([]byte, error) {
	if d.UseAuthorization {
		return json.Marshal(struct {
			HRef          string `json:"href,omitempty"`
			Type          string `json:"type"`
			Authorization bool   `json:"authorization"`
		}{
			HRef:          d.HRef,
			Type:          d.Type,
			Authorization: d.Authorization,
		})
	} else {
		return json.Marshal(struct {
			HRef string `json:"href,omitempty"`
			Type string `json:"type"`
		}{
			HRef: d.HRef,
			Type: d.Type,
		})
	}
}
