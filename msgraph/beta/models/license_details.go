package models

type LicenseDetails struct {
	Id *string `json:"id,omitempty"`
	// Information about the service plans assigned with the license. Read-only, Not nullable
	ServicePlans *[]ServicePlanInfo `json:"servicePlans,omitempty"`
	// Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only
	SkuId *string `json:"skuId,omitempty"`
	// Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: 'AAD_Premium'. Read-only
	SkuPartNumber *string `json:"skuPartNumber,omitempty"`
}
