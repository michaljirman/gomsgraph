package models

import (
	"encoding/json"
	"fmt"
)

// OnlineMeetingProviderType the model 'OnlineMeetingProviderType'
type OnlineMeetingProviderType string

// List of microsoft.graph.onlineMeetingProviderType
const (
	UNKNOWN            OnlineMeetingProviderType = "unknown"
	SKYPE_FOR_BUSINESS OnlineMeetingProviderType = "skypeForBusiness"
	SKYPE_FOR_CONSUMER OnlineMeetingProviderType = "skypeForConsumer"
	TEAMS_FOR_BUSINESS OnlineMeetingProviderType = "teamsForBusiness"
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
