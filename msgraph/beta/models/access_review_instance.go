package models

import (
	"time"
)

type AccessReviewInstance struct {
	Id *string `json:"id,omitempty"`
	// DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. In the case of a single-group review, the scope defined at the accessReviewScheduleDefinition level applies to all instances. In the case of all groups review, scope may be different for each group. Read-only.
	Scope *interface{} `json:"scope,omitempty"`
	// DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Read-only.
	Status *string `json:"status,omitempty"`
	// Each user reviewed in an accessReviewInstance has a decision item representing if their access was approved, denied, or not yet reviewed.
	Decisions *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
	// There is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
	Definition *AccessReviewScheduleDefinition `json:"definition,omitempty"`
}

type MicrosoftGraphAccessReviewInstanceDecisionItem struct {
	Id *string `json:"id,omitempty"`
	// The identifier of the accessReviewInstance parent.
	AccessReviewId *string `json:"accessReviewId,omitempty"`
	// The identifier of the user who applied the decision.
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`
	// The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AppliedDateTime *time.Time `json:"appliedDateTime,omitempty"`
	// The result of applying the decision. Possible values: NotApplied, Success, Failed, NotFound, or NotSupported.
	ApplyResult *string `json:"applyResult,omitempty"`
	// Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow.
	Decision *string `json:"decision,omitempty"`
	// The review decision justification.
	Justification *string `json:"justification,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity.
	Principal     *Identity `json:"principal,omitempty"`
	PrincipalLink *string   `json:"principalLink,omitempty"`
	// A system-generated recommendation for the approval decision. Possible values: Approve, Deny, or NotAvailable.
	Recommendation *string `json:"recommendation,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource
	Resource     *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	ResourceLink *string                                   `json:"resourceLink,omitempty"`
	// The identifier of the reviewer.
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`
	// The timestamp when the review occurred.
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// The target of this specific decision. Decision targets can be of different types – each one with its own specific properties. See accessReviewInstanceDecisionItemTarget.
	Target *interface{} `json:"target,omitempty"`
}

type UserIdentity struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id *string `json:"id,omitempty"`
	// Indicates the client IP address used by user performing the activity (audit log only).
	IpAddress *string `json:"ipAddress,omitempty"`
	// The userPrincipalName attribute of the user.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

type AccessReviewInstanceDecisionItemResource struct {
	// Display name of the resource
	DisplayName *string `json:"displayName,omitempty"`
	// Resource ID
	Id *string `json:"id,omitempty"`
	// Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
	Type *string `json:"type,omitempty"`
}

type AccessReviewInstanceDecisionItem struct {
	Id *string `json:"id,omitempty"`
	// The identifier of the accessReviewInstance parent.
	AccessReviewId *string `json:"accessReviewId,omitempty"`
	// The identifier of the user who applied the decision.
	AppliedBy *Identity `json:"appliedBy,omitempty"`
	// The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AppliedDateTime *time.Time `json:"appliedDateTime,omitempty"`
	// The result of applying the decision. Possible values: NotApplied, Success, Failed, NotFound, or NotSupported.
	ApplyResult *string `json:"applyResult,omitempty"`
	// Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow.
	Decision *string `json:"decision,omitempty"`
	// The review decision justification.
	Justification *string `json:"justification,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity.
	Principal     *Identity `json:"principal,omitempty"`
	PrincipalLink *string   `json:"principalLink,omitempty"`
	// A system-generated recommendation for the approval decision. Possible values: Approve, Deny, or NotAvailable.
	Recommendation *string `json:"recommendation,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource
	Resource     *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	ResourceLink *string                                   `json:"resourceLink,omitempty"`
	// The identifier of the reviewer.
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`
	// The timestamp when the review occurred.
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// The target of this specific decision. Decision targets can be of different types – each one with its own specific properties. See accessReviewInstanceDecisionItemTarget.
	Target *interface{} `json:"target,omitempty"`
}

type AccessReviewScheduleDefinition struct {
	Id *string `json:"id,omitempty"`
	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope.
	BackupReviewers *[]AccessReviewReviewerScope `json:"backupReviewers,omitempty"`
	// User who created this review.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// Timestamp when review series was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description provided by review creators to provide more context of the review to admins.
	DescriptionForAdmins *string `json:"descriptionForAdmins,omitempty"`
	// Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review.
	DescriptionForReviewers *string `json:"descriptionForReviewers,omitempty"`
	// Name of access review series. Required on create.
	DisplayName *string `json:"displayName,omitempty"`
	// In the case of a review of guest users across all Microsoft 365 groups, this determines the scope of which groups will be reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope.
	InstanceEnumerationScope *interface{} `json:"instanceEnumerationScope,omitempty"`
	// Timestamp when review series was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// This collection of access review scopes is used to define who are the reviewers. See accessReviewReviewerScope. Required on create.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`
	// Defines scope of users reviewed. For supported scopes, see accessReviewScope. Required on create.
	Scope *interface{} `json:"scope,omitempty"`
	// The settings for an access review series, see type definition below.
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`
	// This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
	Status *string `json:"status,omitempty"`
	// Set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there will be an instance for each recurrence.
	Instances *[]*AccessReviewInstance `json:"instances,omitempty"`
}

type AccessReviewReviewerScope struct {
	// The query specifying who will be the reviewer. See table for examples.
	Query *string `json:"query,omitempty"`
	// In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query (i.e., ./manager) is specified.
	QueryRoot *string `json:"queryRoot,omitempty"`
	// The type of query. Examples include MicrosoftGraph and ARM.
	QueryType *string `json:"queryType,omitempty"`
}

type AccessReviewScheduleSettings struct {
	// Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction. See accessReviewApplyAction.
	ApplyActions *[]interface{} `json:"applyActions,omitempty"`
	// Flag to indicate whether auto-apply feature is enabled.
	AutoApplyDecisionsEnabled *bool `json:"autoApplyDecisionsEnabled,omitempty"`
	// Decision chosen if defaultDecisionEnabled is enabled. Can be one of 'Approve', 'Deny', or 'Recommendation'.
	DefaultDecision *string `json:"defaultDecision,omitempty"`
	// Flag to indicate whether default decision is enabled/disabled when reviewers do not respond.
	DefaultDecisionEnabled *bool `json:"defaultDecisionEnabled,omitempty"`
	// Duration of each recurrence of review (accessReviewInstance) in number of days.
	InstanceDurationInDays *int32 `json:"instanceDurationInDays,omitempty"`
	// Flag to indicate whether reviewers are required to provide justification with their decision.
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`
	// Flag to indicate whether emails are enabled/disabled.
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`
	// Flag to indicate whether decision recommendations are enabled/disabled.
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`
	// Detailed settings for recurrence. Using standard Outlook recurrence object. Note that dayOfMonth is not supported - use property startDate on recurrenceRange to determine the day the review will start on.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// Flag to indicate whether reminders are enabled/disabled.
	ReminderNotificationsEnabled *bool `json:"reminderNotificationsEnabled,omitempty"`
}
