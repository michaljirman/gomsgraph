package models

// Presence struct for Presence
type Presence struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly.
	Activity *string `json:"activity,omitempty"`
	// The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
	Availability *string `json:"availability,omitempty"`
}
