package models

// CalendarPermission struct for CalendarPermission
type CalendarPermission struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
	AllowedRoles *[]CalendarRoleType `json:"allowedRoles,omitempty"`
	// Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
	// True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
	IsInsideOrganization *bool `json:"isInsideOrganization,omitempty"`
	// True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// Current permission level of the calendar sharee or delegate.
	Role *CalendarRoleType `json:"role,omitempty"`
}
