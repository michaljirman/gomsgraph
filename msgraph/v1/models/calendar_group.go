package models

// CalendarGroup struct for CalendarGroup
type CalendarGroup struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey *string `json:"changeKey,omitempty"`
	// The class identifier. Read-only.
	ClassId *string `json:"classId,omitempty"`
	// The group name.
	Name *string `json:"name,omitempty"`
	// The calendars in the calendar group. Navigation property. Read-only. Nullable.
	Calendars *[]Calendar `json:"calendars,omitempty"`
}
