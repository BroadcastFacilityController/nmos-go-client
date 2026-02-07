package is07v1_0

import "encoding/json"

type TypeNumber struct {
	Type  string  `json:"type"`            // Base type name
	Scale int     `json:"scale,omitempty"` // Denominator typically used for rational numbers in the event payload
	Min   Number  `json:"min"`
	Max   Number  `json:"max"`
	Step  *Number `json:"step,omitempty"` // Pointer to not automatically
	Unit  string  `json:"unit,omitempty"` // Unit of measure
}

func (t *TypeNumber) MarshalJSON() ([]byte, error) {
	type alias *TypeNumber
	r := alias(t)
	r.Type = "number"
	return json.Marshal(r)
}
