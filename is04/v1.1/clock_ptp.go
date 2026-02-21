package is04v1_1

// Describes a clock referenced to PTP
type ClockPTP struct {
	Name      string          `json:"name"`      // Name of this refclock (unique for this set of clocks)
	RefType   ClockRefType    `json:"ref_type"`  // Type of external reference used by this clock
	Traceable bool            `json:"traceable"` // External refclock is synchronised to International Atomic Time (TAI)
	Version   ClockPTPVersion `json:"version"`   // Version of PTP reference used by this clock
	GMID      string          `json:"gmid"`      // ID of the PTP reference used by this clock
	Locked    bool            `json:"locked"`    // Lock state of this clock to the external reference. If true, this device is slaved, otherwise it has no defined relationship to the external reference
}

type ClockPTPVersion string

const (
	CLOCK_PTP_IEEE1588_2008 ClockPTPVersion = "IEEE1588-2008"
)
