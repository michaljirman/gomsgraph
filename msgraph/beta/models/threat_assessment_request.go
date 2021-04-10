package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type ThreatAssessmentRequest struct {
	// The threat category. Possible values are: spam, phishing, malware.
	Category *ThreatCategory `json:"category,omitempty"`
	// The content type of threat assessment. Possible values are: mail, url, file.
	ContentType *ThreatAssessmentContentType `json:"contentType,omitempty"`
	// The threat assessment request creator.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The expected assessment from submitter. Possible values are: block, unblock.
	ExpectedAssessment *ThreatExpectedAssessment `json:"expectedAssessment,omitempty"`
	// The source of the threat assessment request. Possible values are: user, administrator.
	RequestSource *ThreatAssessmentRequestSource `json:"requestSource,omitempty"`
	// The assessment process status. Possible values are: pending, completed.
	Status *ThreatAssessmentStatus `json:"status,omitempty"`
	// A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
	Results *[]ThreatAssessmentResult `json:"results,omitempty"`
}

type ThreatCategory string

const (
	ThreatCategory_UNDEFINED            ThreatCategory = "undefined"
	ThreatCategory_SPAM                 ThreatCategory = "spam"
	ThreatCategory_PHISHING             ThreatCategory = "phishing"
	ThreatCategory_MALWARE              ThreatCategory = "malware"
	ThreatCategory_UNKNOWN_FUTURE_VALUE ThreatCategory = "unknownFutureValue"
)

func (v *ThreatCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatCategory(value)
	for _, existing := range []ThreatCategory{"undefined", "spam", "phishing", "malware", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatCategory", value)
}

type ThreatAssessmentContentType string

const (
	ThreatAssessmentContentType_MAIL ThreatAssessmentContentType = "mail"
	ThreatAssessmentContentType_URL  ThreatAssessmentContentType = "url"
	ThreatAssessmentContentType_FILE ThreatAssessmentContentType = "file"
)

func (v *ThreatAssessmentContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatAssessmentContentType(value)
	for _, existing := range []ThreatAssessmentContentType{"mail", "url", "file"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatAssessmentContentType", value)
}

type ThreatExpectedAssessment string

const (
	ThreatExpectedAssessment_BLOCK   ThreatExpectedAssessment = "block"
	ThreatExpectedAssessment_UNBLOCK ThreatExpectedAssessment = "unblock"
)

func (v *ThreatExpectedAssessment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatExpectedAssessment(value)
	for _, existing := range []ThreatExpectedAssessment{"block", "unblock"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatExpectedAssessment", value)
}

type ThreatAssessmentRequestSource string

const (
	ThreatAssessmentRequestSource_UNDEFINED     ThreatAssessmentRequestSource = "undefined"
	ThreatAssessmentRequestSource_USER          ThreatAssessmentRequestSource = "user"
	ThreatAssessmentRequestSource_ADMINISTRATOR ThreatAssessmentRequestSource = "administrator"
)

func (v *ThreatAssessmentRequestSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatAssessmentRequestSource(value)
	for _, existing := range []ThreatAssessmentRequestSource{"undefined", "user", "administrator"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatAssessmentRequestSource", value)
}

type ThreatAssessmentStatus string

const (
	PENDING   ThreatAssessmentStatus = "pending"
	COMPLETED ThreatAssessmentStatus = "completed"
)

func (v *ThreatAssessmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatAssessmentStatus(value)
	for _, existing := range []ThreatAssessmentStatus{"pending", "completed"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatAssessmentStatus", value)
}

type ThreatAssessmentResult struct {
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The result message for each threat assessment.
	Message *string `json:"message,omitempty"`
	// The threat assessment result type. Possible values are: checkPolicy (only for mail assessment), rescan.
	ResultType *ThreatAssessmentResultType `json:"resultType,omitempty"`
}

type ThreatAssessmentResultType string

const (
	ThreatAssessmentResultType_CHECK_POLICY         ThreatAssessmentResultType = "checkPolicy"
	ThreatAssessmentResultType_RESCAN               ThreatAssessmentResultType = "rescan"
	ThreatAssessmentResultType_UNKNOWN_FUTURE_VALUE ThreatAssessmentResultType = "unknownFutureValue"
)

func (v *ThreatAssessmentResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThreatAssessmentResultType(value)
	for _, existing := range []ThreatAssessmentResultType{"checkPolicy", "rescan", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThreatAssessmentResultType", value)
}
