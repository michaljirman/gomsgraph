package models

// EmailAddress struct for EmailAddress
type EmailAddress struct {
	// The email address of the person or entity.
	Address *string `json:"address,omitempty"`
	// The display name of the person or entity.
	Name *string `json:"name,omitempty"`
}
