package is07v1_0

// Number
type Number struct {
	Value int `json:"value"`           // Number value
	Scale int `json:"scale,omitempty"` // Denominator for rational number represenation (default = 1)
}
