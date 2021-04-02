package models

import (
	"encoding/json"
	"fmt"
)

// OutlookUser struct for OutlookUser
type OutlookUser struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A list of categories defined for the user.
	MasterCategories *[]OutlookCategory `json:"masterCategories,omitempty"`
}

// OutlookCategory struct for OutlookCategory
type OutlookCategory struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
	Color *CategoryColor `json:"color,omitempty"`
	// A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
}

// CategoryColor the model 'CategoryColor'
type CategoryColor string

// List of microsoft.graph.categoryColor
const (
	NONE     CategoryColor = "none"
	PRESET0  CategoryColor = "preset0"
	PRESET1  CategoryColor = "preset1"
	PRESET2  CategoryColor = "preset2"
	PRESET3  CategoryColor = "preset3"
	PRESET4  CategoryColor = "preset4"
	PRESET5  CategoryColor = "preset5"
	PRESET6  CategoryColor = "preset6"
	PRESET7  CategoryColor = "preset7"
	PRESET8  CategoryColor = "preset8"
	PRESET9  CategoryColor = "preset9"
	PRESET10 CategoryColor = "preset10"
	PRESET11 CategoryColor = "preset11"
	PRESET12 CategoryColor = "preset12"
	PRESET13 CategoryColor = "preset13"
	PRESET14 CategoryColor = "preset14"
	PRESET15 CategoryColor = "preset15"
	PRESET16 CategoryColor = "preset16"
	PRESET17 CategoryColor = "preset17"
	PRESET18 CategoryColor = "preset18"
	PRESET19 CategoryColor = "preset19"
	PRESET20 CategoryColor = "preset20"
	PRESET21 CategoryColor = "preset21"
	PRESET22 CategoryColor = "preset22"
	PRESET23 CategoryColor = "preset23"
	PRESET24 CategoryColor = "preset24"
)

func (v *CategoryColor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CategoryColor(value)
	for _, existing := range []CategoryColor{"none", "preset0", "preset1", "preset2", "preset3", "preset4", "preset5", "preset6", "preset7", "preset8", "preset9", "preset10", "preset11", "preset12", "preset13", "preset14", "preset15", "preset16", "preset17", "preset18", "preset19", "preset20", "preset21", "preset22", "preset23", "preset24"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CategoryColor", value)
}
