package models

import (
	"encoding/json"
	"fmt"
)

type CalendarRoleType string

const (
	CalendarRoleType_NONE                                  CalendarRoleType = "none"
	CalendarRoleType_FREE_BUSY_READ                        CalendarRoleType = "freeBusyRead"
	CalendarRoleType_LIMITED_READ                          CalendarRoleType = "limitedRead"
	CalendarRoleType_READ                                  CalendarRoleType = "read"
	CalendarRoleType_WRITE                                 CalendarRoleType = "write"
	CalendarRoleType_DELEGATE_WITHOUT_PRIVATE_EVENT_ACCESS CalendarRoleType = "delegateWithoutPrivateEventAccess"
	CalendarRoleType_DELEGATE_WITH_PRIVATE_EVENT_ACCESS    CalendarRoleType = "delegateWithPrivateEventAccess"
	CalendarRoleType_CUSTOM                                CalendarRoleType = "custom"
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
