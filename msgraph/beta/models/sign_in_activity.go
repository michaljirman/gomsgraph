package models

import (
	"time"
)

// SignInActivity struct for MicrosoftGraphSignInActivity
type SignInActivity struct {
	// The last interactive sign-in date for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
	LastSignInDateTime *time.Time `json:"lastSignInDateTime,omitempty"`
	// Request ID of the last sign-in performed by this user.
	LastSignInRequestId *string `json:"lastSignInRequestId,omitempty"`
}
