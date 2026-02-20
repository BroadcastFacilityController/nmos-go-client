package is04v1_2

// Register a new resource or update an existing resource
type QueryAPISubscriptionPost struct {
	MaxUpdateRateMS int                              `json:"max_update_rate_ms"` // Rate limiting for messages. Sets the minimum interval between consecutive websocket messages
	Persist         bool                             `json:"persist"`            // Whether to destroy the socket when the final client disconnects
	Secure          bool                             `json:"secure"`             // Whether to produce a secure websocket connection (wss://). NB: Default should be 'false' if the API is being presented via HTTP, and 'true' for HTTPS
	ResourcePath    QueryAPISubscriptionResourcePath `json:"resource_path"`      // HTTP resource path in the query API which this subscription relates to
	Params          any                              `json:"params"`             // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
}

type QueryAPISubscriptionResourcePath string

const (
	QUERY_NODES     QueryAPISubscriptionResourcePath = "/nodes"
	QUERY_DEVICES   QueryAPISubscriptionResourcePath = "/devices"
	QUERY_SOURCES   QueryAPISubscriptionResourcePath = "/sources"
	QUERY_FLOWS     QueryAPISubscriptionResourcePath = "/flows"
	QUERY_SENDERS   QueryAPISubscriptionResourcePath = "/senders"
	QUERY_RECEIVERS QueryAPISubscriptionResourcePath = "/receivers"
)
