package models

import (
	"encoding/json"
	"fmt"
)

// RecurrenceRangeType the model 'RecurrenceRangeType'
type RecurrenceRangeType string

// List of microsoft.graph.recurrenceRangeType
const (
	END_DATE RecurrenceRangeType = "endDate"
	NO_END   RecurrenceRangeType = "noEnd"
	NUMBERED RecurrenceRangeType = "numbered"
)

func (v *RecurrenceRangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecurrenceRangeType(value)
	for _, existing := range []RecurrenceRangeType{"endDate", "noEnd", "numbered"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecurrenceRangeType", value)
}
