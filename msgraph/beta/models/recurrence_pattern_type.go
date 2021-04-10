package models

import (
	"encoding/json"
	"fmt"
)

type RecurrencePatternType string

const (
	RecurrencePatternType_DAILY            RecurrencePatternType = "daily"
	RecurrencePatternType_WEEKLY           RecurrencePatternType = "weekly"
	RecurrencePatternType_ABSOLUTE_MONTHLY RecurrencePatternType = "absoluteMonthly"
	RecurrencePatternType_RELATIVE_MONTHLY RecurrencePatternType = "relativeMonthly"
	RecurrencePatternType_ABSOLUTE_YEARLY  RecurrencePatternType = "absoluteYearly"
	RecurrencePatternType_RELATIVE_YEARLY  RecurrencePatternType = "relativeYearly"
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
