package models

import (
	"encoding/json"
	"fmt"
)

type WeekIndex string

const (
	WeekIndex_FIRST  WeekIndex = "first"
	WeekIndex_SECOND WeekIndex = "second"
	WeekIndex_THIRD  WeekIndex = "third"
	WeekIndex_FOURTH WeekIndex = "fourth"
	WeekIndex_LAST   WeekIndex = "last"
)

func (v *WeekIndex) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WeekIndex(value)
	for _, existing := range []WeekIndex{"first", "second", "third", "fourth", "last"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WeekIndex", value)
}
