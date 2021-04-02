package models

type OData struct {
	Context  *string `json:"@odata.context,omitempty"`
	ID       *string `json:"@odata.id,omitempty"`
	Count    *int    `json:"@odata.count,omitempty"`
	NextLink *string `json:"@odata.nextLink,omitempty"`
}

var ODataFields = map[string]struct{}{
	"@odata.context":  {},
	"@odata.id":       {},
	"@odata.count":    {},
	"@odata.nextLink": {},
}
