package models

// LicenseAssignmentState struct for LicenseAssignmentState
type LicenseAssignmentState struct {
	AssignedByGroup *string   `json:"assignedByGroup,omitempty"`
	DisabledPlans   *[]string `json:"disabledPlans,omitempty"`
	Error           *string   `json:"error,omitempty"`
	SkuId           *string   `json:"skuId,omitempty"`
	State           *string   `json:"state,omitempty"`
}
