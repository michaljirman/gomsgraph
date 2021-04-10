package models

type ListOptions struct {
	// The $count and $search parameters are currently not available in Azure AD B2C tenants.
	//Search string `url:"$search,omitempty"`
	//Count  bool   `url:"$count,omitempty"`
	Filter   string `url:"$filter,omitempty"`
	Select   string `url:"$select,omitempty"`
	Top      int    `url:"$top,omitempty"`
	NextLink string `url:"-"`
}
