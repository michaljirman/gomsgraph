package models

// TodoTaskList struct for TodoTaskList
type TodoTaskList struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the task list.
	DisplayName *string `json:"displayName,omitempty"`
	// True if the user is owner of the given task list.
	IsOwner *bool `json:"isOwner,omitempty"`
	// True if the task list is shared with other users
	IsShared *bool `json:"isShared,omitempty"`
	// Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
	WellknownListName *WellknownListName `json:"wellknownListName,omitempty"`
	// The collection of open extensions defined for the task list. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// The tasks in this task list. Read-only. Nullable.
	Tasks *[]TodoTask `json:"tasks,omitempty"`
}
