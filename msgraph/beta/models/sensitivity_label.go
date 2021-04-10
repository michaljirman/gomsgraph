package models

import (
	"encoding/json"
	"fmt"
)

type SensitivityLabel struct {
	ApplicableTo                *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	ApplicationMode             *ApplicationMode        `json:"applicationMode,omitempty"`
	AssignedPolicies            *[]LabelPolicy          `json:"assignedPolicies,omitempty"`
	AutoLabeling                *AutoLabeling           `json:"autoLabeling,omitempty"`
	Description                 *string                 `json:"description,omitempty"`
	DisplayName                 *string                 `json:"displayName,omitempty"`
	IsDefault                   *bool                   `json:"isDefault,omitempty"`
	IsEndpointProtectionEnabled *bool                   `json:"isEndpointProtectionEnabled,omitempty"`
	LabelActions                *[]LabelActionBase      `json:"labelActions,omitempty"`
	Name                        *string                 `json:"name,omitempty"`
	Priority                    *int                    `json:"priority,omitempty"`
	ToolTip                     *string                 `json:"toolTip,omitempty"`
	Sublabels                   *[]SensitivityLabel     `json:"sublabels,omitempty"`
}

type LabelActionBase struct {
	Name *string `json:"name,omitempty"`
}

type AutoLabeling struct {
	Message          *string   `json:"message,omitempty"`
	SensitiveTypeIds *[]string `json:"sensitiveTypeIds,omitempty"`
}

type LabelPolicy struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ApplicationMode string

const (
	APPLICATION_MODE_MANUAL      ApplicationMode = "manual"
	APPLICATION_MODE_AUTOMATIC   ApplicationMode = "automatic"
	APPLICATION_MODE_RECOMMENDED ApplicationMode = "recommended"
)

func (v *ApplicationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationMode(value)
	for _, existing := range []ApplicationMode{"manual", "automatic", "recommended"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationMode", value)
}
