package models

import (
	"time"
)

// Group struct for MicrosoftGraphGroup
type Group struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select.
	//AssignedLabels *[]AnyOfmicrosoftGraphAssignedLabel `json:"assignedLabels,omitempty"`
	// The licenses that are assigned to the group. Returned only on $select. Read-only.
	//AssignedLicenses *[]AnyOfmicrosoftGraphAssignedLicense `json:"assignedLicenses,omitempty"`
	// Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default.
	Classification *string `json:"classification,omitempty"`
	// App ID of the app used to create the group. Can be null for some groups. Returned by default. Read-only. Supports $filter.
	CreatedByAppId *string `json:"createdByAppId,omitempty"`
	// Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// An optional description for the group. Returned by default.
	Description *string `json:"description,omitempty"`
	// The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter and $orderby.
	DisplayName *string `json:"displayName,omitempty"`
	// Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter.
	GroupTypes *[]string `json:"groupTypes,omitempty"`
	// Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true).
	HasMembersWithLicenseErrors *bool `json:"hasMembersWithLicenseErrors,omitempty"`
	// Identifies the info segments assigned to the group. Returned by default.
	InfoCatalogs *[]string `json:"infoCatalogs,omitempty"`
	// Indicates whether this group can be assigned to an Azure Active Directory role or not.This property can only be set while creating the group and is immutable. Only Global Administrator and Privileged Role Administrator roles can set this property. For more information, see Using a group to manage Azure AD role assignmentsReturned by default.
	IsAssignableToRole *bool `json:"isAssignableToRole,omitempty"`
	// Indicates status of the group license assignment to all members of the group. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete. Returned only on $select. Read-only.
	//LicenseProcessingState NullableAnyOfmicrosoftGraphLicenseProcessingState `json:"licenseProcessingState,omitempty"`
	// The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter.
	Mail *string `json:"mail,omitempty"`
	// Specifies whether the group is mail-enabled. Returned by default.
	MailEnabled *bool `json:"mailEnabled,omitempty"`
	// The mail alias for the group, unique in the organization. This property must be specified when a group is created. These characters cannot be used in the mailNickName: @()/[]';:.<>,SPACE. Returned by default. Supports $filter.
	MailNickname *string `json:"mailNickname,omitempty"`
	MdmAppId     *string `json:"mdmAppId,omitempty"`
	// The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default.
	MembershipRule *string `json:"membershipRule,omitempty"`
	// Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default.
	MembershipRuleProcessingState *string `json:"membershipRuleProcessingState,omitempty"`
	// Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter.
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
	OnPremisesNetBiosName *string `json:"onPremisesNetBiosName,omitempty"`
	// Errors when using Microsoft synchronization product during provisioning. Returned by default.
	//OnPremisesProvisioningErrors *[]AnyOfmicrosoftGraphOnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Read-only.
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter.
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// The preferred data location for the group. For more information, see  OneDrive Online Multi-Geo. Returned by default.
	PreferredDataLocation *string `json:"preferredDataLocation,omitempty"`
	// The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example 'en-US'. Returned by default.
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required for filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter.
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`
	// Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
	RenewedDateTime *time.Time `json:"renewedDateTime,omitempty"`
	// Specifies the group behaviors that can be set for a Microsoft 365 group during creation. This can be set only as part of creation (POST). Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled. For more information, see Set Microsoft 365 group behaviors and provisioning options.
	ResourceBehaviorOptions *[]string `json:"resourceBehaviorOptions,omitempty"`
	// Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team. For more information, see Set Microsoft 365 group behaviors and provisioning options.
	ResourceProvisioningOptions *[]string `json:"resourceProvisioningOptions,omitempty"`
	// Specifies whether the group is a security group. Returned by default. Supports $filter.
	SecurityEnabled *bool `json:"securityEnabled,omitempty"`
	// Security identifier of the group, used in Windows scenarios. Returned by default.
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	// Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
	Theme *string `json:"theme,omitempty"`
	// Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. See group visibility options to learn more. Returned by default.
	Visibility *string `json:"visibility,omitempty"`
	//AccessType NullableAnyOfmicrosoftGraphGroupAccessType `json:"accessType,omitempty"`
	// Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	AllowExternalSenders *bool `json:"allowExternalSenders,omitempty"`
	// Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	AutoSubscribeNewMembers *bool `json:"autoSubscribeNewMembers,omitempty"`
	// true if the group is not displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups; false otherwise. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	HideFromAddressLists *bool `json:"hideFromAddressLists,omitempty"`
	// true if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web, false otherwise. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	HideFromOutlookClients *bool `json:"hideFromOutlookClients,omitempty"`
	IsFavorite             *bool `json:"isFavorite,omitempty"`
	// Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	IsSubscribedByMail *bool `json:"isSubscribedByMail,omitempty"`
	// Count of conversations that have been delivered one or more new posts since the signed-in user's last visit to the group. This property is the same as unseenCount. Returned only on $select.
	UnseenConversationsCount *int `json:"unseenConversationsCount,omitempty"`
	// Count of conversations that have received new posts since the signed-in user last visited the group. This property is the same as unseenConversationsCount.Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
	UnseenCount *int `json:"unseenCount,omitempty"`
	// Count of new posts that have been delivered to the group's conversations since the signed-in user's last visit to the group. Returned only on $select.
	UnseenMessagesCount *int `json:"unseenMessagesCount,omitempty"`
	// Describes the processing status for rules-based dynamic groups. The property is null for non-rule based dynamic groups or if the dynamic group processing has been paused. Returned only on $select. Supports $filter. Read-only.
	//MembershipRuleProcessingStatus NullableAnyOfmicrosoftGraphMembershipRuleProcessingStatus `json:"membershipRuleProcessingStatus,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	// Represents the app roles a group has been granted for an application.
	//AppRoleAssignments *[]MicrosoftGraphAppRoleAssignment `json:"appRoleAssignments,omitempty"`
	//// The user (or application) that created the group. Note: This is not set if the user is an administrator. Read-only.
	//CreatedOnBehalfOf NullableAnyOfmicrosoftGraphDirectoryObject `json:"createdOnBehalfOf,omitempty"`
	//// Endpoints for the group. Read-only. Nullable.
	//Endpoints *[]MicrosoftGraphEndpoint `json:"endpoints,omitempty"`
	//// Groups and administrative units that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable.
	//MemberOf *[]MicrosoftGraphDirectoryObject `json:"memberOf,omitempty"`
	//// Users, contacts, and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for security groups and mail-enabled security groups), DELETE (supported only for security groups) Read-only. Nullable.
	//Members *[]MicrosoftGraphDirectoryObject `json:"members,omitempty"`
	//// A list of group members with license errors from this group-based license assignment. Read-only.
	//MembersWithLicenseErrors *[]MicrosoftGraphDirectoryObject `json:"membersWithLicenseErrors,omitempty"`
	//// The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. HTTP Methods: GET (supported for all groups), POST (supported for security groups and mail-enabled security groups), DELETE (supported only for security groups) Read-only. Nullable.
	//Owners           *[]MicrosoftGraphDirectoryObject                 `json:"owners,omitempty"`
	//PermissionGrants *[]MicrosoftGraphResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	//// Settings that can govern this group's behavior, like whether members can invite guest users to the group. Nullable.
	//Settings           *[]MicrosoftGraphDirectorySetting `json:"settings,omitempty"`
	//TransitiveMemberOf *[]MicrosoftGraphDirectoryObject  `json:"transitiveMemberOf,omitempty"`
	//TransitiveMembers  *[]MicrosoftGraphDirectoryObject  `json:"transitiveMembers,omitempty"`
	//// The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
	//AcceptedSenders *[]MicrosoftGraphDirectoryObject `json:"acceptedSenders,omitempty"`
	//// The group's calendar. Read-only.
	//Calendar NullableAnyOfmicrosoftGraphCalendar `json:"calendar,omitempty"`
	//// The calendar view for the calendar. Read-only.
	//CalendarView *[]MicrosoftGraphEvent `json:"calendarView,omitempty"`
	//// The group's conversations.
	//Conversations *[]MicrosoftGraphConversation `json:"conversations,omitempty"`
	//// The group's events.
	//Events *[]MicrosoftGraphEvent `json:"events,omitempty"`
	//// The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
	//RejectedSenders *[]MicrosoftGraphDirectoryObject `json:"rejectedSenders,omitempty"`
	//// The group's conversation threads. Nullable.
	//Threads *[]MicrosoftGraphConversationThread `json:"threads,omitempty"`
	//// The group's default drive. Read-only.
	//Drive NullableAnyOfmicrosoftGraphDrive `json:"drive,omitempty"`
	//// The group's drives. Read-only.
	//Drives *[]MicrosoftGraphDrive `json:"drives,omitempty"`
	//// The list of SharePoint sites in this group. Access the default site with /sites/root.
	//Sites *[]MicrosoftGraphSite `json:"sites,omitempty"`
	//// The collection of open extensions defined for the group. Read-only. Nullable.
	//Extensions *[]MicrosoftGraphExtension `json:"extensions,omitempty"`
	//// The collection of lifecycle policies for this group. Read-only. Nullable.
	//GroupLifecyclePolicies *[]MicrosoftGraphGroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`
	//// Selective Planner services available to the group. Read-only. Nullable.
	//Planner NullableAnyOfmicrosoftGraphPlannerGroup `json:"planner,omitempty"`
	//// Read-only.
	//Onenote NullableAnyOfmicrosoftGraphOnenote `json:"onenote,omitempty"`
	//// The group's profile photo.
	//Photo NullableAnyOfmicrosoftGraphProfilePhoto `json:"photo,omitempty"`
	//// The profile photos owned by the group. Read-only. Nullable.
	//Photos *[]MicrosoftGraphProfilePhoto   `json:"photos,omitempty"`
	//Team   NullableAnyOfmicrosoftGraphTeam `json:"team,omitempty"`
}
