package models

// ScopedRoleMembership struct for ScopedRoleMembership
type ScopedRoleMembership struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Unique identifier for the administrative unit that the directory role is scoped to
	AdministrativeUnitId *string `json:"administrativeUnitId,omitempty"`
	// Unique identifier for the directory role that the member is in.
	RoleId         *string   `json:"roleId,omitempty"`
	RoleMemberInfo *Identity `json:"roleMemberInfo,omitempty"`
}
