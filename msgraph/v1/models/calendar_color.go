package models

import (
	"encoding/json"
	"fmt"
)

// CalendarColor the model 'CalendarColor'
type CalendarColor string

// List of microsoft.graph.calendarColor
const (
	AUTO         CalendarColor = "auto"
	LIGHT_BLUE   CalendarColor = "lightBlue"
	LIGHT_GREEN  CalendarColor = "lightGreen"
	LIGHT_ORANGE CalendarColor = "lightOrange"
	LIGHT_GRAY   CalendarColor = "lightGray"
	LIGHT_YELLOW CalendarColor = "lightYellow"
	LIGHT_TEAL   CalendarColor = "lightTeal"
	LIGHT_PINK   CalendarColor = "lightPink"
	LIGHT_BROWN  CalendarColor = "lightBrown"
	LIGHT_RED    CalendarColor = "lightRed"
	MAX_COLOR    CalendarColor = "maxColor"
)

func (v *CalendarColor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CalendarColor(value)
	for _, existing := range []CalendarColor{"auto", "lightBlue", "lightGreen", "lightOrange", "lightGray", "lightYellow", "lightTeal", "lightPink", "lightBrown", "lightRed", "maxColor"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CalendarColor", value)
}
