package models

import (
	"encoding/json"
	"fmt"
)

type SensitivityLabelTarget string

// List of microsoft.graph.sensitivityLabelTarget
const (
	SensitivityLabelTarget_EMAIL                SensitivityLabelTarget = "email"
	SensitivityLabelTarget_SITE                 SensitivityLabelTarget = "site"
	SensitivityLabelTarget_UNIFIED_GROUP        SensitivityLabelTarget = "unifiedGroup"
	SensitivityLabelTarget_UNKNOWN_FUTURE_VALUE SensitivityLabelTarget = "unknownFutureValue"
)

func (v *SensitivityLabelTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SensitivityLabelTarget(value)
	for _, existing := range []SensitivityLabelTarget{"email", "site", "unifiedGroup", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SensitivityLabelTarget", value)
}
