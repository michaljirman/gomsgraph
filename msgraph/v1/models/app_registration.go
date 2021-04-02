package models

import "time"

// ManagedAppRegistration struct for ManagedAppRegistration
type ManagedAppRegistration struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The app package Identifier
	AppIdentifier *interface{} `json:"appIdentifier,omitempty"`
	// App version
	ApplicationVersion *string `json:"applicationVersion,omitempty"`
	// Date and time of creation
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Host device name
	DeviceName *string `json:"deviceName,omitempty"`
	// App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
	DeviceTag *string `json:"deviceTag,omitempty"`
	// Host device type
	DeviceType *string `json:"deviceType,omitempty"`
	// Zero or more reasons an app registration is flagged. E.g. app running on rooted device
	FlaggedReasons *[]ManagedAppFlaggedReason `json:"flaggedReasons,omitempty"`
	// Date and time of last the app synced with management service.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// App management SDK version
	ManagementSdkVersion *string `json:"managementSdkVersion,omitempty"`
	// Operating System version
	PlatformVersion *string `json:"platformVersion,omitempty"`
	// The user Id to who this app registration belongs.
	UserId *string `json:"userId,omitempty"`
	// Version of the entity.
	Version *string `json:"version,omitempty"`
	// Zero or more policys already applied on the registered app when it last synchronized with managment service.
	AppliedPolicies *[]ManagedAppPolicy `json:"appliedPolicies,omitempty"`
	// Zero or more policies admin intended for the app as of now.
	IntendedPolicies *[]ManagedAppPolicy `json:"intendedPolicies,omitempty"`
	// Zero or more long running operations triggered on the app registration.
	Operations *[]ManagedAppOperation `json:"operations,omitempty"`
}
