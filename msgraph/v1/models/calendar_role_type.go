package models

import (
	"encoding/json"
	"fmt"
)

// CalendarRoleType the model 'CalendarRoleType'
type CalendarRoleType string

// List of microsoft.graph.calendarRoleType
const (
	CALENDARROLETYPE_NONE                                  CalendarRoleType = "none"
	CALENDARROLETYPE_FREE_BUSY_READ                        CalendarRoleType = "freeBusyRead"
	CALENDARROLETYPE_LIMITED_READ                          CalendarRoleType = "limitedRead"
	CALENDARROLETYPE_READ                                  CalendarRoleType = "read"
	CALENDARROLETYPE_WRITE                                 CalendarRoleType = "write"
	CALENDARROLETYPE_DELEGATE_WITHOUT_PRIVATE_EVENT_ACCESS CalendarRoleType = "delegateWithoutPrivateEventAccess"
	CALENDARROLETYPE_DELEGATE_WITH_PRIVATE_EVENT_ACCESS    CalendarRoleType = "delegateWithPrivateEventAccess"
	CALENDARROLETYPE_CUSTOM                                CalendarRoleType = "custom"
)

func (v *CalendarRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CalendarRoleType(value)
	for _, existing := range []CalendarRoleType{"none", "freeBusyRead", "limitedRead", "read", "write", "delegateWithoutPrivateEventAccess", "delegateWithPrivateEventAccess", "custom"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CalendarRoleType", value)
}
