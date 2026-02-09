package is08v1_0

import "github.com/guregu/null/v6"

// IO Information view for entire API
type IOResponse struct {
	Inputs  map[string]IOResponseInput  `json:"inputs"`
	Outputs map[string]IOResponseOutput `json:"outputs"`
}

type IOResponseInput struct {
	Properties InputProperties       `json:"properties"`
	Parent     InputParentResponse   `json:"parent"`
	Channels   InputChannelsResponse `json:"channels"`
	Caps       InputCapsResponse     `json:"caps"`
}

type IOResponseOutput struct {
	Properties OutputProperties       `json:"properties"`
	SourceID   null.String            `json:"source_id"`
	Channels   OutputChannelsResponse `json:"channels"`
	Caps       OutputCapsResponse     `json:"caps"`
}
