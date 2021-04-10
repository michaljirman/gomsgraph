package models

import (
	"encoding/json"
	"fmt"
)

type OnlineMeetingProviderType string

const (
	ONLINE_MEETING_PROVIDER_TYPE_UNKNOWN            OnlineMeetingProviderType = "unknown"
	ONLINE_MEETING_PROVIDER_TYPE_SKYPE_FOR_BUSINESS OnlineMeetingProviderType = "skypeForBusiness"
	ONLINE_MEETING_PROVIDER_TYPE_SKYPE_FOR_CONSUMER OnlineMeetingProviderType = "skypeForConsumer"
	ONLINE_MEETING_PROVIDER_TYPE_TEAMS_FOR_BUSINESS OnlineMeetingProviderType = "teamsForBusiness"
)

func (v *OnlineMeetingProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OnlineMeetingProviderType(value)
	for _, existing := range []OnlineMeetingProviderType{"unknown", "skypeForBusiness", "skypeForConsumer", "teamsForBusiness"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OnlineMeetingProviderType", value)
}
