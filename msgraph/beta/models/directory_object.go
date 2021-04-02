package models

import "time"

// DirectoryObject struct for DirectoryObject
type DirectoryObject struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
}
