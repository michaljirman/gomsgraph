package models

import "time"

type List struct {
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
	// The displayable title of the list.
	DisplayName *string `json:"displayName,omitempty"`
	// Provides additional details about the list.
	List *ListInfo `json:"list,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	// If present, indicates that this is a system-managed list. Read-only.
	System *interface{} `json:"system,omitempty"`
	// The collection of field definitions for this list.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`
	// The collection of content types present in this list.
	ContentTypes *[]ContentType `json:"contentTypes,omitempty"`
	// Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
	Drive *Drive `json:"drive,omitempty"`
	// All items contained in the list.
	Items *[]ListItem `json:"items,omitempty"`
	// The set of subscriptions on the list.
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}

// ListInfo struct for ListInfo
type ListInfo struct {
	// If true, indicates that content types are enabled for this list.
	ContentTypesEnabled *bool `json:"contentTypesEnabled,omitempty"`
	// If true, indicates that the list is not normally visible in the SharePoint user experience.
	Hidden *bool `json:"hidden,omitempty"`
	// An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
	Template *string `json:"template,omitempty"`
}

// ColumnDefinition struct for ColumnDefinition
type ColumnDefinition struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// This column stores boolean values.
	Boolean *interface{} `json:"boolean,omitempty"`
	// This column's data is calculated based on other columns.
	Calculated *CalculatedColumn `json:"calculated,omitempty"`
	// This column stores data from a list of choices.
	Choice *ChoiceColumn `json:"choice,omitempty"`
	// For site columns, the name of the group this column belongs to. Helps organize related columns.
	ColumnGroup *string `json:"columnGroup,omitempty"`
	// This column stores currency values.
	Currency *CurrencyColumn `json:"currency,omitempty"`
	// This column stores DateTime values.
	DateTime *DateTimeColumn `json:"dateTime,omitempty"`
	// The default value for this column.
	DefaultValue *DefaultColumnValue `json:"defaultValue,omitempty"`
	// The user-facing description of the column.
	Description *string `json:"description,omitempty"`
	// The user-facing name of the column.
	DisplayName *string `json:"displayName,omitempty"`
	// If true, no two list items may have the same value for this column.
	EnforceUniqueValues *bool `json:"enforceUniqueValues,omitempty"`
	// This column stores a geolocation.
	Geolocation *interface{} `json:"geolocation,omitempty"`
	// Specifies whether the column is displayed in the user interface.
	Hidden *bool `json:"hidden,omitempty"`
	// Specifies whether the column values can used for sorting and searching.
	Indexed *bool `json:"indexed,omitempty"`
	// This column's data is looked up from another source in the site.
	Lookup *LookupColumn `json:"lookup,omitempty"`
	// The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
	Name *string `json:"name,omitempty"`
	// This column stores number values.
	Number *NumberColumn `json:"number,omitempty"`
	// This column stores Person or Group values.
	PersonOrGroup *PersonOrGroupColumn `json:"personOrGroup,omitempty"`
	// Specifies whether the column values can be modified.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Specifies whether the column value is not optional.
	Required *bool `json:"required,omitempty"`
	// This column stores text values.
	Text *TextColumn `json:"text,omitempty"`
}

type ContentType struct {
	Id *string `json:"id,omitempty"`
	// The descriptive text for the item.
	Description *string `json:"description,omitempty"`
	// The name of the group this content type belongs to. Helps organize related content types.
	Group *string `json:"group,omitempty"`
	// Indicates whether the content type is hidden in the list's 'New' menu.
	Hidden *bool `json:"hidden,omitempty"`
	// If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`
	// The name of the content type.
	Name *string `json:"name,omitempty"`
	// Specifies the order in which the content type appears in the selection UI.
	Order *ContentTypeOrder `json:"order,omitempty"`
	// The unique identifier of the content type.
	ParentId *string `json:"parentId,omitempty"`
	// If true, the content type cannot be modified unless this value is first set to false.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// If true, the content type cannot be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
	Sealed *bool `json:"sealed,omitempty"`
	// The collection of columns that are required by this content type
	ColumnLinks *[]ColumnLink `json:"columnLinks,omitempty"`
}

type ListItem struct {
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
	// The content type of this list item
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	// Analytics about the view activities that took place on this item.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// The values of the columns set on this list item.
	Fields *FieldValueSet `json:"fields,omitempty"`
	// The list of previous versions of the list item.
	Versions *[]ListItemVersion `json:"versions,omitempty"`
}

type Subscription struct {
	Id *string `json:"id,omitempty"`
	// Identifier of the application used to create the subscription. Read-only.
	ApplicationId *string `json:"applicationId,omitempty"`
	// Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
	ChangeType *string `json:"changeType,omitempty"`
	// Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
	ClientState *string `json:"clientState,omitempty"`
	// Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
	CreatorId *string `json:"creatorId,omitempty"`
	// A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true.
	EncryptionCertificate *string `json:"encryptionCertificate,omitempty"`
	// A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional.
	EncryptionCertificateId *string `json:"encryptionCertificateId,omitempty"`
	// Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// When set to true, change notifications include resource data (such as content of a chat message). Optional.
	IncludeResourceData *bool `json:"includeResourceData,omitempty"`
	// Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
	LatestSupportedTlsVersion *string `json:"latestSupportedTlsVersion,omitempty"`
	// The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications.
	LifecycleNotificationUrl *string `json:"lifecycleNotificationUrl,omitempty"`
	// Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
	NotificationUrl *string `json:"notificationUrl,omitempty"`
	// Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
	Resource *string `json:"resource,omitempty"`
}

type CalculatedColumn struct {
	// For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
	Format *string `json:"format,omitempty"`
	// The formula used to compute the value for this column.
	Formula *string `json:"formula,omitempty"`
	// The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
	OutputType *string `json:"outputType,omitempty"`
}

type ChoiceColumn struct {
	// If true, allows custom values that aren't in the configured choices.
	AllowTextEntry *bool `json:"allowTextEntry,omitempty"`
	// The list of values available for this column.
	Choices *[]string `json:"choices,omitempty"`
	// How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
	DisplayAs *string `json:"displayAs,omitempty"`
}

type CurrencyColumn struct {
	// Specifies the locale from which to infer the currency symbol.
	Locale *string `json:"locale,omitempty"`
}

type DateTimeColumn struct {
	// How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
	DisplayAs *string `json:"displayAs,omitempty"`
	// Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
	Format *string `json:"format,omitempty"`
}

type DefaultColumnValue struct {
	// The formula used to compute the default value for this column.
	Formula *string `json:"formula,omitempty"`
	// The direct value to use as the default value for this column.
	Value *string `json:"value,omitempty"`
}

type LookupColumn struct {
	// Indicates whether multiple values can be selected from the source.
	AllowMultipleValues *bool `json:"allowMultipleValues,omitempty"`
	// Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
	AllowUnlimitedLength *bool `json:"allowUnlimitedLength,omitempty"`
	// The name of the lookup source column.
	ColumnName *string `json:"columnName,omitempty"`
	// The unique identifier of the lookup source list.
	ListId *string `json:"listId,omitempty"`
	// If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
	PrimaryLookupColumnId *string `json:"primaryLookupColumnId,omitempty"`
}

type NumberColumn struct {
	// How many decimal places to display. See below for information about the possible values.
	DecimalPlaces *string `json:"decimalPlaces,omitempty"`
	// How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
	DisplayAs *string `json:"displayAs,omitempty"`
	// The maximum permitted value.
	Maximum *int `json:"maximum,omitempty"`
	// The minimum permitted value.
	Minimum *int `json:"minimum,omitempty"`
}

type PersonOrGroupColumn struct {
	// Indicates whether multiple values can be selected from the source.
	AllowMultipleSelection *bool `json:"allowMultipleSelection,omitempty"`
	// Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
	ChooseFromType *string `json:"chooseFromType,omitempty"`
	// How to display the information about the person or group chosen. See below.
	DisplayAs *string `json:"displayAs,omitempty"`
}

type TextColumn struct {
	// Whether to allow multiple lines of text.
	AllowMultipleLines *bool `json:"allowMultipleLines,omitempty"`
	// Whether updates to this column should replace existing text, or append to it.
	AppendChangesToExistingText *bool `json:"appendChangesToExistingText,omitempty"`
	// The size of the text box.
	LinesForEditing *int `json:"linesForEditing,omitempty"`
	// The maximum number of characters for the value.
	MaxLength *int `json:"maxLength,omitempty"`
	// The type of text being stored. Must be one of plain or richText
	TextType *string `json:"textType,omitempty"`
}

type ContentTypeOrder struct {
	// Whether this is the default Content Type
	Default *bool `json:"default,omitempty"`
	// Specifies the position in which the Content Type appears in the selection UI.
	Position *int `json:"position,omitempty"`
}

type ColumnLink struct {
	Id *string `json:"id,omitempty"`
	// The name of the column  in this content type.
	Name *string `json:"name,omitempty"`
}

type ContentTypeInfo struct {
	// The id of the content type.
	Id *string `json:"id,omitempty"`
	// The name of the content type.
	Name *string `json:"name,omitempty"`
}

type ItemAnalytics struct {
	Id                *string             `json:"id,omitempty"`
	AllTime           *ItemActivityStat   `json:"allTime,omitempty"`
	ItemActivityStats *[]ItemActivityStat `json:"itemActivityStats,omitempty"`
	LastSevenDays     *ItemActivityStat   `json:"lastSevenDays,omitempty"`
}

type FieldValueSet struct {
	Id *string `json:"id,omitempty"`
}

type ListItemVersion struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the version was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Indicates the publication status of this particular version. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`
	// A collection of the fields and values for this version of the list item.
	Fields *FieldValueSet `json:"fields,omitempty"`
}

type PublicationFacet struct {
	// The state of publication for this document. Either published or checkout. Read-only.
	Level *string `json:"level,omitempty"`
	// The unique identifier for the version that is visible to the current caller. Read-only.
	VersionId *string `json:"versionId,omitempty"`
}

type ItemActivityStat struct {
	Id *string `json:"id,omitempty"`
	// Statistics about the access actions in this interval. Read-only.
	Access *ItemActionStat `json:"access,omitempty"`
	// Statistics about the create actions in this interval. Read-only.
	Create *ItemActionStat `json:"create,omitempty"`
	// Statistics about the delete actions in this interval. Read-only.
	Delete *ItemActionStat `json:"delete,omitempty"`
	// Statistics about the edit actions in this interval. Read-only.
	Edit *ItemActionStat `json:"edit,omitempty"`
	// When the interval ends. Read-only.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Indicates that the statistics in this interval are based on incomplete data. Read-only.
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	// Indicates whether the item is 'trending.' Read-only.
	IsTrending *bool `json:"isTrending,omitempty"`
	// Statistics about the move actions in this interval. Read-only.
	Move *ItemActionStat `json:"move,omitempty"`
	// When the interval starts. Read-only.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Exposes the itemActivities represented in this itemActivityStat resource.
	Activities *[]ItemActivity `json:"activities,omitempty"`
}

type ItemActionStat struct {
	// The number of times the action took place. Read-only.
	ActionCount *int `json:"actionCount,omitempty"`
	// The number of distinct actors that performed the action. Read-only.
	ActorCount *int `json:"actorCount,omitempty"`
}

type IncompleteData struct {
	// The service does not have source data before the specified time.
	MissingDataBeforeDateTime *time.Time `json:"missingDataBeforeDateTime,omitempty"`
	// Some data was not recorded due to excessive activity.
	WasThrottled *bool `json:"wasThrottled,omitempty"`
}

type ItemActivity struct {
	Id *string `json:"id,omitempty"`
	// An item was accessed.
	Access *interface{} `json:"access,omitempty"`
	// Details about when the activity took place. Read-only.
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// Identity of who performed the action. Read-only.
	Actor *IdentitySet `json:"actor,omitempty"`
	// Exposes the driveItem that was the target of this activity.
	DriveItem *DriveItem `json:"driveItem,omitempty"`
}
