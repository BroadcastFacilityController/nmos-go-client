package is04v1_2

import (
	"encoding/json"

	"github.com/guregu/null/v6"
)

// Describes a Source
type SourceCore struct {
	ResourceCore
	GrainRate SourceCoreGrainRate `json:"grain_rate,omitempty"` // Maximum number of Grains per second for Flows derived from this Source. Corresponding Flow Grain rates may override this attribute. Grain rate matches the frame rate for video (see NMOS Content Model). Specified for periodic Sources only.
	Caps      any                 `json:"caps"`                 // Capabilities (not yet defined)
	DeviceID  string              `json:"device_id"`            // Globally unique identifier for the Device which initially created the Source. This attribute is used to ensure referential integrity by registry implementations.
	Parents   []string            `json:"parents"`              // Array of UUIDs representing the Source IDs of Grains which came together at the input to this Source (may change over the lifetime of this Source)
	ClockName null.String         `json:"clock_name"`           // Reference to clock in the originating Node
}

type SourceCoreGrainRate struct {
	Numerator   int `json:"numerator"`             // Numerator
	Denominator int `json:"denominator,omitempty"` // Denominator
}

func (r *SourceCoreGrainRate) UnmarshalJSON(data []byte) error {
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	_, denominator_ok := dataTest["denominator"]
	err = json.Unmarshal(data, r)
	if err != nil {
		return err
	}
	if !denominator_ok {
		r.Denominator = 1
	}
	return nil
}

func (r *SourceCoreGrainRate) MarshalJSON() ([]byte, error) {
	rCopy := *r
	// Omit field when default = 1
	if rCopy.Denominator == 1 {
		rCopy.Denominator = 0
	}
	return json.Marshal(rCopy)
}
