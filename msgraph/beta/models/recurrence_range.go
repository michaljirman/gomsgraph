package models

type RecurrenceRange struct {
	// The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
	EndDate *string `json:"endDate,omitempty"`
	// The number of times to repeat the event. Required and must be positive if type is numbered.
	NumberOfOccurrences *int32 `json:"numberOfOccurrences,omitempty"`
	// Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
	RecurrenceTimeZone *string `json:"recurrenceTimeZone,omitempty"`
	// The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
	StartDate *string `json:"startDate,omitempty"`
	// The recurrence range. The possible values are: endDate, noEnd, numbered. Required.
	Type *RecurrenceRangeType `json:"type,omitempty"`
}
