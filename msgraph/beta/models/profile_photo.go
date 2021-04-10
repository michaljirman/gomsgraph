package models

type ProfilePhoto struct {
	Id *string `json:"id,omitempty"`
	// The height of the photo. Read-only.
	Height *int `json:"height,omitempty"`
	// The width of the photo. Read-only.
	Width *int `json:"width,omitempty"`
}
