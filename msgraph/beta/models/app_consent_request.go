package models

import (
	"time"
)

type AppConsentRequest struct {
	// The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// The identifier of the application. Required. Supports $filter (eq only) and $orderby.
	AppId *string `json:"appId,omitempty"`
	// The consent type of the request. Possible values are: Static and Dynamic. These represent static and dynamic permissions, respectively, requested in the consent workflow. Supports $filter (eq only) and $orderby. Required.
	ConsentType *string `json:"consentType,omitempty"`
	// A list of pending scopes waiting for approval. This is empty if the consentType is Static. Required.
	PendingScopes *[]AppConsentRequestScope `json:"pendingScopes,omitempty"`
	// A list of pending user consent requests.
	UserConsentRequests *[]UserConsentRequest `json:"userConsentRequests,omitempty"`
}

type AppConsentRequestScope struct {
	// The name of the scope.
	DisplayName *string `json:"displayName,omitempty"`
}

type UserConsentRequest struct {
	Id                *string      `json:"id,omitempty"`
	ApprovalId        *string      `json:"approvalId,omitempty"`
	CompletedDateTime *time.Time   `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime   *time.Time   `json:"createdDateTime,omitempty"`
	CustomData        *string      `json:"customData,omitempty"`
	Status            *string      `json:"status,omitempty"`
	// The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
	Reason *string `json:"reason,omitempty"`
	// Approval decisions associated with a request.
	Approval *Approval `json:"approval,omitempty"`
}

type Approval struct {
	Id    *string         `json:"id,omitempty"`
	Steps *[]ApprovalStep `json:"steps,omitempty"`
}

type ApprovalStep struct {
	Id *string `json:"id,omitempty"`
	// Indicates whether the step is assigned to the calling user to review. Read-only.
	AssignedToMe *bool `json:"assignedToMe,omitempty"`
	// The label provided by the policy creator to identify an approval step. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
	// The justification associated with the approval step decision.
	Justification *string `json:"justification,omitempty"`
	// The identifier of the reviewer. Read-only.
	ReviewedBy *Identity `json:"reviewedBy,omitempty"`
	// The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
	ReviewResult *string `json:"reviewResult,omitempty"`
	// The step status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
	Status *string `json:"status,omitempty"`
}
