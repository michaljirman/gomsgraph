package models

import "time"

// UserSettings struct for UserSettings
type UserSettings struct {
	// Read-only.
	Id                                                   *string           `json:"id,omitempty"`
	ContributionToContentDiscoveryAsOrganizationDisabled *bool             `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	ContributionToContentDiscoveryDisabled               *bool             `json:"contributionToContentDiscoveryDisabled,omitempty"`
	ShiftPreferences                                     *ShiftPreferences `json:"shiftPreferences,omitempty"`
}

// ShiftPreferences struct for ShiftPreferences
type ShiftPreferences struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Identity of the person who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Availability of the user to be scheduled for work and its recurrence pattern.
	Availability *[]*ShiftAvailability `json:"availability,omitempty"`
}
