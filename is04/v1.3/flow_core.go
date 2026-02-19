package is04v1_3

// Describes a flow
type FlowCore struct {
	ResourceCore
	GrainRate FlowCoreGrainRate `json:"grain_rate"` // Number of Grains per second for this Flow. Must be an integer division of, or equal to the Grain rate specified by the parent Source. Grain rate matches the frame rate for video (see NMOS Content Model). Specified for periodic Flows only.
	SourceID  string            `json:"source_id"`  // Globally unique identifier for the Source which initially created the Flow. This attribute is used to ensure referential integrity by registry implementations (v1.0 only).
	DeviceID  string            `json:"device_id"`  // Globally unique identifier for the Device which initially created the Flow. This attribute is used to ensure referential integrity by registry implementations (v1.1 onwards).
	Parents   []string          `json:"parents"`    // Array of UUIDs representing the Flow IDs of Grains which came together to generate this Flow (may change over the lifetime of this Flow)
}

// Number of Grains per second for this Flow. Must be an integer division of, or equal to the Grain rate specified by the parent Source. Grain rate matches the frame rate for video (see NMOS Content Model). Specified for periodic Flows only.
type FlowCoreGrainRate struct {
	Numerator   int `json:"numerator"`   // Numerator
	Denominator int `json:"denominator"` // Denominator. Can be empty
}
