package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type InformationProtection struct {
	Bitlocker                  *Bitlocker                   `json:"bitlocker,omitempty"`
	DataLossPreventionPolicies *[]DataLossPreventionPolicy  `json:"dataLossPreventionPolicies,omitempty"`
	SensitivityLabels          *[]SensitivityLabel          `json:"sensitivityLabels,omitempty"`
	SensitivityPolicySettings  *SensitivityPolicySettings   `json:"sensitivityPolicySettings,omitempty"`
	Policy                     *InformationProtectionPolicy `json:"policy,omitempty"`
	ThreatAssessmentRequests   *[]ThreatAssessmentRequest   `json:"threatAssessmentRequests,omitempty"`
}

type DataLossPreventionPolicy struct {
	Name *string `json:"name,omitempty"`
}

type Bitlocker struct {
	// The recovery keys associated with the bitlocker entity.
	RecoveryKeys *[]BitlockerRecoveryKey `json:"recoveryKeys,omitempty"`
}

type BitlockerRecoveryKey struct {
	// The date and time when the key was originally backed up to Azure Active Directory.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ID of the device the BitLocker key is originally backed up from.
	DeviceId *string `json:"deviceId,omitempty"`
	// The BitLocker recovery key.
	Key *string `json:"key,omitempty"`
	// Indicates the type of volume the BitLocker key is associated with. Possible values are: operatingSystemVolume, fixedDataVolume, removableDataVolume, unknownFutureValue.
	VolumeType *VolumeType `json:"volumeType,omitempty"`
}

type VolumeType string

// List of microsoft.graph.volumeType
const (
	VOLUMETYPE_OPERATING_SYSTEM_VOLUME VolumeType = "operatingSystemVolume"
	VOLUMETYPE_FIXED_DATA_VOLUME       VolumeType = "fixedDataVolume"
	VOLUMETYPE_REMOVABLE_DATA_VOLUME   VolumeType = "removableDataVolume"
	VOLUMETYPE_UNKNOWN_FUTURE_VALUE    VolumeType = "unknownFutureValue"
)

func (v *VolumeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VolumeType(value)
	for _, existing := range []VolumeType{"operatingSystemVolume", "fixedDataVolume", "removableDataVolume", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VolumeType", value)
}

type SensitivityPolicySettings struct {
	ApplicableTo                              *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	DowngradeSensitivityRequiresJustification *bool                   `json:"downgradeSensitivityRequiresJustification,omitempty"`
	HelpWebUrl                                *string                 `json:"helpWebUrl,omitempty"`
	IsMandatory                               *bool                   `json:"isMandatory,omitempty"`
}
