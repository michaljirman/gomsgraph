package models

// ProfilePhoto struct for ProfilePhoto
type ProfilePhoto struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The height of the photo. Read-only.
	Height *int `json:"height,omitempty"`
	// The width of the photo. Read-only.
	Width *int `json:"width,omitempty"`
}
