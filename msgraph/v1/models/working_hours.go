package models

// WorkingHours struct for WorkingHours
type WorkingHours struct {
	// The days of the week on which the user works.
	DaysOfWeek *[]DayOfWeek `json:"daysOfWeek,omitempty"`
	// The time of the day that the user stops working.
	EndTime *string `json:"endTime,omitempty"`
	// The time of the day that the user starts working.
	StartTime *string `json:"startTime,omitempty"`
	// The time zone to which the working hours apply.
	TimeZone *TimeZoneBase `json:"timeZone,omitempty"`
}
