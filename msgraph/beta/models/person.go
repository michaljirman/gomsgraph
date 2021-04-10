package models

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id *string `json:"id,omitempty"`
	// The person's birthday.
	Birthday *string `json:"birthday,omitempty"`
	// The name of the person's company.
	CompanyName *string `json:"companyName,omitempty"`
	// The person's department.
	Department *string `json:"department,omitempty"`
	// The person's display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The person's given name.
	GivenName *string `json:"givenName,omitempty"`
	// The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only.
	ImAddress *string `json:"imAddress,omitempty"`
	// true if the user has flagged this person as a favorite.
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// The person's job title.
	JobTitle *string `json:"jobTitle,omitempty"`
	// The location of the person's office.
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Free-form notes that the user has taken about this person.
	PersonNotes *string `json:"personNotes,omitempty"`
	// The type of person.
	PersonType *PersonType `json:"personType,omitempty"`
	// The person's phone numbers.
	Phones *[]Phone `json:"phones,omitempty"`
	// The person's addresses.
	PostalAddresses *[]Location `json:"postalAddresses,omitempty"`
	// The person's profession.
	Profession *string `json:"profession,omitempty"`
	// The person's email addresses.
	ScoredEmailAddresses *[]ScoredEmailAddress `json:"scoredEmailAddresses,omitempty"`
	// The person's surname.
	Surname *string `json:"surname,omitempty"`
	// The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// The person's websites.
	Websites *[]Website `json:"websites,omitempty"`
	// The phonetic Japanese name of the person's company.
	YomiCompany *string `json:"yomiCompany,omitempty"`
}

// PersonType struct for PersonType
type PersonType struct {
	// The type of data source, such as Person.
	Class *string `json:"class,omitempty"`
	// The secondary type of data source, such as OrganizationUser.
	Subclass *string `json:"subclass,omitempty"`
}

type Phone struct {
	Language *string `json:"language,omitempty"`
	// The phone number.
	Number *string `json:"number,omitempty"`
	Region *string `json:"region,omitempty"`
	// The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
	Type *PhoneType `json:"type,omitempty"`
}

type Website struct {
	// The URL of the website.
	Address *string `json:"address,omitempty"`
	// The display name of the web site.
	DisplayName *string `json:"displayName,omitempty"`
	// The possible values are: other, home, work, blog, profile.
	Type WebsiteType `json:"type,omitempty"`
}

type WebsiteType string

const (
	WebsiteType_OTHER   WebsiteType = "other"
	WebsiteType_HOME    WebsiteType = "home"
	WebsiteType_WORK    WebsiteType = "work"
	WebsiteType_BLOG    WebsiteType = "blog"
	WebsiteType_PROFILE WebsiteType = "profile"
)

func (v *WebsiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebsiteType(value)
	for _, existing := range []WebsiteType{"other", "home", "work", "blog", "profile"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebsiteType", value)
}

type ScoredEmailAddress struct {
	// The email address.
	Address *string `json:"address,omitempty"`
	ItemId  *string `json:"itemId,omitempty"`
	// The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
	RelevanceScore      *float32 `json:"relevanceScore,omitempty"`
	SelectionLikelihood *float32 `json:"selectionLikelihood,omitempty"`
}
