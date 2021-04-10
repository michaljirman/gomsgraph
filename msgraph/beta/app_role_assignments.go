package beta

import (
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
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
