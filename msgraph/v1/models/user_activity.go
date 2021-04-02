package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// UserActivity struct for UserActivity
type UserActivity struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
	ActivationUrl *string `json:"activationUrl,omitempty"`
	// Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
	ActivitySourceHost *string `json:"activitySourceHost,omitempty"`
	// Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
	AppActivityId *string `json:"appActivityId,omitempty"`
	// Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
	ContentInfo *interface{} `json:"contentInfo,omitempty"`
	// Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
	ContentUrl *string `json:"contentUrl,omitempty"`
	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Set by the server. DateTime in UTC when the object expired on the server.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Optional. URL used to launch the activity in a web-based app, if available.
	FallbackUrl *string `json:"fallbackUrl,omitempty"`
	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status *Status `json:"status,omitempty"`
	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone   *string     `json:"userTimezone,omitempty"`
	VisualElements *VisualInfo `json:"visualElements,omitempty"`
	// Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
	HistoryItems *[]ActivityHistoryItem `json:"historyItems,omitempty"`
}

// Status the model 'Status'
type Status string

// List of microsoft.graph.status
const (
	Status_ACTIVE               Status = "active"
	Status_UPDATED              Status = "updated"
	Status_DELETED              Status = "deleted"
	Status_IGNORED              Status = "ignored"
	Status_UNKNOWN_FUTURE_VALUE Status = "unknownFutureValue"
)

func (v *Status) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status(value)
	for _, existing := range []Status{"active", "updated", "deleted", "ignored", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status", value)
}

// VisualInfo struct for VisualInfo
type VisualInfo struct {
	// Optional. JSON object used to represent an icon which represents the application used to generate the activity
	Attribution *ImageInfo `json:"attribution,omitempty"`
	// Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
	BackgroundColor *string `json:"backgroundColor,omitempty"`
	// Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
	Content *interface{} `json:"content,omitempty"`
	// Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
	Description *string `json:"description,omitempty"`
	// Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
	DisplayText *string `json:"displayText,omitempty"`
}

// ActivityHistoryItem struct for ActivityHistoryItem
type ActivityHistoryItem struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Optional. The duration of active user engagement. if not supplied, this is calculated from the startedDateTime and lastActiveDateTime.
	ActiveDurationSeconds *int `json:"activeDurationSeconds,omitempty"`
	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Optional. UTC DateTime when the historyItem will undergo hard-delete. Can be set by the client.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Optional. UTC DateTime when the historyItem (activity session) was last understood as active or finished - if null, historyItem status should be Ongoing.
	LastActiveDateTime *time.Time `json:"lastActiveDateTime,omitempty"`
	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Required. UTC DateTime when the historyItem (activity session) was started. Required for timeline history.
	StartedDateTime *time.Time `json:"startedDateTime,omitempty"`
	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status *Status `json:"status,omitempty"`
	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation time. Values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone *string       `json:"userTimezone,omitempty"`
	Activity     *UserActivity `json:"activity,omitempty"`
}

// ImageInfo struct for ImageInfo
type ImageInfo struct {
	// Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
	AddImageQuery *bool `json:"addImageQuery,omitempty"`
	// Optional; alt-text accessible content for the image
	AlternateText   *string `json:"alternateText,omitempty"`
	AlternativeText *string `json:"alternativeText,omitempty"`
	// Optional; URI that points to an icon which represents the application used to generate the activity
	IconUrl *string `json:"iconUrl,omitempty"`
}
