package models

// TimeZoneBase struct for TimeZoneBase
type TimeZoneBase struct {
	// The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
	Name *string `json:"name,omitempty"`
}
