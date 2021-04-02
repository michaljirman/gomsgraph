package models

import "time"

// AppRoleAssignment struct for AppRoleAssignment
type AppRoleAssignment struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create. Does not support $filter.
	AppRoleId       *string    `json:"appRoleId,omitempty"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith).
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	// The unique identifier (id) for the user, group or service principal being granted the app role. Required on create. Does not support $filter.
	PrincipalId *string `json:"principalId,omitempty"`
	// The type of the assigned principal. This can either be 'User', 'Group' or 'ServicePrincipal'. Read-only. Does not support $filter.
	PrincipalType *string `json:"principalType,omitempty"`
	// The display name of the resource app's service principal to which the assignment is made. Does not support $filter.
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only).
	ResourceId *string `json:"resourceId,omitempty"`
}
