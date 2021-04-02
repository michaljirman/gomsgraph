package models

// Identity struct for Identity
type Identity struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id *string `json:"id,omitempty"`
}
