package models

import (
	"encoding/json"
	"fmt"
)

// MailFolder struct for MailFolder
type MailFolder struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The number of immediate child mailFolders in the current mailFolder.
	ChildFolderCount *int `json:"childFolderCount,omitempty"`
	// The mailFolder's display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The unique identifier for the mailFolder's parent mailFolder.
	ParentFolderId *string `json:"parentFolderId,omitempty"`
	// The number of items in the mailFolder.
	TotalItemCount *int `json:"totalItemCount,omitempty"`
	// The number of items in the mailFolder marked as unread.
	UnreadItemCount *int `json:"unreadItemCount,omitempty"`
	// The collection of child folders in the mailFolder.
	ChildFolders *[]MailFolder `json:"childFolders,omitempty"`
	// The collection of rules that apply to the user's Inbox folder.
	MessageRules *[]MessageRule `json:"messageRules,omitempty"`
	// The collection of messages in the mailFolder.
	Messages *[]Message `json:"messages,omitempty"`
	// The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// MessageRule struct for MessageRule
type MessageRule struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Actions to be taken on a message when the corresponding conditions are fulfilled.
	Actions *MessageRuleActions `json:"actions,omitempty"`
	// Conditions that when fulfilled, will trigger the corresponding actions for that rule.
	Conditions *MessageRulePredicates `json:"conditions,omitempty"`
	// The display name of the rule.
	DisplayName *string `json:"displayName,omitempty"`
	// Exception conditions for the rule.
	Exceptions *MessageRulePredicates `json:"exceptions,omitempty"`
	// Indicates whether the rule is in an error condition. Read-only.
	HasError *bool `json:"hasError,omitempty"`
	// Indicates whether the rule is enabled to be applied to messages.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	// Indicates the order in which the rule is executed, among other rules.
	Sequence *int `json:"sequence,omitempty"`
}

// MessageRuleActions struct for MessageRuleActions
type MessageRuleActions struct {
	// A list of categories to be assigned to a message.
	AssignCategories *[]string `json:"assignCategories,omitempty"`
	// The ID of a folder that a message is to be copied to.
	CopyToFolder *string `json:"copyToFolder,omitempty"`
	// Indicates whether a message should be moved to the Deleted Items folder.
	Delete *bool `json:"delete,omitempty"`
	// The email addresses of the recipients to which a message should be forwarded as an attachment.
	ForwardAsAttachmentTo *[]Recipient `json:"forwardAsAttachmentTo,omitempty"`
	// The email addresses of the recipients to which a message should be forwarded.
	ForwardTo *[]Recipient `json:"forwardTo,omitempty"`
	// Indicates whether a message should be marked as read.
	MarkAsRead *bool `json:"markAsRead,omitempty"`
	// Sets the importance of the message, which can be: low, normal, high.
	MarkImportance *Importance `json:"markImportance,omitempty"`
	// The ID of the folder that a message will be moved to.
	MoveToFolder *string `json:"moveToFolder,omitempty"`
	// Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
	PermanentDelete *bool `json:"permanentDelete,omitempty"`
	// The email addresses to which a message should be redirected.
	RedirectTo *[]Recipient `json:"redirectTo,omitempty"`
	// Indicates whether subsequent rules should be evaluated.
	StopProcessingRules *bool `json:"stopProcessingRules,omitempty"`
}

// MessageRulePredicates struct for MessageRulePredicates
type MessageRulePredicates struct {
	// Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
	BodyContains *[]string `json:"bodyContains,omitempty"`
	// Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
	BodyOrSubjectContains *[]string `json:"bodyOrSubjectContains,omitempty"`
	// Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
	Categories *[]string `json:"categories,omitempty"`
	// Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
	FromAddresses *[]Recipient `json:"fromAddresses,omitempty"`
	// Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
	HeaderContains *[]string `json:"headerContains,omitempty"`
	// The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
	Importance *Importance `json:"importance,omitempty"`
	// Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
	IsApprovalRequest *bool `json:"isApprovalRequest,omitempty"`
	// Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
	IsAutomaticForward *bool `json:"isAutomaticForward,omitempty"`
	// Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
	IsAutomaticReply *bool `json:"isAutomaticReply,omitempty"`
	// Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
	IsMeetingRequest *bool `json:"isMeetingRequest,omitempty"`
	// Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
	IsMeetingResponse *bool `json:"isMeetingResponse,omitempty"`
	// Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
	IsNonDeliveryReport *bool `json:"isNonDeliveryReport,omitempty"`
	// Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
	IsPermissionControlled *bool `json:"isPermissionControlled,omitempty"`
	// Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
	IsReadReceipt *bool `json:"isReadReceipt,omitempty"`
	// Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
	IsSigned *bool `json:"isSigned,omitempty"`
	// Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
	IsVoicemail *bool `json:"isVoicemail,omitempty"`
	// Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
	MessageActionFlag *MessageActionFlag `json:"messageActionFlag,omitempty"`
	// Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
	NotSentToMe *bool `json:"notSentToMe,omitempty"`
	// Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
	RecipientContains *[]string `json:"recipientContains,omitempty"`
	// Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
	SenderContains *[]string `json:"senderContains,omitempty"`
	// Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
	SentCcMe *bool `json:"sentCcMe,omitempty"`
	// Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
	SentOnlyToMe *bool `json:"sentOnlyToMe,omitempty"`
	// Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
	SentToAddresses *[]Recipient `json:"sentToAddresses,omitempty"`
	// Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
	SentToMe *bool `json:"sentToMe,omitempty"`
	// Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
	SentToOrCcMe *bool `json:"sentToOrCcMe,omitempty"`
	// Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
	SubjectContains *[]string `json:"subjectContains,omitempty"`
	// Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
	WithinSizeRange *SizeRange `json:"withinSizeRange,omitempty"`
}

// Recipient struct for Recipient
type Recipient struct {
	// The recipient's email address.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
}

// SizeRange struct for SizeRange
type SizeRange struct {
	// The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
	MaximumSize *int `json:"maximumSize,omitempty"`
	// The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
	MinimumSize *int `json:"minimumSize,omitempty"`
}

// Sensitivity the model 'Sensitivity'
type Sensitivity string

// List of microsoft.graph.sensitivity
const (
	NORMAL       Sensitivity = "normal"
	PERSONAL     Sensitivity = "personal"
	PRIVATE      Sensitivity = "private"
	CONFIDENTIAL Sensitivity = "confidential"
)

func (v *Sensitivity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Sensitivity(value)
	for _, existing := range []Sensitivity{"normal", "personal", "private", "confidential"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Sensitivity", value)
}

// MessageActionFlag the model 'MessageActionFlag'
type MessageActionFlag string

// List of microsoft.graph.messageActionFlag
const (
	ANY                   MessageActionFlag = "any"
	CALL                  MessageActionFlag = "call"
	DO_NOT_FORWARD        MessageActionFlag = "doNotForward"
	FOLLOW_UP             MessageActionFlag = "followUp"
	FYI                   MessageActionFlag = "fyi"
	FORWARD               MessageActionFlag = "forward"
	NO_RESPONSE_NECESSARY MessageActionFlag = "noResponseNecessary"
	READ                  MessageActionFlag = "read"
	REPLY                 MessageActionFlag = "reply"
	REPLY_TO_ALL          MessageActionFlag = "replyToAll"
	REVIEW                MessageActionFlag = "review"
)

func (v *MessageActionFlag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageActionFlag(value)
	for _, existing := range []MessageActionFlag{"any", "call", "doNotForward", "followUp", "fyi", "forward", "noResponseNecessary", "read", "reply", "replyToAll", "review"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageActionFlag", value)
}

// Importance the model 'Importance'
type Importance string

// List of microsoft.graph.importance
const (
	Importance_LOW    Importance = "low"
	Importance_NORMAL Importance = "normal"
	Importance_HIGH   Importance = "high"
)

func (v *Importance) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Importance(value)
	for _, existing := range []Importance{"low", "normal", "high"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Importance", value)
}
