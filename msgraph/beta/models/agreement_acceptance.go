package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type AgreementAcceptance struct {
	Id *string `json:"id,omitempty"`
	// ID of the agreement file accepted by the user.
	AgreementFileId *string `json:"agreementFileId,omitempty"`
	// ID of the agreement.
	AgreementId *string `json:"agreementId,omitempty"`
	// The display name of the device used for accepting the agreement.
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	// The unique identifier of the device used for accepting the agreement.
	DeviceId *string `json:"deviceId,omitempty"`
	// The operating system used for accepting the agreement.
	DeviceOSType *string `json:"deviceOSType,omitempty"`
	// The operating system version of the device used for accepting the agreement.
	DeviceOSVersion *string `json:"deviceOSVersion,omitempty"`
	// The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	RecordedDateTime *time.Time `json:"recordedDateTime,omitempty"`
	// Possible values are: accepted, declined.
	State *AgreementAcceptanceState `json:"state,omitempty"`
	// Display name of the user when the acceptance was recorded.
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// Email of the user when the acceptance was recorded.
	UserEmail *string `json:"userEmail,omitempty"`
	// ID of the user who accepted the agreement.
	UserId *string `json:"userId,omitempty"`
	// UPN of the user when the acceptance was recorded.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

type AgreementAcceptanceState string

const (
	AgreementAcceptanceState_ACCEPTED             AgreementAcceptanceState = "accepted"
	AgreementAcceptanceState_DECLINED             AgreementAcceptanceState = "declined"
	AgreementAcceptanceState_UNKNOWN_FUTURE_VALUE AgreementAcceptanceState = "unknownFutureValue"
)

func (v *AgreementAcceptanceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgreementAcceptanceState(value)
	for _, existing := range []AgreementAcceptanceState{"accepted", "declined", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgreementAcceptanceState", value)
}
