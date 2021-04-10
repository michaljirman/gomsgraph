package models

type DateTimeTimeZone struct {
	// A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).
	DateTime *string `json:"dateTime,omitempty"`
	// Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
	TimeZone *string `json:"timeZone,omitempty"`
}
