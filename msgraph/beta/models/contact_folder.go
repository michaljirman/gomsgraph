package models

type ContactFolder struct {
	Id *string `json:"id,omitempty"`
	// The folder's display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the folder's parent folder.
	ParentFolderId *string `json:"parentFolderId,omitempty"`
	// The collection of child folders in the folder. Navigation property. Read-only. Nullable.
	ChildFolders *[]ContactFolder `json:"childFolders,omitempty"`
	// The contacts in the folder. Navigation property. Read-only. Nullable.
	Contacts *[]Contact `json:"contacts,omitempty"`
	// The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
