package models

type SingleValueLegacyExtendedProperty struct {
	Id *string `json:"id,omitempty"`
	// A property value.
	Value *string `json:"value,omitempty"`
}
