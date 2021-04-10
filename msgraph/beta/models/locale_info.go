package models

type LocaleInfo struct {
	// A name representing the user's locale in natural language, for example, 'English (United States)'.
	DisplayName *string `json:"displayName,omitempty"`
	// A locale representation for the user, which includes the user's preferred language and country/region. For example, 'en-us'. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2.
	Locale *string `json:"locale,omitempty"`
}
