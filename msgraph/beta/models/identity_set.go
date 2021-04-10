package models

type IdentitySet struct {
	// Optional. The application associated with this action.
	Application *Identity `json:"application,omitempty"`
	// Optional. The device associated with this action.
	Device *Identity `json:"device,omitempty"`
	// Optional. The user associated with this action.
	User *Identity `json:"user,omitempty"`
}
