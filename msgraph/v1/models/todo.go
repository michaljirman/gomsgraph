package models

// Todo struct for Todo
type Todo struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The task lists in the users mailbox.
	Lists *[]TodoTaskList `json:"lists,omitempty"`
}
