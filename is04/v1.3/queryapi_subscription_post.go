package is04v1_3

// Register a new resource or update an existing resource
type QueryAPISubscriptionPost struct {
	MaxUpdateRateMS int                              `json:"max_update_rate_ms"`      // Rate limiting for messages. Sets the minimum interval (in milliseconds) between consecutive WebSocket messages.
	Persist         bool                             `json:"persist"`                 // Whether the API will retain or destroy the subscription after the final client disconnects
	Secure          bool                             `json:"secure"`                  // Whether a secure WebSocket connection (wss://) is required. NB: Default should be 'true' if the API is being presented via HTTPS, and 'false' otherwise.
	ResourcePath    QueryAPISubscriptionResourcePath `json:"resource_path"`           // HTTP resource path in the Query API to which this subscription relates
	Params          any                              `json:"params"`                  // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
	Authorization   bool                             `json:"authorization,omitempty"` // Whether the WebSocket connection requires authorization. NB: Default should be 'true' if the API requires authorization, and 'false' otherwise.
}

type QueryAPISubscriptionResourcePath string

const (
	QUERY_SUBSCRIPTION_RESOURCE_PATH_NODES     QueryAPISubscriptionResourcePath = "/nodes"
	QUERY_SUBSCRIPTION_RESOURCE_PATH_DEVICES   QueryAPISubscriptionResourcePath = "/devices"
	QUERY_SUBSCRIPTION_RESOURCE_PATH_SOURCES   QueryAPISubscriptionResourcePath = "/sources"
	QUERY_SUBSCRIPTION_RESOURCE_PATH_FLOWS     QueryAPISubscriptionResourcePath = "/flows"
	QUERY_SUBSCRIPTION_RESOURCE_PATH_SENDERS   QueryAPISubscriptionResourcePath = "/senders"
	QUERY_SUBSCRIPTION_RESOURCE_PATH_RECEIVERS QueryAPISubscriptionResourcePath = "/receivers"
)
