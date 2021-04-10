package models

import (
	"encoding/json"
	"fmt"
)

type CalendarColor string

const (
	CalendarColor_AUTO         CalendarColor = "auto"
	CalendarColor_LIGHT_BLUE   CalendarColor = "lightBlue"
	CalendarColor_LIGHT_GREEN  CalendarColor = "lightGreen"
	CalendarColor_LIGHT_ORANGE CalendarColor = "lightOrange"
	CalendarColor_LIGHT_GRAY   CalendarColor = "lightGray"
	CalendarColor_LIGHT_YELLOW CalendarColor = "lightYellow"
	CalendarColor_LIGHT_TEAL   CalendarColor = "lightTeal"
	CalendarColor_LIGHT_PINK   CalendarColor = "lightPink"
	CalendarColor_LIGHT_BROWN  CalendarColor = "lightBrown"
	CalendarColor_LIGHT_RED    CalendarColor = "lightRed"
	CalendarColor_MAX_COLOR    CalendarColor = "maxColor"
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
