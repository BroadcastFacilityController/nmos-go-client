package is05v1_1

// The constraints for a single transport parameter
type Constraint struct {
	Max         int    `json:"maximum,omitempty"`     // The inclusive maximum value the parameter can be set to
	Min         int    `json:"minimum,omitempty"`     // The inclusive minimum value the parameter can be set to
	Enum        []any  `json:"enum,omitempty"`        // An array of allowed values
	Pattern     string `json:"pattern,omitempty"`     // A regex pattern that must be satisfied for this parameter
	Description string `json:"description,omitempty"` // A human readable string describing the constraint (optional)
}
