package models

import (
	"encoding/json"
	"fmt"
)

type UserAnalytics struct {
	// The current settings for a user to use the analytics API.
	Settings *Settings `json:"settings,omitempty"`
	// The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
	ActivityStatistics *[]ActivityStatistics `json:"activityStatistics,omitempty"`
}

type Settings struct {
	// Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
	HasGraphMailbox *bool `json:"hasGraphMailbox,omitempty"`
	// Specifies if the user has a MyAnalytics license assigned.
	HasLicense *bool `json:"hasLicense,omitempty"`
	// Specifies if the user opted out of MyAnalytics.
	HasOptedOut *bool `json:"hasOptedOut,omitempty"`
}

type ActivityStatistics struct {
	// The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
	Activity AnalyticsActivityType `json:"activity,omitempty"`
	// Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
	Duration *string `json:"duration,omitempty"`
	// Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
	EndDate *string `json:"endDate,omitempty"`
	// Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
	StartDate *string `json:"startDate,omitempty"`
	// The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
	TimeZoneUsed *string `json:"timeZoneUsed,omitempty"`
}

type AnalyticsActivityType string

// List of microsoft.graph.analyticsActivityType
const (
	EMAIL   AnalyticsActivityType = "Email"
	MEETING AnalyticsActivityType = "Meeting"
	FOCUS   AnalyticsActivityType = "Focus"
	CHAT    AnalyticsActivityType = "Chat"
	CALL    AnalyticsActivityType = "Call"
)

func (v *AnalyticsActivityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnalyticsActivityType(value)
	for _, existing := range []AnalyticsActivityType{"Email", "Meeting", "Focus", "Chat", "Call"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnalyticsActivityType", value)
}
