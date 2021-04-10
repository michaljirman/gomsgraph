package models

import (
	"encoding/json"
	"fmt"
)

type UsageRight struct {
	// Product id corresponding to the usage right.
	CatalogId *string `json:"catalogId,omitempty"`
	// Identifier of the service corresponding to the usage right.
	ServiceIdentifier *string `json:"serviceIdentifier,omitempty"`
	// The state of the usage right. Possible values are: active, inactive, warning, suspended.
	State *UsageRightState `json:"state,omitempty"`
}

type UsageRightState string

const (
	ACTIVE               UsageRightState = "active"
	INACTIVE             UsageRightState = "inactive"
	WARNING              UsageRightState = "warning"
	SUSPENDED            UsageRightState = "suspended"
	UNKNOWN_FUTURE_VALUE UsageRightState = "unknownFutureValue"
)

func (v *UsageRightState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageRightState(value)
	for _, existing := range []UsageRightState{"active", "inactive", "warning", "suspended", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageRightState", value)
}
