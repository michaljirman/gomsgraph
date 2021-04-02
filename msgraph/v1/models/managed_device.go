package models

// ManagedDevice struct for ManagedDevice
type ManagedDevice struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Code that allows the Activation Lock on a device to be bypassed.
	//ActivationLockBypassCode *string `json:"activationLockBypassCode,omitempty"`
	//// Android security patch level
	//AndroidSecurityPatchLevel *string `json:"androidSecurityPatchLevel,omitempty"`
	//// The unique identifier for the Azure Active Directory device. Read only.
	//AzureADDeviceId *string `json:"azureADDeviceId,omitempty"`
	//// Whether the device is Azure Active Directory registered.
	//AzureADRegistered *bool `json:"azureADRegistered,omitempty"`
	//// The DateTime when device compliance grace period expires
	//ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	//// Compliance state of the device. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
	//ComplianceState *AnyOfComplianceState `json:"complianceState,omitempty"`
	//// ConfigrMgr client enabled features
	//ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`
	//// List of ComplexType deviceActionResult objects.
	//DeviceActionResults *[]AnyOfDeviceActionResult `json:"deviceActionResults,omitempty"`
	//// Device category display name
	//DeviceCategoryDisplayName *string `json:"deviceCategoryDisplayName,omitempty"`
	//// Enrollment type of the device. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
	//DeviceEnrollmentType *AnyOfDeviceEnrollmentType `json:"deviceEnrollmentType,omitempty"`
	//// The device health attestation state.
	//DeviceHealthAttestationState *DeviceHealthAttestationState `json:"deviceHealthAttestationState,omitempty"`
	//// Name of the device
	//DeviceName *string `json:"deviceName,omitempty"`
	//// Device registration state. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
	//DeviceRegistrationState *AnyOfDeviceRegistrationState `json:"deviceRegistrationState,omitempty"`
	//// Whether the device is Exchange ActiveSync activated.
	//EasActivated *bool `json:"easActivated,omitempty"`
	//// Exchange ActivationSync activation time of the device.
	//EasActivationDateTime *time.Time `json:"easActivationDateTime,omitempty"`
	//// Exchange ActiveSync Id of the device.
	//EasDeviceId *string `json:"easDeviceId,omitempty"`
	//// Email(s) for the user associated with the device
	//EmailAddress *string `json:"emailAddress,omitempty"`
	//// Enrollment time of the device.
	//EnrolledDateTime *time.Time `json:"enrolledDateTime,omitempty"`
	//// The Access State of the device in Exchange. Possible values are: none, unknown, allowed, blocked, quarantined.
	//ExchangeAccessState *AnyOfDeviceManagementExchangeAccessState `json:"exchangeAccessState,omitempty"`
	//// The reason for the device's access state in Exchange. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
	//ExchangeAccessStateReason *AnyOfDeviceManagementExchangeAccessStateReason `json:"exchangeAccessStateReason,omitempty"`
	//// Last time the device contacted Exchange.
	//ExchangeLastSuccessfulSyncDateTime *time.Time `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`
	//// Free Storage in Bytes
	//FreeStorageSpaceInBytes *int64 `json:"freeStorageSpaceInBytes,omitempty"`
	//// IMEI
	//Imei *string `json:"imei,omitempty"`
	//// Device encryption status
	//IsEncrypted *bool `json:"isEncrypted,omitempty"`
	//// Device supervised status
	//IsSupervised *bool `json:"isSupervised,omitempty"`
	//// whether the device is jail broken or rooted.
	//JailBroken *string `json:"jailBroken,omitempty"`
	//// The date and time that the device last completed a successful sync with Intune.
	//LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	//// Automatically generated name to identify a device. Can be overwritten to a user friendly name.
	//ManagedDeviceName *string `json:"managedDeviceName,omitempty"`
	//// Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
	//ManagedDeviceOwnerType *AnyOfManagedDeviceOwnerType `json:"managedDeviceOwnerType,omitempty"`
	//// Management channel of the device. Intune, EAS, etc. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
	//ManagementAgent *AnyOfManagementAgentType `json:"managementAgent,omitempty"`
	//// Manufacturer of the device
	//Manufacturer *string `json:"manufacturer,omitempty"`
	//// MEID
	//Meid *string `json:"meid,omitempty"`
	//// Model of the device
	//Model *string `json:"model,omitempty"`
	//// Operating system of the device. Windows, iOS, etc.
	//OperatingSystem *string `json:"operatingSystem,omitempty"`
	//// Operating system version of the device.
	//OsVersion *string `json:"osVersion,omitempty"`
	//// Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
	//PartnerReportedThreatState *AnyOfManagedDevicePartnerReportedHealthState `json:"partnerReportedThreatState,omitempty"`
	//// Phone number of the device
	//PhoneNumber *string `json:"phoneNumber,omitempty"`
	//// An error string that identifies issues when creating Remote Assistance session objects.
	//RemoteAssistanceSessionErrorDetails *string `json:"remoteAssistanceSessionErrorDetails,omitempty"`
	//// Url that allows a Remote Assistance session to be established with the device.
	//RemoteAssistanceSessionUrl *string `json:"remoteAssistanceSessionUrl,omitempty"`
	//// SerialNumber
	//SerialNumber *string `json:"serialNumber,omitempty"`
	//// Subscriber Carrier
	//SubscriberCarrier *string `json:"subscriberCarrier,omitempty"`
	//// Total Storage in Bytes
	//TotalStorageSpaceInBytes *int64 `json:"totalStorageSpaceInBytes,omitempty"`
	//// User display name
	//UserDisplayName *string `json:"userDisplayName,omitempty"`
	//// Unique Identifier for the user associated with the device
	//UserId *string `json:"userId,omitempty"`
	//// Device user principal name
	//UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	//// Wi-Fi MAC
	//WiFiMacAddress *string `json:"wiFiMacAddress,omitempty"`
	//// Device compliance policy states for this device.
	//DeviceCompliancePolicyStates *[]DeviceCompliancePolicyState `json:"deviceCompliancePolicyStates,omitempty"`
	//// Device configuration states for this device.
	//DeviceConfigurationStates *[]DeviceConfigurationState `json:"deviceConfigurationStates,omitempty"`
	//// Device category
	//DeviceCategory *DeviceCategory `json:"deviceCategory,omitempty"`
}
