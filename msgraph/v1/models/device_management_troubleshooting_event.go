package models

import "time"

// DeviceManagementTroubleshootingEvent struct for DeviceManagementTroubleshootingEvent
type DeviceManagementTroubleshootingEvent struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Id used for tracing the failure in the service.
	CorrelationId *string `json:"correlationId,omitempty"`
	// Time when the event occurred .
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
}
