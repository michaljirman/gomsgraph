package v1

import (
	. "github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

type AppRoleAssignmentsService service

type AppRoleAssignmentsResponse struct {
	OData
	AppRoleAssignments []*AppRoleAssignment `json:"value"`
}

type AppRoleAssignmentResponse struct {
	OData
	AppRoleAssignment
}
