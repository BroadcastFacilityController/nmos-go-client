package is04v1_1

// A single subscription resource registered with a Query API
type QueryAPISubscription struct {
	ID              string `json:"id"`                 // Globally unique identifier for the subscription
	WSHRef          string `json:"ws_href"`            // Address to connect to for the websocket subscription
	MaxUpdateRateMS int    `json:"max_update_rate_ms"` // Rate limiting for messages. Sets the minimum interval between consecutive websocket messages
	Persist         bool   `json:"persist"`            // Whether to destroy the socket when the final client disconnects
	Secure          bool   `json:"secure"`             // Whether to produce a secure websocket connection (wss://). NB: Default should be 'false' if the API is being presented via HTTP, and 'true' for HTTPS
	ResourcePath    string `json:"resource_path"`      // HTTP resource path in the query API which this subscription relates to
	Params          any    `json:"params"`             // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
}
