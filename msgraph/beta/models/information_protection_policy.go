package models

type InformationProtectionPolicy struct {
	Labels *[]InformationProtectionLabel `json:"labels,omitempty"`
}

type InformationProtectionLabel struct {
	// The color that the UI should display for the label, if configured.
	Color *string `json:"color,omitempty"`
	// The admin-defined description for the label.
	Description *string `json:"description,omitempty"`
	// Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
	IsActive *bool `json:"isActive,omitempty"`
	// The plaintext name of the label.
	Name   *string             `json:"name,omitempty"`
	Parent *ParentLabelDetails `json:"parent,omitempty"`
	// The sensitivity value of the label, where lower is less sensitive.
	Sensitivity *int32 `json:"sensitivity,omitempty"`
	// The tooltip that should be displayed for the label in a UI.
	Tooltip *string `json:"tooltip,omitempty"`
}

type ParentLabelDetails struct {
	Color       *string             `json:"color,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *string             `json:"id,omitempty"`
	IsActive    *bool               `json:"isActive,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Parent      *ParentLabelDetails `json:"parent,omitempty"`
	Sensitivity *int32              `json:"sensitivity,omitempty"`
	Tooltip     *string             `json:"tooltip,omitempty"`
}
