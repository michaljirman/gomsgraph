package models

import (
	"encoding/json"
	"fmt"
)

type RecurrenceRangeType string

const (
	RecurrenceRangeType_END_DATE RecurrenceRangeType = "endDate"
	RecurrenceRangeType_NO_END   RecurrenceRangeType = "noEnd"
	RecurrenceRangeType_NUMBERED RecurrenceRangeType = "numbered"
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
