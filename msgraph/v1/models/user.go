package models

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// User struct for User
type User struct {
	// Read-only.
	Id              *string    `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// true if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information.
	AgeGroup *string `json:"ageGroup,omitempty"`
	// The licenses that are assigned to the user. Not nullable.
	AssignedLicenses *[]AssignedLicense `json:"assignedLicenses,omitempty"`
	// The plans that are assigned to the user. Read-only. Not nullable.
	AssignedPlans *[]AssignedPlan `json:"assignedPlans,omitempty"`
	// The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`
	// The city in which the user is located. Supports $filter.
	City *string `json:"city,omitempty"`
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 chararcters.Returned only on $select.
	CompanyName *string `json:"companyName,omitempty"`
	// Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information.
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty"`
	// The country/region in which the user is located; for example, 'US' or 'UK'. Supports $filter.
	Country *string `json:"country,omitempty"`
	// The created date of the user object.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Indicates whether the user account was created as a regular school or work account (null), an external account (Invitation), a local account for an Azure Active Directory B2C tenant (LocalAccount) or self-service sign-up using email verification (EmailVerified). Read-only.
	CreationType *string `json:"creationType,omitempty"`
	// The name for the department in which the user works. Supports $filter.
	Department *string `json:"department,omitempty"`
	// The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $filter and $orderby.
	DisplayName *string `json:"displayName,omitempty"`
	// The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter.
	EmployeeHireDate *time.Time `json:"employeeHireDate,omitempty"`
	// The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter.
	EmployeeId *string `json:"employeeId,omitempty"`
	// Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select.
	EmployeeOrgData *EmployeeOrgData `json:"employeeOrgData,omitempty"`
	// Captures enterprise worker type: Employee, Contractor, Consultant, Vendor, etc. Returned only on $select. Supports $filter.
	EmployeeType *string `json:"employeeType,omitempty"`
	// For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter with the supported values. For example: $filter=externalUserState eq 'PendingAcceptance'.
	ExternalUserState *string `json:"externalUserState,omitempty"`
	// Shows the timestamp for the latest change to the externalUserState property. Returned only on $select.
	ExternalUserStateChangeDateTime *time.Time `json:"externalUserStateChangeDateTime,omitempty"`
	// The fax number of the user.
	FaxNumber *string `json:"faxNumber,omitempty"`
	// The given name (first name) of the user. Returned by default. Supports $filter.
	GivenName *string `json:"givenName,omitempty"`
	// Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Supports $filter.
	Identities *[]ObjectIdentity `json:"identities,omitempty"`
	// The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only.
	ImAddresses *[]string `json:"imAddresses,omitempty"`
	// Do not use – reserved for future use.
	IsResourceAccount *bool `json:"isResourceAccount,omitempty"`
	// The user's job title. Returned by default. Supports $filter.
	JobTitle *string `json:"jobTitle,omitempty"`
	// The time when this Azure AD user last changed their password. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastPasswordChangeDateTime *time.Time `json:"lastPasswordChangeDateTime,omitempty"`
	// Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information.
	LegalAgeGroupClassification *string `json:"legalAgeGroupClassification,omitempty"`
	// State of license assignments for this user. Read-only.
	LicenseAssignmentStates *[]LicenseAssignmentState `json:"licenseAssignmentStates,omitempty"`
	// The SMTP address for the user, for example, 'jeff@contoso.onmicrosoft.com'. Returned by default. Supports $filter and endsWith.
	Mail *string `json:"mail,omitempty"`
	// The mail alias for the user. This property must be specified when a user is created. Supports $filter.
	MailNickname *string `json:"mailNickname,omitempty"`
	// The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Returned by default.
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// The office location in the user's place of business. Returned by default.
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
	OnPremisesDistinguishedName *string `json:"onPremisesDistinguishedName,omitempty"`
	// Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// Contains extensionAttributes 1-15 for the user. Note that the individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties may be set during creation or update. These extension attributes are also known as Exchange custom attributes 1-15.
	OnPremisesExtensionAttributes *OnPremisesExtensionAttributes `json:"onPremisesExtensionAttributes,omitempty"`
	// This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. Important: The $ and _ characters cannot be used when specifying this property. Supports $filter.
	OnPremisesImmutableId *string `json:"onPremisesImmutableId,omitempty"`
	// Indicates the last time at which the object was synced with the on-premises directory; for example: '2013-02-16T03:04:54Z'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only.
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// Errors when using Microsoft synchronization product during provisioning.
	OnPremisesProvisioningErrors *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only.
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
	OnPremisesUserPrincipalName *string `json:"onPremisesUserPrincipalName,omitempty"`
	// A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. Supports $filter.
	OtherMails *[]string `json:"otherMails,omitempty"`
	// Specifies password policies for the user. This value is an enumeration with one possible value being 'DisableStrongPassword', which allows weaker passwords than the default policy to be specified. 'DisablePasswordExpiration' can also be specified. The two may be specified together; for example: 'DisablePasswordExpiration, DisableStrongPassword'.
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The preferred language for the user. Should follow ISO 639-1 Code; for example 'en-US'. Returned by default.
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// The plans that are provisioned for the user. Read-only. Not nullable.
	ProvisionedPlans *[]ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'] The any operator is required for filter expressions on multi-valued properties. Read-only, Not nullable. Supports $filter.
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`
	// true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset.
	SignInSessionsValidFromDateTime *time.Time `json:"signInSessionsValidFromDateTime,omitempty"`
	// The state or province in the user's address. Supports $filter.
	State *string `json:"state,omitempty"`
	// The street address of the user's place of business.
	StreetAddress *string `json:"streetAddress,omitempty"`
	// The user's surname (family name or last name). Returned by default. Supports $filter.
	Surname *string `json:"surname,omitempty"`
	// A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: 'US', 'JP', and 'GB'. Not nullable. Supports $filter.
	UsageLocation *string `json:"usageLocation,omitempty"`
	// The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Returned by default. Supports $filter, $orderby, and endsWith.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports $filter.
	UserType *string `json:"userType,omitempty"`
	// Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.
	MailboxSettings *MailboxSettings `json:"mailboxSettings,omitempty"`
	// The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
	DeviceEnrollmentLimit *int32 `json:"deviceEnrollmentLimit,omitempty"`
	// A freeform text entry field for the user to describe themselves.
	AboutMe *string `json:"aboutMe,omitempty"`
	// The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	Birthday *time.Time `json:"birthday,omitempty"`
	// The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
	HireDate *time.Time `json:"hireDate,omitempty"`
	// A list for the user to describe their interests.
	Interests *[]string `json:"interests,omitempty"`
	// The URL for the user's personal site.
	MySite *string `json:"mySite,omitempty"`
	// A list for the user to enumerate their past projects.
	PastProjects *[]string `json:"pastProjects,omitempty"`
	// The preferred name for the user.
	PreferredName *string `json:"preferredName,omitempty"`
	// A list for the user to enumerate their responsibilities.
	Responsibilities *[]string `json:"responsibilities,omitempty"`
	// A list for the user to enumerate the schools they have attended.
	Schools *[]string `json:"schools,omitempty"`
	// A list for the user to enumerate their skills.
	Skills             *[]string            `json:"skills,omitempty"`
	AppRoleAssignments *[]AppRoleAssignment `json:"appRoleAssignments,omitempty"`
	// Directory objects that were created by the user. Read-only. Nullable.
	CreatedObjects *[]DirectoryObject `json:"createdObjects,omitempty"`
	// The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable.
	DirectReports *[]DirectoryObject `json:"directReports,omitempty"`
	// A collection of this user's license details. Read-only.
	LicenseDetails *[]LicenseDetails `json:"licenseDetails,omitempty"`
	// The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.)
	Manager *DirectoryObject `json:"manager,omitempty"`
	// The groups and directory roles that the user is a member of. Read-only. Nullable.
	MemberOf               *[]DirectoryObject       `json:"memberOf,omitempty"`
	Oauth2PermissionGrants *[]OAuth2PermissionGrant `json:"oauth2PermissionGrants,omitempty"`
	// Devices that are owned by the user. Read-only. Nullable.
	OwnedDevices *[]DirectoryObject `json:"ownedDevices,omitempty"`
	// Directory objects that are owned by the user. Read-only. Nullable.
	OwnedObjects *[]DirectoryObject `json:"ownedObjects,omitempty"`
	// Devices that are registered for the user. Read-only. Nullable.
	RegisteredDevices  *[]DirectoryObject      `json:"registeredDevices,omitempty"`
	ScopedRoleMemberOf *[]ScopedRoleMembership `json:"scopedRoleMemberOf,omitempty"`
	TransitiveMemberOf *[]DirectoryObject      `json:"transitiveMemberOf,omitempty"`
	// The user's primary calendar. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`
	// The user's calendar groups. Read-only. Nullable.
	CalendarGroups *[]CalendarGroup `json:"calendarGroups,omitempty"`
	// The user's calendars. Read-only. Nullable.
	Calendars *[]Calendar `json:"calendars,omitempty"`
	// The calendar view for the calendar. Read-only. Nullable.
	CalendarView *[]Event `json:"calendarView,omitempty"`
	// The user's contacts folders. Read-only. Nullable.
	ContactFolders *[]ContactFolder `json:"contactFolders,omitempty"`
	// The user's contacts. Read-only. Nullable.
	Contacts *[]Contact `json:"contacts,omitempty"`
	// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
	Events *[]Event `json:"events,omitempty"`
	// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
	InferenceClassification *InferenceClassification `json:"inferenceClassification,omitempty"`
	// The user's mail folders. Read-only. Nullable.
	MailFolders *[]MailFolder `json:"mailFolders,omitempty"`
	// The messages in a mailbox or folder. Read-only. Nullable.
	Messages *[]Message `json:"messages,omitempty"`
	// Read-only.
	Outlook *OutlookUser `json:"outlook,omitempty"`
	// People that are relevant to the user. Read-only. Nullable.
	People *[]Person `json:"people,omitempty"`
	// The user's profile photo. Read-only.
	Photo  *ProfilePhoto   `json:"photo,omitempty"`
	Photos *[]ProfilePhoto `json:"photos,omitempty"`
	// The user's OneDrive. Read-only.
	Drive *Drive `json:"drive,omitempty"`
	// A collection of drives available for this user. Read-only.
	Drives        *[]Drive `json:"drives,omitempty"`
	FollowedSites *[]Site  `json:"followedSites,omitempty"`
	// The collection of open extensions defined for the user. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// The managed devices associated with the user.
	ManagedDevices *[]ManagedDevice `json:"managedDevices,omitempty"`
	// Zero or more managed app registrations that belong to the user.
	ManagedAppRegistrations *[]ManagedAppRegistration `json:"managedAppRegistrations,omitempty"`
	// The list of troubleshooting events for this user.
	DeviceManagementTroubleshootingEvents *[]DeviceManagementTroubleshootingEvent `json:"deviceManagementTroubleshootingEvents,omitempty"`
	// Entry-point to the Planner resource that might exist for a user. Read-only.
	Planner *PlannerUser `json:"planner,omitempty"`
	// Read-only. Nullable.
	Insights *OfficeGraphInsights `json:"insights,omitempty"`
	Settings *UserSettings        `json:"settings,omitempty"`
	// Read-only.
	Onenote *Onenote `json:"onenote,omitempty"`
	// The user's activities across devices. Read-only. Nullable.
	Activities     *[]UserActivity  `json:"activities,omitempty"`
	OnlineMeetings *[]OnlineMeeting `json:"onlineMeetings,omitempty"`
	Presence       *Presence        `json:"presence,omitempty"`
	JoinedTeams    *[]Team          `json:"joinedTeams,omitempty"`
	Teamwork       *UserTeamwork    `json:"teamwork,omitempty"`
	// Represents the To Do services available to a user.
	Todo           *Todo                      `json:"todo,omitempty"`
	AdditionalData map[string]json.RawMessage `json:"-"`
}

// To avoid json unmarshall recursion
type _user User

// UnmarshalJSON unmarshal json data into UserResponse struct.
// Any unknown data will not be thrown away but placed into AdditionalData. This could be useful when selecting schema extensions attributes.
func (user *User) UnmarshalJSON(data []byte) error {
	var u _user
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	*user = User(u)

	objValue := reflect.ValueOf(user).Elem()
	knownFields := map[string]reflect.Value{}
	for i := 0; i != objValue.NumField(); i++ {
		jsonName := strings.Split(objValue.Type().Field(i).Tag.Get("json"), ",")[0]
		knownFields[jsonName] = objValue.Field(i)
	}

	otherFields := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &otherFields); err != nil {
		return err
	}

	for key := range otherFields {
		if _, found := knownFields[key]; found {
			delete(otherFields, key)
		}
		if _, found := ODataFields[key]; found {
			delete(otherFields, key)
		}
	}
	user.AdditionalData = otherFields

	return nil
}
