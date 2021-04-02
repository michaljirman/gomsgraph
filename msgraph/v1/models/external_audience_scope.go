package models

import (
	"encoding/json"
	"fmt"
)

// ExternalAudienceScope the model 'ExternalAudienceScope'
type ExternalAudienceScope string

// List of microsoft.graph.externalAudienceScope
const (
	ExternalAudienceScope_NONE          ExternalAudienceScope = "none"
	ExternalAudienceScope_CONTACTS_ONLY ExternalAudienceScope = "contactsOnly"
	ExternalAudienceScope_ALL           ExternalAudienceScope = "all"
)

func (v *ExternalAudienceScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalAudienceScope(value)
	for _, existing := range []ExternalAudienceScope{"none", "contactsOnly", "all"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalAudienceScope", value)
}
