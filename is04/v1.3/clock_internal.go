package is04v1_3

// Describes a clock with no external reference
type ClockInternal struct {
	Name    string       `json:"name"`     // Name of this refclock (unique for this set of clocks)
	RefType ClockRefType `json:"ref_type"` // Type of external reference used by this clock
}
