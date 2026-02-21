package is04v1_0

// Register a new resource or update an existing resource
type RegistrationPost struct {
	Type RegistrationPostType `json:"type"` // Singular form of the resource type to be registered
	Data any                  `json:"data"`
}

type RegistrationPostType string

const (
	REGISTRATION_TYPE_NODE     RegistrationPostType = "node"
	REGISTRATION_TYPE_DEVICE   RegistrationPostType = "device"
	REGISTRATION_TYPE_SENDER   RegistrationPostType = "sender"
	REGISTRATION_TYPE_RECEIVER RegistrationPostType = "receiver"
	REGISTRATION_TYPE_SOURCE   RegistrationPostType = "source"
	REGISTRATION_TYPE_FLOW     RegistrationPostType = "flow"
)
