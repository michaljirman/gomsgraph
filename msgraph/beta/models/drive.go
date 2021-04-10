package models

import "time"

type Drive struct {
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
	// Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
	DriveType *string `json:"driveType,omitempty"`
	// Optional. The user account that owns the drive. Read-only.
	Owner *IdentitySet `json:"owner,omitempty"`
	// Optional. Information about the drive's storage space quota. Read-only.
	Quota         *Quota         `json:"quota,omitempty"`
	SharePointIds *SharepointIds `json:"sharePointIds,omitempty"`
	// If present, indicates that this is a system-managed drive. Read-only.
	System *interface{} `json:"system,omitempty"`
	// The list of items the user is following. Only in OneDrive for Business.
	Following *[]DriveItem `json:"following,omitempty"`
	// All items contained in the drive. Read-only. Nullable.
	Items *[]DriveItem `json:"items,omitempty"`
	// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
	List *List `json:"list,omitempty"`
	// The root folder of the drive. Read-only.
	Root *DriveItem `json:"root,omitempty"`
	// Collection of common folders available in OneDrive. Read-only. Nullable.
	Special *[]DriveItem `json:"special,omitempty"`
}
