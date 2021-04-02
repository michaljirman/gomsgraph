package models

// ProvisionedPlan struct for ProvisionedPlan
type ProvisionedPlan struct {
	// For example, 'Enabled'.
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// For example, 'Success'.
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	// The name of the service; for example, 'AccessControlS2S'
	Service *string `json:"service,omitempty"`
}
