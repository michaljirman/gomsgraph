package models

import "time"

// ManagedAppOperation struct for ManagedAppOperation
type ManagedAppOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The operation name.
	DisplayName *string `json:"displayName,omitempty"`
	// The last time the app operation was modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The current state of the operation
	State *string `json:"state,omitempty"`
	// Version of the entity.
	Version *string `json:"version,omitempty"`
}
