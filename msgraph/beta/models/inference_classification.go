package models

import (
	"encoding/json"
	"fmt"
)

type InferenceClassification struct {
	Id *string `json:"id,omitempty"`
	// A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
	Overrides *[]InferenceClassificationOverride `json:"overrides,omitempty"`
}

type InferenceClassificationOverride struct {
	Id *string `json:"id,omitempty"`
	// Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
	ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`
	// The email address information of the sender for whom the override is created.
	SenderEmailAddress *EmailAddress `json:"senderEmailAddress,omitempty"`
}

type InferenceClassificationType string

const (
	InferenceClassificationType_FOCUSED InferenceClassificationType = "focused"
	InferenceClassificationType_OTHER   InferenceClassificationType = "other"
)

func (v *InferenceClassificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InferenceClassificationType(value)
	for _, existing := range []InferenceClassificationType{"focused", "other"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InferenceClassificationType", value)
}
