package models

import (
	"encoding/json"
	"fmt"
)

// DayOfWeek the model 'DayOfWeek'
type DayOfWeek string

// List of microsoft.graph.dayOfWeek
const (
	SUNDAY    DayOfWeek = "sunday"
	MONDAY    DayOfWeek = "monday"
	TUESDAY   DayOfWeek = "tuesday"
	WEDNESDAY DayOfWeek = "wednesday"
	THURSDAY  DayOfWeek = "thursday"
	FRIDAY    DayOfWeek = "friday"
	SATURDAY  DayOfWeek = "saturday"
)

func (v *DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayOfWeek(value)
	for _, existing := range []DayOfWeek{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayOfWeek", value)
}
