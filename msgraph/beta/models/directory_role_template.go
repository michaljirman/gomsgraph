package models

import (
	"time"
)

// DirectoryRoleTemplate Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type DirectoryRoleTemplate struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// The description to set for the directory role. Read-only.
	Description *string `json:"description,omitempty"`
	// The display name to set for the directory role. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
}
