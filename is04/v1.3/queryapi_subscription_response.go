package is04v1_3

// A single subscription resource registered with a Query API
type QueryAPISubscription struct {
	ID              string `json:"id"`                      // Globally unique identifier for the subscription
	WSHRef          string `json:"ws_href"`                 // Address to connect to for the websocket subscription
	MaxUpdateRateMS int    `json:"max_update_rate_ms"`      // Rate limiting for messages. Sets the minimum interval (in milliseconds) between consecutive WebSocket messages.
	Persist         bool   `json:"persist"`                 // Whether the API will retain or destroy the subscription after the final client disconnects
	Secure          bool   `json:"secure"`                  // Whether a secure WebSocket connection (wss://) is required.
	ResourcePath    string `json:"resource_path"`           // HTTP resource path in the Query API to which this subscription relates
	Params          any    `json:"params"`                  // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
	Authorization   bool   `json:"authorization,omitempty"` // Whether the WebSocket connection requires authorization.
}
