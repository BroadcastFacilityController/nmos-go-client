package is04v1_2

import (
	"encoding/json"
	"fmt"
)

// Interface for holding multiple types of clocks.
// Used only for unmarshalling JSON.
// Type points to the method that stores the actual type. All others are null.
type Clock struct {
	Type          ClockRefType
	ClockInternal *ClockInternal
	ClockPTP      *ClockPTP
}

func (c *Clock) UnmarshalJSON(data []byte) error {
	// Unmarshal to a map[string]interface{} to test
	var dataTest map[string]interface{}
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	clockType := dataTest["ref_type"].(string)
	switch ClockRefType(clockType) {
	case CLOCK_REF_INTERNAL:
		c.Type = CLOCK_REF_INTERNAL
		c.ClockInternal = &ClockInternal{}
		return json.Unmarshal(data, c.ClockInternal)
	case CLOCK_REF_PTP:
		c.Type = CLOCK_REF_PTP
		c.ClockPTP = &ClockPTP{}
		return json.Unmarshal(data, c.ClockPTP)
	default:
		return fmt.Errorf("unable to parse ref_type for data: %s", string(data))
	}
}

func (c *Clock) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case CLOCK_REF_PTP:
		return json.Marshal(c.ClockPTP)
	case CLOCK_REF_INTERNAL:
		return json.Marshal(c.ClockInternal)
	default:
		return nil, fmt.Errorf("could not marshal clock %v", *c)
	}
}

func (c Clock) String() string {
	switch c.Type {
	case CLOCK_REF_INTERNAL:
		return fmt.Sprint(*c.ClockInternal)
	case CLOCK_REF_PTP:
		return fmt.Sprint(*c.ClockPTP)
	default:
		return ""
	}
}

type ClockRefType string

const (
	CLOCK_REF_INTERNAL ClockRefType = "internal"
	CLOCK_REF_PTP      ClockRefType = "ptp"
)
