package models

// ItemReference struct for ItemReference
type ItemReference struct {
	// Unique identifier of the drive instance that contains the item. Read-only.
	DriveId *string `json:"driveId,omitempty"`
	// Identifies the type of drive. See [drive][] resource for values.
	DriveType *string `json:"driveType,omitempty"`
	// Unique identifier of the item in the drive. Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the item being referenced. Read-only.
	Name *string `json:"name,omitempty"`
	// Path that can be used to navigate to the item. Read-only.
	Path *string `json:"path,omitempty"`
	// A unique identifier for a shared resource that can be accessed via the [Shares][] API.
	ShareId *string `json:"shareId,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	// For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated.
	SiteId *string `json:"siteId,omitempty"`
}

// SharepointIds struct for SharepointIds
type SharepointIds struct {
	// The unique identifier (guid) for the item's list in SharePoint.
	ListId *string `json:"listId,omitempty"`
	// An integer identifier for the item within the containing list.
	ListItemId *string `json:"listItemId,omitempty"`
	// The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
	ListItemUniqueId *string `json:"listItemUniqueId,omitempty"`
	// The unique identifier (guid) for the item's site collection (SPSite).
	SiteId *string `json:"siteId,omitempty"`
	// The SharePoint URL for the site that contains the item.
	SiteUrl *string `json:"siteUrl,omitempty"`
	// The unique identifier (guid) for the tenancy.
	TenantId *string `json:"tenantId,omitempty"`
	// The unique identifier (guid) for the item's site (SPWeb).
	WebId *string `json:"webId,omitempty"`
}
