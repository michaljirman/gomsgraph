package models

import (
	"encoding/json"
	"fmt"
)

// ManagedAppFlaggedReason the model 'ManagedAppFlaggedReason'
type ManagedAppFlaggedReason string

// List of microsoft.graph.managedAppFlaggedReason
const (
	MANAGEDAPPFLAGGEDREASON_NONE          ManagedAppFlaggedReason = "none"
	MANAGEDAPPFLAGGEDREASON_ROOTED_DEVICE ManagedAppFlaggedReason = "rootedDevice"
)

func (v *ManagedAppFlaggedReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManagedAppFlaggedReason(value)
	for _, existing := range []ManagedAppFlaggedReason{"none", "rootedDevice"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManagedAppFlaggedReason", value)
}
