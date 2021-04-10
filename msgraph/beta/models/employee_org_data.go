package models

// EmployeeOrgData struct for EmployeeOrgData
type EmployeeOrgData struct {
	// The cost center associated with the user. Returned only on $select. Supports $filter.
	CostCenter *string `json:"costCenter,omitempty"`
	// The name of the division in which the user works. Returned only on $select. Supports $filter.
	Division *string `json:"division,omitempty"`
}
