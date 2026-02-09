package is08v1_0

import "github.com/guregu/null/v6"

// Describes the map table entries
// Key is output unique identifier
type MapEntries map[string]MapEntryChannel

// Indexes of channels
// Key is channel on the output
type MapEntryChannel map[string]MapEntry

// Unique identifiers of outputs
type MapEntry struct {
	Input        null.String `json:"input"`         // Unique identifier of the input to which the channel to be routed belongs. null for unrouted.
	ChannelIndex null.Int    `json:"channel_index"` // Index of the channel to be routed. null for unrouted.
}
