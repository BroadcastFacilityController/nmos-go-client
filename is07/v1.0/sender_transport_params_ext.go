package is07v1_0

import "github.com/guregu/null/v6"

// Describes external Sender transport parameters defined for IS-07 NMOS Event & Tally specification. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint.
type SenderTransportParamsExt struct {
	ExtIS07RestAPIUrl null.String `json:"ext_is_07_rest_api_url"` // the URL for the Events API resource for the associated source
	ExtIS07SourceID   null.String `json:"ext_is_07_source_id"`    // ID of the related source
}
