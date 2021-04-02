package models

import "time"

// AssignedPlan struct for AssignedPlan
type AssignedPlan struct {
	// The date and time at which the plan was assigned; for example: 2013-01-02T19:32:30Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut.
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// The name of the service; for example, 'Exchange'.
	Service *string `json:"service,omitempty"`
	// A GUID that identifies the service plan.
	ServicePlanId *string `json:"servicePlanId,omitempty"`
}
