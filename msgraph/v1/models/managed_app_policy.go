package models

import "time"

// ManagedAppPolicy struct for ManagedAppPolicy
type ManagedAppPolicy struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The date and time the policy was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The policy's description.
	Description *string `json:"description,omitempty"`
	// Policy display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Last time the policy was modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Version of the entity.
	Version *string `json:"version,omitempty"`
}
