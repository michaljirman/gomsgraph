package models

type PatternedRecurrence struct {
	// The frequency of an event.
	Pattern *RecurrencePattern `json:"pattern,omitempty"`
	// The duration of an event.
	Range *RecurrenceRange `json:"range,omitempty"`
}

type RecurrencePattern struct {
	// The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
	DayOfMonth *int32 `json:"dayOfMonth,omitempty"`
	// A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
	DaysOfWeek *[]DayOfWeek `json:"daysOfWeek,omitempty"`
	// The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
	FirstDayOfWeek *DayOfWeek `json:"firstDayOfWeek,omitempty"`
	// Specifies on which instance of the allowed days specified in daysOfsWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
	Index *WeekIndex `json:"index,omitempty"`
	// The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
	Interval *int32 `json:"interval,omitempty"`
	// The month in which the event occurs.  This is a number from 1 to 12.
	Month *int32 `json:"month,omitempty"`
	// The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required.
	Type *RecurrencePatternType `json:"type,omitempty"`
}
