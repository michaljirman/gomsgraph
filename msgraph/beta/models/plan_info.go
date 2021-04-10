package models

type ServicePlanInfo struct {
	// The object the service plan can be assigned to. Possible values:'User' - service plan can be assigned to individual users.'Company' - service plan can be assigned to the entire tenant.
	AppliesTo *string `json:"appliesTo,omitempty"`
	// The provisioning status of the service plan. Possible values:'Success' - Service is fully provisioned.'Disabled' - Service has been disabled.'PendingInput' - Service is not yet provisioned; awaiting service confirmation.'PendingActivation' - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)'PendingProvisioning' - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	// The unique identifier of the service plan.
	ServicePlanId *string `json:"servicePlanId,omitempty"`
	// The name of the service plan.
	ServicePlanName *string `json:"servicePlanName,omitempty"`
}
