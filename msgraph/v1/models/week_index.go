package models

import (
	"encoding/json"
	"fmt"
)

// WeekIndex the model 'WeekIndex'
type WeekIndex string

// List of microsoft.graph.weekIndex
const (
	FIRST  WeekIndex = "first"
	SECOND WeekIndex = "second"
	THIRD  WeekIndex = "third"
	FOURTH WeekIndex = "fourth"
	LAST   WeekIndex = "last"
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
