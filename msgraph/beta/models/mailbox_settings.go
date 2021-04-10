package models

import (
	"encoding/json"
	"fmt"
)

type MailboxSettings struct {
	// Folder ID of an archive folder for the user.
	ArchiveFolder *string `json:"archiveFolder,omitempty"`
	// Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
	AutomaticRepliesSetting *AutomaticRepliesSetting `json:"automaticRepliesSetting,omitempty"`
	// The date format for the user's mailbox.
	DateFormat *string `json:"dateFormat,omitempty"`
	// If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
	DelegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions `json:"delegateMeetingMessageDeliveryOptions,omitempty"`
	// The locale information for the user, including the preferred language and country/region.
	Language *LocaleInfo `json:"language,omitempty"`
	// The time format for the user's mailbox.
	TimeFormat *string `json:"timeFormat,omitempty"`
	// The default time zone for the user's mailbox.
	TimeZone *string `json:"timeZone,omitempty"`
	// The purpose of the mailbox. Used to differentiate a mailbox for a single user from a shared mailbox and equipment mailbox in Exchange Online. Read only.
	UserPurpose   *UserPurpose          `json:"userPurpose,omitempty"`
	UserPurposeV2 *MailboxRecipientType `json:"userPurposeV2,omitempty"`
	// The days of the week and hours in a specific time zone that the user works.
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}

type UserPurpose struct {
	// Represents the user's recipient or mailbox type in Exchange Online. Possible values are: unknown, user, linked, shared, room, equipment, and others. See the next section for more information.
	Value *MailboxRecipientType `json:"value,omitempty"`
}

type MailboxRecipientType string

const (
	MAILBOXRECIPIENTTYPE_UNKNOWN   MailboxRecipientType = "unknown"
	MAILBOXRECIPIENTTYPE_USER      MailboxRecipientType = "user"
	MAILBOXRECIPIENTTYPE_LINKED    MailboxRecipientType = "linked"
	MAILBOXRECIPIENTTYPE_SHARED    MailboxRecipientType = "shared"
	MAILBOXRECIPIENTTYPE_ROOM      MailboxRecipientType = "room"
	MAILBOXRECIPIENTTYPE_EQUIPMENT MailboxRecipientType = "equipment"
	MAILBOXRECIPIENTTYPE_OTHERS    MailboxRecipientType = "others"
)

func (v *MailboxRecipientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MailboxRecipientType(value)
	for _, existing := range []MailboxRecipientType{"unknown", "user", "linked", "shared", "room", "equipment", "others"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MailboxRecipientType", value)
}
