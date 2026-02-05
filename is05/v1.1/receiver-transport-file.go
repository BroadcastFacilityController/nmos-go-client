package is05v1_1

import "github.com/guregu/null/v6"

// Transport file parameters. 'data' and 'type' must both be strings or both be null
type ReceiverTransportFile struct {
	Data null.String `json:"data"` // Content of the transport file
	Type null.String `json:"type"` // IANA assigned media type for file (e.g application/sdp)
}
