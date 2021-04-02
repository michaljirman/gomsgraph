package models

import (
	"encoding/json"
	"fmt"
)

type DelegateMeetingMessageDeliveryOptions string

// List of microsoft.graph.delegateMeetingMessageDeliveryOptions
const (
	SEND_TO_DELEGATE_AND_INFORMATION_TO_PRINCIPAL DelegateMeetingMessageDeliveryOptions = "sendToDelegateAndInformationToPrincipal"
	SEND_TO_DELEGATE_AND_PRINCIPAL                DelegateMeetingMessageDeliveryOptions = "sendToDelegateAndPrincipal"
	SEND_TO_DELEGATE_ONLY                         DelegateMeetingMessageDeliveryOptions = "sendToDelegateOnly"
)

func (v *DelegateMeetingMessageDeliveryOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DelegateMeetingMessageDeliveryOptions(value)
	for _, existing := range []DelegateMeetingMessageDeliveryOptions{"sendToDelegateAndInformationToPrincipal", "sendToDelegateAndPrincipal", "sendToDelegateOnly"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DelegateMeetingMessageDeliveryOptions", value)
}
