package models

import "time"

type Contact struct {
	Id *string `json:"id,omitempty"`
	// The categories associated with the item
	Categories *[]string `json:"categories,omitempty"`
	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey *string `json:"changeKey,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the contact's assistant.
	AssistantName *string `json:"assistantName,omitempty"`
	// The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	Birthday *time.Time `json:"birthday,omitempty"`
	// The contact's business address.
	BusinessAddress *PhysicalAddress `json:"businessAddress,omitempty"`
	// The business home page of the contact.
	BusinessHomePage *string `json:"businessHomePage,omitempty"`
	// The contact's business phone numbers.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`
	// The names of the contact's children.
	Children *[]string `json:"children,omitempty"`
	// The name of the contact's company.
	CompanyName *string `json:"companyName,omitempty"`
	// The contact's department.
	Department *string `json:"department,omitempty"`
	// The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
	DisplayName *string `json:"displayName,omitempty"`
	// The contact's email addresses.
	EmailAddresses *[]EmailAddress `json:"emailAddresses,omitempty"`
	// The name the contact is filed under.
	FileAs *string `json:"fileAs,omitempty"`
	// The contact's generation.
	Generation *string `json:"generation,omitempty"`
	// The contact's given name.
	GivenName *string `json:"givenName,omitempty"`
	// The contact's home address.
	HomeAddress *PhysicalAddress `json:"homeAddress,omitempty"`
	// The contact's home phone numbers.
	HomePhones *[]string `json:"homePhones,omitempty"`
	// The contact's instant messaging (IM) addresses.
	ImAddresses *[]string `json:"imAddresses,omitempty"`
	// The contact's initials.
	Initials *string `json:"initials,omitempty"`
	// The contact’s job title.
	JobTitle *string `json:"jobTitle,omitempty"`
	// The name of the contact's manager.
	Manager *string `json:"manager,omitempty"`
	// The contact's middle name.
	MiddleName *string `json:"middleName,omitempty"`
	// The contact's mobile phone number.
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// The contact's nickname.
	NickName *string `json:"nickName,omitempty"`
	// The location of the contact's office.
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Other addresses for the contact.
	OtherAddress PhysicalAddress `json:"otherAddress,omitempty"`
	// The ID of the contact's parent folder.
	ParentFolderId *string `json:"parentFolderId,omitempty"`
	// The user's notes about the contact.
	PersonalNotes *string `json:"personalNotes,omitempty"`
	// The contact's profession.
	Profession *string `json:"profession,omitempty"`
	// The name of the contact's spouse/partner.
	SpouseName *string `json:"spouseName,omitempty"`
	// The contact's surname.
	Surname *string `json:"surname,omitempty"`
	// The contact's title.
	Title *string `json:"title,omitempty"`
	// The phonetic Japanese company name of the contact.
	YomiCompanyName *string `json:"yomiCompanyName,omitempty"`
	// The phonetic Japanese given name (first name) of the contact.
	YomiGivenName *string `json:"yomiGivenName,omitempty"`
	// The phonetic Japanese surname (last name)  of the contact.
	YomiSurname *string `json:"yomiSurname,omitempty"`
	// The collection of open extensions defined for the contact. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Optional contact picture. You can get or set a photo for a contact.
	Photo ProfilePhoto `json:"photo,omitempty"`
	// The collection of single-value extended properties defined for the contact. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
