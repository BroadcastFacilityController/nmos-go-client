package is05v1_0

type BulkReceiverPost struct {
	ID     string        `json:"id"`     // ID of the target receiver to apply parameters to
	Params ReceiverStage `json:"params"` // Describes a receiver
}
