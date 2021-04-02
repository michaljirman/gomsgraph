package models

import (
	"encoding/json"
	"fmt"
)

// AutomaticRepliesStatus the model 'AutomaticRepliesStatus'
type AutomaticRepliesStatus string

// List of microsoft.graph.automaticRepliesStatus
const (
	DISABLED       AutomaticRepliesStatus = "disabled"
	ALWAYS_ENABLED AutomaticRepliesStatus = "alwaysEnabled"
	SCHEDULED      AutomaticRepliesStatus = "scheduled"
)

func (v *AutomaticRepliesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutomaticRepliesStatus(value)
	for _, existing := range []AutomaticRepliesStatus{"disabled", "alwaysEnabled", "scheduled"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutomaticRepliesStatus", value)
}
