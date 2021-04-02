package models

// DirectoryRole Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type DirectoryRole struct {
	Id *string `json:"id,omitempty"`
	// The description for the directory role. Read-only.
	Description *string `json:"description,omitempty"`
	// The display name for the directory role. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.
	RoleTemplateId *string `json:"roleTemplateId,omitempty"`
	// Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.
	Members       *[]DirectoryObject      `json:"members,omitempty"`
	ScopedMembers *[]ScopedRoleMembership `json:"scopedMembers,omitempty"`
}
