package models

// AssignedLicense struct for AssignedLicense
type AssignedLicense struct {
	// A collection of the unique identifiers for plans that have been disabled.
	DisabledPlans *[]string `json:"disabledPlans,omitempty"`
	// The unique identifier for the SKU.
	SkuId *string `json:"skuId,omitempty"`
}
