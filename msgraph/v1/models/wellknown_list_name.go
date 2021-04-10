package models

import (
	"encoding/json"
	"fmt"
)

type WellknownListName string

const (
	WELLKNOWNLISTNAME_NONE                 WellknownListName = "none"
	WELLKNOWNLISTNAME_DEFAULT_LIST         WellknownListName = "defaultList"
	WELLKNOWNLISTNAME_FLAGGED_EMAILS       WellknownListName = "flaggedEmails"
	WELLKNOWNLISTNAME_UNKNOWN_FUTURE_VALUE WellknownListName = "unknownFutureValue"
)

func (v *WellknownListName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WellknownListName(value)
	for _, existing := range []WellknownListName{"none", "defaultList", "flaggedEmails", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WellknownListName", value)
}
