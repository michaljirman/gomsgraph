package models

type LicenseAssignmentState struct {
	// The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only.
	AssignedByGroup *string `json:"assignedByGroup,omitempty"`
	// The service plans that are disabled in this assignment. Read-Only.
	DisabledPlans *[]string `json:"disabledPlans,omitempty"`
	// License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here.
	Error *string `json:"error,omitempty"`
	// The unique identifier for the SKU. Read-Only.
	SkuId *string `json:"skuId,omitempty"`
	// Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error.
	State *string `json:"state,omitempty"`
}
