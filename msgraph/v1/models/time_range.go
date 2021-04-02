package models

// TimeRange struct for TimeRange
type TimeRange struct {
	// End time for the time range.
	EndTime *string `json:"endTime,omitempty"`
	// Start time for the time range.
	StartTime *string `json:"startTime,omitempty"`
}
