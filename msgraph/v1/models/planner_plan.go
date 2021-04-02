package models

// PlannerPlan struct for PlannerPlan
type PlannerPlan struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Read-only. The user who created the plan.
	//CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	//// Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	//CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	//// ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated.
	//Owner *string `json:"owner,omitempty"`
	//// Required. Title of the plan.
	//Title *string `json:"title,omitempty"`
	//// Read-only. Nullable. Collection of buckets in the plan.
	//Buckets *[]PlannerBucket `json:"buckets,omitempty"`
	//// Read-only. Nullable. Additional details about the plan.
	//Details *PlannerPlanDetails `json:"details,omitempty"`
	//// Read-only. Nullable. Collection of tasks in the plan.
	//Tasks *[]PlannerTask `json:"tasks,omitempty"`
}
