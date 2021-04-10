package models

import "time"

type OnPremisesProvisioningError struct {
	// Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
	Category *string `json:"category,omitempty"`
	// The date and time at which the error occurred.
	OccurredDateTime *time.Time `json:"occurredDateTime,omitempty"`
	// Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
	PropertyCausingError *string `json:"propertyCausingError,omitempty"`
	// Value of the property causing the error.
	Value *string `json:"value,omitempty"`
}
