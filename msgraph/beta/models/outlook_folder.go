package models

import (
	"encoding/json"
	"fmt"
)

type OutlookUser struct {
	Id *string `json:"id,omitempty"`
	// A list of categories defined for the user.
	MasterCategories *[]OutlookCategory `json:"masterCategories,omitempty"`
}

type OutlookCategory struct {
	Id *string `json:"id,omitempty"`
	// A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
	Color *CategoryColor `json:"color,omitempty"`
	// A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
	DisplayName *string `json:"displayName,omitempty"`
}

type CategoryColor string

const (
	CategoryColor_NONE CategoryColor = "none"
	CategoryColor_0    CategoryColor = "preset0"
	CategoryColor_1    CategoryColor = "preset1"
	CategoryColor_2    CategoryColor = "preset2"
	CategoryColor_3    CategoryColor = "preset3"
	CategoryColor_4    CategoryColor = "preset4"
	CategoryColor_5    CategoryColor = "preset5"
	CategoryColor_6    CategoryColor = "preset6"
	CategoryColor_7    CategoryColor = "preset7"
	CategoryColor_8    CategoryColor = "preset8"
	CategoryColor_9    CategoryColor = "preset9"
	CategoryColor_10   CategoryColor = "preset10"
	CategoryColor_11   CategoryColor = "preset11"
	CategoryColor_12   CategoryColor = "preset12"
	CategoryColor_13   CategoryColor = "preset13"
	CategoryColor_14   CategoryColor = "preset14"
	CategoryColor_15   CategoryColor = "preset15"
	CategoryColor_16   CategoryColor = "preset16"
	CategoryColor_17   CategoryColor = "preset17"
	CategoryColor_18   CategoryColor = "preset18"
	CategoryColor_19   CategoryColor = "preset19"
	CategoryColor_20   CategoryColor = "preset20"
	CategoryColor_21   CategoryColor = "preset21"
	CategoryColor_22   CategoryColor = "preset22"
	CategoryColor_24   CategoryColor = "preset24"
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
