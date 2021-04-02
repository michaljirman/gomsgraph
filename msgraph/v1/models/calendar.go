package models

// Calendar struct for Calendar
type Calendar struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
	AllowedOnlineMeetingProviders *[]OnlineMeetingProviderType `json:"allowedOnlineMeetingProviders,omitempty"`
	// True if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
	CanEdit *bool `json:"canEdit,omitempty"`
	// True if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
	CanShare *bool `json:"canShare,omitempty"`
	// True if the user can read calendar items that have been marked private, false otherwise.
	CanViewPrivateItems *bool `json:"canViewPrivateItems,omitempty"`
	// Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey *string `json:"changeKey,omitempty"`
	// Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: LightBlue=0, LightGreen=1, LightOrange=2, LightGray=3, LightYellow=4, LightTeal=5, LightPink=6, LightBrown=7, LightRed=8, MaxColor=9, Auto=-1
	Color *CalendarColor `json:"color,omitempty"`
	// The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
	DefaultOnlineMeetingProvider *OnlineMeetingProviderType `json:"defaultOnlineMeetingProvider,omitempty"`
	// The calendar color, expressed in a hex color code of three hexidecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
	HexColor *string `json:"hexColor,omitempty"`
	// True if this is the default calendar where new events are created by default, false otherwise.
	IsDefaultCalendar *bool `json:"isDefaultCalendar,omitempty"`
	// Indicates whether this user calendar can be deleted from the user mailbox.
	IsRemovable *bool `json:"isRemovable,omitempty"`
	// Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
	IsTallyingResponses *bool `json:"isTallyingResponses,omitempty"`
	// The calendar name.
	Name *string `json:"name,omitempty"`
	// If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
	Owner *EmailAddress `json:"owner,omitempty"`
	// The permissions of the users with whom the calendar is shared.
	CalendarPermissions *[]CalendarPermission `json:"calendarPermissions,omitempty"`
	// The calendar view for the calendar. Navigation property. Read-only.
	CalendarView *[]Event `json:"calendarView,omitempty"`
	// The events in the calendar. Navigation property. Read-only.
	Events *[]Event `json:"events,omitempty"`
	// The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
