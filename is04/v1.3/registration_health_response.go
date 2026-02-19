package is04v1_3

// Response to a request to update a resource's health
type RegistrationHealthResponse struct {
	Health string `json:"health"` // Timestamp indicating the time in seconds at which the server recorded the heartbeat
}
