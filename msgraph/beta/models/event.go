package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Id *string `json:"id,omitempty"`
	// The categories associated with the item
	Categories *[]string `json:"categories,omitempty"`
	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey *string `json:"changeKey,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// True if the meeting organizer allows invitees to propose a new time when responding, false otherwise. Optional. Default is true.
	AllowNewTimeProposals *bool `json:"allowNewTimeProposals,omitempty"`
	// The collection of attendees for the event.
	Attendees *[]Attendee `json:"attendees,omitempty"`
	// The body of the message associated with the event. It can be in HTML or text format.
	Body *ItemBody `json:"body,omitempty"`
	// The preview of the message associated with the event. It is in text format.
	BodyPreview *string `json:"bodyPreview,omitempty"`
	// The date, time, and time zone that the event ends. By default, the end time is in UTC.
	End *DateTimeTimeZone `json:"end,omitempty"`
	// Set to true if the event has attachments.
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
	HideAttendees *bool `json:"hideAttendees,omitempty"`
	// A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
	ICalUId *string `json:"iCalUId,omitempty"`
	// The importance of the event. The possible values are: low, normal, high.
	Importance *Importance `json:"importance,omitempty"`
	// Set to true if the event lasts all day.
	IsAllDay *bool `json:"isAllDay,omitempty"`
	// Set to true if the event has been canceled.
	IsCancelled *bool `json:"isCancelled,omitempty"`
	// Set to true if the user has updated the meeting in Outlook but has not sent the updates to attendees. Set to false if all changes have been sent, or if the event is an appointment without any attendees.
	IsDraft *bool `json:"isDraft,omitempty"`
	// True if this event has online meeting information, false otherwise. Default is false. Optional.
	IsOnlineMeeting *bool `json:"isOnlineMeeting,omitempty"`
	// Set to true if the calendar owner (specified by the owner property of the calendar) is the organizer of the event (specified by the organizer property of the event). This also applies if a delegate organized the event on behalf of the owner.
	IsOrganizer *bool `json:"isOrganizer,omitempty"`
	// Set to true if an alert is set to remind the user of the event.
	IsReminderOn *bool `json:"isReminderOn,omitempty"`
	// The location of the event.
	Location *Location `json:"location,omitempty"`
	// The locations where the event is held or attended from. The location and locations properties always correspond with each other. If you update the location property, any prior locations in the locations collection would be removed and replaced by the new location value.
	Locations *[]Location `json:"locations,omitempty"`
	// Details for an attendee to join the meeting online. Read-only.
	OnlineMeeting *OnlineMeetingInfo `json:"onlineMeeting,omitempty"`
	// Represents the online meeting service provider. The possible values are teamsForBusiness, skypeForBusiness, and skypeForConsumer. Optional.
	OnlineMeetingProvider *OnlineMeetingProviderType `json:"onlineMeetingProvider,omitempty"`
	// A URL for an online meeting. The property is set only when an organizer specifies an event as an online meeting such as a Skype meeting. Read-only.
	OnlineMeetingUrl *string `json:"onlineMeetingUrl,omitempty"`
	// The organizer of the event.
	Organizer *Recipient `json:"organizer,omitempty"`
	// The end time zone that was set when the event was created. A value of tzone://Microsoft/Custom indicates that a legacy custom time zone was set in desktop Outlook.
	OriginalEndTimeZone *string `json:"originalEndTimeZone,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	OriginalStart *time.Time `json:"originalStart,omitempty"`
	// The start time zone that was set when the event was created. A value of tzone://Microsoft/Custom indicates that a legacy custom time zone was set in desktop Outlook.
	OriginalStartTimeZone *string `json:"originalStartTimeZone,omitempty"`
	// The recurrence pattern for the event.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// The number of minutes before the event start time that the reminder alert occurs.
	ReminderMinutesBeforeStart *int `json:"reminderMinutesBeforeStart,omitempty"`
	// Default is true, which represents the organizer would like an invitee to send a response to the event.
	ResponseRequested *bool `json:"responseRequested,omitempty"`
	// Indicates the type of response sent in response to an event message.
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
	// The possible values are: normal, personal, private, confidential.
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// The ID for the recurring series master item, if this event is part of a recurring series.
	SeriesMasterId *string `json:"seriesMasterId,omitempty"`
	// The status to show. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
	ShowAs *FreeBusyStatus `json:"showAs,omitempty"`
	// The date, time, and time zone that the event starts. By default, the start time is in UTC.
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// The text of the event's subject line.
	Subject *string `json:"subject,omitempty"`
	// A custom identifier specified by a client app for the server to avoid redundant POST operations in case of client retries to create the same event. This is useful when low network connectivity causes the client to time out before receiving a response from the server for the client's prior create-event request. After you set transactionId when creating an event, you cannot change transactionId in a subsequent update. This property is only returned in a response payload if an app has set it. Optional.
	TransactionId *string `json:"transactionId,omitempty"`
	// The event type. The possible values are: singleInstance, occurrence, exception, seriesMaster. Read-only.
	Type *EventType `json:"type,omitempty"`
	// The URL to open the event in Outlook on the web.Outlook on the web opens the event in the browser if you are signed in to your mailbox. Otherwise, Outlook on the web prompts you to sign in.This URL cannot be accessed from within an iFrame.
	WebLink *string `json:"webLink,omitempty"`
	// The collection of fileAttachment and itemAttachment attachments for the event. Navigation property. Read-only. Nullable.
	Attachments *[]Attachment `json:"attachments,omitempty"`
	// The calendar that contains the event. Navigation property. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`
	// The collection of open extensions defined for the event. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// The instances of the event. Navigation property. Read-only. Nullable.
	Instances *[]Event `json:"instances,omitempty"`
	// The collection of multi-value extended properties defined for the event. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// The collection of single-value extended properties defined for the event. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

type Attendee struct {
	// The recipient's email address.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
	// The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
	Type *AttendeeType `json:"type,omitempty"`
	// An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't proposed another time, then this property is not included in a response of a GET event.
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`
	// The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
	Status *ResponseStatus `json:"status,omitempty"`
}

type ItemBody struct {
	// The content of the item.
	Content *string `json:"content,omitempty"`
	// The type of the content. Possible values are text and html.
	ContentType *BodyType `json:"contentType,omitempty"`
}

type Location struct {
	// The street address of the location.
	Address *PhysicalAddress `json:"address,omitempty"`
	// The geographic coordinates and elevation of the location.
	Coordinates *OutlookGeoCoordinates `json:"coordinates,omitempty"`
	// The name associated with the location.
	DisplayName *string `json:"displayName,omitempty"`
	// Optional email address of the location.
	LocationEmailAddress *string `json:"locationEmailAddress,omitempty"`
	// The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
	LocationType *LocationType `json:"locationType,omitempty"`
	// Optional URI representing the location.
	LocationUri *string `json:"locationUri,omitempty"`
	// For internal use only.
	UniqueId *string `json:"uniqueId,omitempty"`
	// For internal use only.
	UniqueIdType *LocationUniqueIdType `json:"uniqueIdType,omitempty"`
}

type OnlineMeetingInfo struct {
	// The ID of the conference.
	ConferenceId *string `json:"conferenceId,omitempty"`
	// The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
	JoinUrl *string `json:"joinUrl,omitempty"`
	// All of the phone numbers associated with this conference.
	Phones *[]Phone `json:"phones,omitempty"`
	// The pre-formatted quickdial for this call.
	QuickDial *string `json:"quickDial,omitempty"`
	// The toll free numbers that can be used to join the conference.
	TollFreeNumbers *[]string `json:"tollFreeNumbers,omitempty"`
	// The toll number that can be used to join the conference.
	TollNumber *string `json:"tollNumber,omitempty"`
}

type ResponseStatus struct {
	// The response type. The possible values are: None, Organizer, TentativelyAccepted, Accepted, Declined, NotResponded.
	Response *ResponseType `json:"response,omitempty"`
	// The date and time that the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	Time *time.Time `json:"time,omitempty"`
}

type FreeBusyStatus string

const (
	FreeBusyStatus_UNKNOWN           FreeBusyStatus = "unknown"
	FreeBusyStatus_FREE              FreeBusyStatus = "free"
	FreeBusyStatus_TENTATIVE         FreeBusyStatus = "tentative"
	FreeBusyStatus_BUSY              FreeBusyStatus = "busy"
	FreeBusyStatus_OOF               FreeBusyStatus = "oof"
	FreeBusyStatus_WORKING_ELSEWHERE FreeBusyStatus = "workingElsewhere"
)

func (v *FreeBusyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FreeBusyStatus(value)
	for _, existing := range []FreeBusyStatus{"unknown", "free", "tentative", "busy", "oof", "workingElsewhere"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FreeBusyStatus", value)
}

// EventType the model 'EventType'
type EventType string

// List of microsoft.graph.eventType
const (
	EventType_SINGLE_INSTANCE EventType = "singleInstance"
	EventType_OCCURRENCE      EventType = "occurrence"
	EventType_EXCEPTION       EventType = "exception"
	EventType_SERIES_MASTER   EventType = "seriesMaster"
)

func (v *EventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventType(value)
	for _, existing := range []EventType{"singleInstance", "occurrence", "exception", "seriesMaster"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventType", value)
}

type Attachment struct {
	Id *string `json:"id,omitempty"`
	// The MIME type.
	ContentType *string `json:"contentType,omitempty"`
	// true if the attachment is an inline attachment; otherwise, false.
	IsInline *bool `json:"isInline,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The attachment's file name.
	Name *string `json:"name,omitempty"`
	// The length of the attachment in bytes.
	Size *int32 `json:"size,omitempty"`
}

type AttendeeType string

const (
	AttendeeType_REQUIRED AttendeeType = "required"
	AttendeeType_OPTIONAL AttendeeType = "optional"
	AttendeeType_RESOURCE AttendeeType = "resource"
)

func (v *AttendeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttendeeType(value)
	for _, existing := range []AttendeeType{"required", "optional", "resource"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttendeeType", value)
}

// TimeSlot struct for TimeSlot
type TimeSlot struct {
	End   *DateTimeTimeZone `json:"end,omitempty"`
	Start *DateTimeTimeZone `json:"start,omitempty"`
}

type BodyType string

const (
	BodyType_TEXT BodyType = "text"
	BodyType_HTML BodyType = "html"
)

func (v *BodyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BodyType(value)
	for _, existing := range []BodyType{"text", "html"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BodyType", value)
}

type OutlookGeoCoordinates struct {
	// The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
	Accuracy *float32 `json:"accuracy,omitempty"`
	// The altitude of the location.
	Altitude *float32 `json:"altitude,omitempty"`
	// The accuracy of the altitude.
	AltitudeAccuracy *float32 `json:"altitudeAccuracy,omitempty"`
	// The latitude of the location.
	Latitude *float32 `json:"latitude,omitempty"`
	// The longitude of the location.
	Longitude *float32 `json:"longitude,omitempty"`
}

// ResponseType the model 'ResponseType'
type ResponseType string

// List of microsoft.graph.responseType
const (
	ResponseType_NONE                 ResponseType = "none"
	ResponseType_ORGANIZER            ResponseType = "organizer"
	ResponseType_TENTATIVELY_ACCEPTED ResponseType = "tentativelyAccepted"
	ResponseType_ACCEPTED             ResponseType = "accepted"
	ResponseType_DECLINED             ResponseType = "declined"
	ResponseType_NOT_RESPONDED        ResponseType = "notResponded"
)

func (v *ResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseType(value)
	for _, existing := range []ResponseType{"none", "organizer", "tentativelyAccepted", "accepted", "declined", "notResponded"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponseType", value)
}

type LocationType string

const (
	LocationType_DEFAULT          LocationType = "default"
	LocationType_CONFERENCE_ROOM  LocationType = "conferenceRoom"
	LocationType_HOME_ADDRESS     LocationType = "homeAddress"
	LocationType_BUSINESS_ADDRESS LocationType = "businessAddress"
	LocationType_GEO_COORDINATES  LocationType = "geoCoordinates"
	LocationType_STREET_ADDRESS   LocationType = "streetAddress"
	LocationType_HOTEL            LocationType = "hotel"
	LocationType_RESTAURANT       LocationType = "restaurant"
	LocationType_LOCAL_BUSINESS   LocationType = "localBusiness"
	LocationType_POSTAL_ADDRESS   LocationType = "postalAddress"
)

func (v *LocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationType(value)
	for _, existing := range []LocationType{"default", "conferenceRoom", "homeAddress", "businessAddress", "geoCoordinates", "streetAddress", "hotel", "restaurant", "localBusiness", "postalAddress"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationType", value)
}

type LocationUniqueIdType string

const (
	LocationUniqueIdType_UNKNOWN        LocationUniqueIdType = "unknown"
	LocationUniqueIdType_LOCATION_STORE LocationUniqueIdType = "locationStore"
	LocationUniqueIdType_DIRECTORY      LocationUniqueIdType = "directory"
	LocationUniqueIdType_PRIVATE        LocationUniqueIdType = "private"
	LocationUniqueIdType_BING           LocationUniqueIdType = "bing"
)

func (v *LocationUniqueIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationUniqueIdType(value)
	for _, existing := range []LocationUniqueIdType{"unknown", "locationStore", "directory", "private", "bing"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationUniqueIdType", value)
}

type PhoneType string

const (
	PhoneType_HOME         PhoneType = "home"
	PhoneType_BUSINESS     PhoneType = "business"
	PhoneType_MOBILE       PhoneType = "mobile"
	PhoneType_OTHER        PhoneType = "other"
	PhoneType_ASSISTANT    PhoneType = "assistant"
	PhoneType_HOME_FAX     PhoneType = "homeFax"
	PhoneType_BUSINESS_FAX PhoneType = "businessFax"
	PhoneType_OTHER_FAX    PhoneType = "otherFax"
	PhoneType_PAGER        PhoneType = "pager"
	PhoneType_RADIO        PhoneType = "radio"
)

func (v *PhoneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhoneType(value)
	for _, existing := range []PhoneType{"home", "business", "mobile", "other", "assistant", "homeFax", "businessFax", "otherFax", "pager", "radio"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhoneType", value)
}
