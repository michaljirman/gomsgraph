package models

// PasswordProfile struct for MicrosoftGraphPasswordProfile
type PasswordProfile struct {
	// If true, at next sign-in, the user must change their password. After a password change, this property will be automatically reset to false. If not set, default is false.
	ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
	// If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
	ForceChangePasswordNextSignInWithMfa *bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
	// The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
	Password *string `json:"password,omitempty"`
}
