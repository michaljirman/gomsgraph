package models

type MultiValueLegacyExtendedProperty struct {
	Id *string `json:"id,omitempty"`
	// A collection of property values.
	Value *[]string `json:"value,omitempty"`
}
