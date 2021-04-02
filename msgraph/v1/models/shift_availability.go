package models

// ShiftAvailability struct for ShiftAvailability
type ShiftAvailability struct {
	// Specifies the pattern for recurrence
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// The time slot(s) preferred by the user.
	TimeSlots *[]TimeRange `json:"timeSlots,omitempty"`
	// Specifies the time zone for the indicated time.
	TimeZone *string `json:"timeZone,omitempty"`
}
