package models

// PlannerUser struct for PlannerUser
type PlannerUser struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Read-only. Nullable. Returns the plannerTasks assigned to the user.
	//Plans *[]PlannerPlan `json:"plans,omitempty"`
	//// Read-only. Nullable. Returns the plannerPlans shared with the user.
	//Tasks *[]PlannerTask `json:"tasks,omitempty"`
}
