package models

// MultiValueLegacyExtendedProperty struct for MultiValueLegacyExtendedProperty
type MultiValueLegacyExtendedProperty struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A collection of property values.
	Value *[]string `json:"value,omitempty"`
}
