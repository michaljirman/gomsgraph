package models

import (
	"encoding/json"
	"fmt"
)

// RecurrencePatternType the model 'RecurrencePatternType'
type RecurrencePatternType string

// List of microsoft.graph.recurrencePatternType
const (
	RECURRENCEPATTERNTYPE_DAILY            RecurrencePatternType = "daily"
	RECURRENCEPATTERNTYPE_WEEKLY           RecurrencePatternType = "weekly"
	RECURRENCEPATTERNTYPE_ABSOLUTE_MONTHLY RecurrencePatternType = "absoluteMonthly"
	RECURRENCEPATTERNTYPE_RELATIVE_MONTHLY RecurrencePatternType = "relativeMonthly"
	RECURRENCEPATTERNTYPE_ABSOLUTE_YEARLY  RecurrencePatternType = "absoluteYearly"
	RECURRENCEPATTERNTYPE_RELATIVE_YEARLY  RecurrencePatternType = "relativeYearly"
)

func (v *RecurrencePatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecurrencePatternType(value)
	for _, existing := range []RecurrencePatternType{"daily", "weekly", "absoluteMonthly", "relativeMonthly", "absoluteYearly", "relativeYearly"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecurrencePatternType", value)
}
