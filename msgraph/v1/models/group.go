package models

import "time"

// Group struct for Group
type Group struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// The list of sensitivity label pairs (label ID, label name) associated with an Microsoft 365 group. Returned only on $select. Read-only.
	AssignedLabels *[]AssignedLabel `json:"assignedLabels,omitempty"`
	// The licenses that are assigned to the group. Returned only on $select. Read-only.
	AssignedLicenses *[]AssignedLicense `json:"assignedLicenses,omitempty"`
	// Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default.
	Classification *string `json:"classification,omitempty"`
	// Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// An optional description for the group. Returned by default.
	Description *string `json:"description,omitempty"`
	// The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter and $orderby.
	DisplayName *string `json:"displayName,omitempty"`
	// Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter.
	GroupTypes *[]string `json:"groupTypes,omitempty"`
	// Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example.
	HasMembersWithLicenseErrors *bool `json:"hasMembersWithLicenseErrors,omitempty"`
	// Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
	LicenseProcessingState *LicenseProcessingState `json:"licenseProcessingState,omitempty"`
	// The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter.
	Mail *string `json:"mail,omitempty"`
	// Specifies whether the group is mail-enabled. Returned by default.
	MailEnabled  *bool   `json:"mailEnabled,omitempty"`
	MailNickname *string `json:"mailNickname,omitempty"`
	// The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default.
	MembershipRule *string `json:"membershipRule,omitempty"`
	// Indicates whether the dynamic membership processing is on or paused. Possible values are 'On' or 'Paused'. Returned by default.
	MembershipRuleProcessingState *string                        `json:"membershipRuleProcessingState,omitempty"`
	OnPremisesDomainName          *string                        `json:"onPremisesDomainName,omitempty"`
	OnPremisesLastSyncDateTime    *time.Time                     `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesNetBiosName         *string                        `json:"onPremisesNetBiosName,omitempty"`
	OnPremisesProvisioningErrors  *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Read-only.
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter.
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// The preferred data location for the group. For more information, see  OneDrive Online Multi-Geo. Returned by default.
	PreferredDataLocation *string `json:"preferredDataLocation,omitempty"`
	// The preferred language for an Microsoft 365 group. Should follow ISO 639-1 Code; for example 'en-US'. Returned by default.
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter.
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`
	// Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	RenewedDateTime *time.Time `json:"renewedDateTime,omitempty"`
	// Specifies whether the group is a security group. Returned by default. Supports $filter.
	SecurityEnabled *bool `json:"securityEnabled,omitempty"`
	// Security identifier of the group, used in Windows scenarios. Returned by default.
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	// Specifies an Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
	Theme *string `json:"theme,omitempty"`
	// Specifies the visibility of a Microsoft 365 group. Possible values are: Private, Public, or Hiddenmembership; blank values are treated as public.  See group visibility options to learn more.Visibility can be set only when a group is created; it is not editable.Visibility is supported only for unified groups; it is not supported for security groups. Returned by default.
	Visibility *string `json:"visibility,omitempty"`
	// Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select.
	AllowExternalSenders *bool `json:"allowExternalSenders,omitempty"`
	// Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select.
	AutoSubscribeNewMembers *bool `json:"autoSubscribeNewMembers,omitempty"`
	// True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select.
	HideFromAddressLists *bool `json:"hideFromAddressLists,omitempty"`
	// True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select.
	HideFromOutlookClients *bool `json:"hideFromOutlookClients,omitempty"`
	// Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select.
	IsSubscribedByMail *bool `json:"isSubscribedByMail,omitempty"`
	// Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select.
	UnseenCount        *int                 `json:"unseenCount,omitempty"`
	IsArchived         *bool                `json:"isArchived,omitempty"`
	AppRoleAssignments *[]AppRoleAssignment `json:"appRoleAssignments,omitempty"`
	// The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`
	// Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable.
	Members *[]DirectoryObject `json:"members,omitempty"`
	// A list of group members with license errors from this group-based license assignment. Read-only.
	MembersWithLicenseErrors *[]DirectoryObject `json:"membersWithLicenseErrors,omitempty"`
	// The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable.
	Owners           *[]DirectoryObject                 `json:"owners,omitempty"`
	PermissionGrants *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	// Read-only. Nullable.
	Settings           *[]GroupSetting    `json:"settings,omitempty"`
	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`
	TransitiveMembers  *[]DirectoryObject `json:"transitiveMembers,omitempty"`
	// The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
	AcceptedSenders *[]DirectoryObject `json:"acceptedSenders,omitempty"`
	// The group's calendar. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`
	// The calendar view for the calendar. Read-only.
	CalendarView *[]Event `json:"calendarView,omitempty"`
	// The group's conversations.
	// TODO Conversations *[]Conversation `json:"conversations,omitempty"`
	// The group's calendar events.
	Events *[]Event `json:"events,omitempty"`
	// The group's profile photo
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// The profile photos owned by the group. Read-only. Nullable.
	Photos *[]ProfilePhoto `json:"photos,omitempty"`
	// The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
	RejectedSenders *[]DirectoryObject `json:"rejectedSenders,omitempty"`
	// The group's conversation threads. Nullable.
	// TODO Threads *[]ConversationThread `json:"threads,omitempty"`
	// The group's default drive. Read-only.
	Drive *Drive `json:"drive,omitempty"`
	// The group's drives. Read-only.
	Drives *[]Drive `json:"drives,omitempty"`
	// The list of SharePoint sites in this group. Access the default site with /sites/root.
	Sites *[]Site `json:"sites,omitempty"`
	// The collection of open extensions defined for the group. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// The collection of lifecycle policies for this group. Read-only. Nullable.
	GroupLifecyclePolicies *[]GroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`
	// Entry-point to Planner resource that might exist for a Unified Group.
	Planner *PlannerGroup `json:"planner,omitempty"`
	// Read-only.
	Onenote *Onenote `json:"onenote,omitempty"`
	Team    *Team    `json:"team,omitempty"`
}

// AssignedLabel struct for AssignedLabel
type AssignedLabel struct {
	// The display name of the label. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
	// The unique identifier of the label.
	LabelId *string `json:"labelId,omitempty"`
}

// LicenseProcessingState struct for LicenseProcessingState
type LicenseProcessingState struct {
	State *string `json:"state,omitempty"`
}

// ResourceSpecificPermissionGrant struct for ResourceSpecificPermissionGrant
type ResourceSpecificPermissionGrant struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	ClientAppId     *string    `json:"clientAppId,omitempty"`
	ClientId        *string    `json:"clientId,omitempty"`
	Permission      *string    `json:"permission,omitempty"`
	PermissionType  *string    `json:"permissionType,omitempty"`
	ResourceAppId   *string    `json:"resourceAppId,omitempty"`
}

// GroupSetting struct for GroupSetting
type GroupSetting struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Display name of this group of settings, which comes from the associated template.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique identifier for the template used to create this group of settings. Read-only.
	TemplateId *string `json:"templateId,omitempty"`
	// Collection of name value pairs. Must contain and set all the settings defined in the template.
	Values *[]SettingValue `json:"values,omitempty"`
}

// PlannerGroup struct for PlannerGroup
type PlannerGroup struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Read-only. Nullable. Returns the plannerPlans owned by the group.
	Plans *[]PlannerPlan `json:"plans,omitempty"`
}

// GroupLifecyclePolicy struct for GroupLifecyclePolicy
type GroupLifecyclePolicy struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon.
	AlternateNotificationEmails *string `json:"alternateNotificationEmails,omitempty"`
	// Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined.
	GroupLifetimeInDays *int `json:"groupLifetimeInDays,omitempty"`
	// The group type for which the expiration policy applies. Possible values are All, Selected or None.
	ManagedGroupTypes *string `json:"managedGroupTypes,omitempty"`
}

// SettingValue struct for SettingValue
type SettingValue struct {
	// Name of the setting (as defined by the groupSettingTemplate).
	Name *string `json:"name,omitempty"`
	// Value of the setting.
	Value *string `json:"value,omitempty"`
}
