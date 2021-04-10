package models

type PhysicalAddress struct {
	// The city.
	City *string `json:"city,omitempty"`
	// The country or region. It's a free-format string value, for example, 'United States'.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The state.
	State *string `json:"state,omitempty"`
	// The street.
	Street *string `json:"street,omitempty"`
}
