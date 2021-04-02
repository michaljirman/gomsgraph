package models

// MailboxSettings struct for MailboxSettings
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
	// The days of the week and hours in a specific time zone that the user works.
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}
