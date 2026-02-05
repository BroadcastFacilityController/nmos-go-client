package is05v1_1

// Describes a bulk sender update resource
type BulkSenderPost struct {
	ID     string `json:"id"` // ID of the target sender to apply parameters to
	Params Sender `json:"params"`
}
