package is07v1_0

import "encoding/json"

type TypeBooleanEnum struct {
	Type   string                  `json:"type"`   // Base type name
	Values []TypeBooleanEnumOption `json:"values"` // List of allowed values
}

type TypeBooleanEnumOption struct {
	Value       bool   `json:"value"`       // Enum value
	Label       string `json:"label"`       // Enum label
	Description string `json:"description"` // Enum description
}

func (t *TypeBooleanEnum) MarshalJSON() ([]byte, error) {
	type alias *TypeBooleanEnum
	r := alias(t)
	r.Type = "boolean"
	return json.Marshal(r)
}
