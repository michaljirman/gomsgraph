package models

// DeviceKey struct for MicrosoftGraphDeviceKey
type DeviceKey struct {
	DeviceId    *string `json:"deviceId,omitempty"`
	KeyMaterial *string `json:"keyMaterial,omitempty"`
	KeyType     *string `json:"keyType,omitempty"`
}
