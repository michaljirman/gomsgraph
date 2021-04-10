package models

type Message struct {
	Id *string `json:"id,omitempty"`
	//// The categories associated with the item
	//Categories *[]string `json:"categories,omitempty"`
	//// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	//ChangeKey *string `json:"changeKey,omitempty"`
	//// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	//CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	//// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	//LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	//// The Bcc: recipients for the message.
	//BccRecipients *[]AnyOfRecipient `json:"bccRecipients,omitempty"`
	//// The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
	//Body *ItemBody `json:"body,omitempty"`
	//// The first 255 characters of the message body. It is in text format.
	//BodyPreview *string `json:"bodyPreview,omitempty"`
	//// The Cc: recipients for the message.
	//CcRecipients *[]AnyOfRecipient `json:"ccRecipients,omitempty"`
	//// The ID of the conversation the email belongs to.
	//ConversationId *string `json:"conversationId,omitempty"`
	//// Indicates the position of the message within the conversation.
	//ConversationIndex *string `json:"conversationIndex,omitempty"`
	//// The flag value that indicates the status, start date, due date, or completion date for the message.
	//Flag *FollowupFlag `json:"flag,omitempty"`
	//// The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
	//From *Recipient `json:"from,omitempty"`
	//// Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
	//HasAttachments *bool `json:"hasAttachments,omitempty"`
	//// The importance of the message: Low, Normal, High.
	//Importance *Importance `json:"importance,omitempty"`
	//// The classification of the message for the user, based on inferred relevance or importance, or on an explicit override. The possible values are: focused or other.
	//InferenceClassification *InferenceClassificationType `json:"inferenceClassification,omitempty"`
	//// A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
	//InternetMessageHeaders *[]AnyOfInternetMessageHeader `json:"internetMessageHeaders,omitempty"`
	//// The message ID in the format specified by RFC2822.
	//InternetMessageId *string `json:"internetMessageId,omitempty"`
	//// Indicates whether a read receipt is requested for the message.
	//IsDeliveryReceiptRequested *bool `json:"isDeliveryReceiptRequested,omitempty"`
	//// Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
	//IsDraft *bool `json:"isDraft,omitempty"`
	//// Indicates whether the message has been read.
	//IsRead *bool `json:"isRead,omitempty"`
	//// Indicates whether a read receipt is requested for the message.
	//IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty"`
	//// The unique identifier for the message's parent mailFolder.
	//ParentFolderId *string `json:"parentFolderId,omitempty"`
	//// The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
	//ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`
	//// The email addresses to use when replying.
	//ReplyTo *[]AnyOfRecipient `json:"replyTo,omitempty"`
	//// The account that is actually used to generate the message. In most cases, this value is the same as the from property. You can set this property to a different value when sending a message from a shared mailbox, for a shared calendar, or as a delegate. In any case, the value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
	//Sender *Recipient `json:"sender,omitempty"`
	//// The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
	//SentDateTime *time.Time `json:"sentDateTime,omitempty"`
	//// The subject of the message.
	//Subject *string `json:"subject,omitempty"`
	//// The To: recipients for the message.
	//ToRecipients *[]AnyOfRecipient `json:"toRecipients,omitempty"`
	//// The part of the body of the message that is unique to the current message. uniqueBody is not returned by default but can be retrieved for a given message by use of the ?$select=uniqueBody query. It can be in HTML or text format.
	//UniqueBody *ItemBody `json:"uniqueBody,omitempty"`
	//// The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
	//WebLink *string `json:"webLink,omitempty"`
	//// The fileAttachment and itemAttachment attachments for the message.
	//Attachments *[]Attachment `json:"attachments,omitempty"`
	//// The collection of open extensions defined for the message. Nullable.
	//Extensions *[]Extension `json:"extensions,omitempty"`
	//// The collection of multi-value extended properties defined for the message. Nullable.
	//MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	//// The collection of single-value extended properties defined for the message. Nullable.
	//SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
