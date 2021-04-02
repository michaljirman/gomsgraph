package models

import "time"

// Site struct for Site
type Site struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
	// The full title for the site. Read-only.
	DisplayName *string      `json:"displayName,omitempty"`
	Error       *PublicError `json:"error,omitempty"`
	// If present, indicates that this is the root site in the site collection. Read-only.
	Root *interface{} `json:"root,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	// Provides details about the site's site collection. Available only on the root site. Read-only.
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`
	// Analytics about the view activities that took place in this site.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// The collection of column definitions reusable across lists under this site.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`
	// The collection of content types defined for this site.
	ContentTypes *[]ContentType `json:"contentTypes,omitempty"`
	// The default drive (document library) for this site.
	Drive *Drive `json:"drive,omitempty"`
	// The collection of drives (document libraries) under this site.
	Drives *[]Drive `json:"drives,omitempty"`
	// Used to address any item contained in this site. This collection cannot be enumerated.
	Items *[]BaseItem `json:"items,omitempty"`
	// The collection of lists under this site.
	Lists *[]List `json:"lists,omitempty"`
	// The collection of the sub-sites under this site.
	Sites *[]Site `json:"sites,omitempty"`
	// Calls the OneNote service for notebook related operations.
	Onenote *Onenote `json:"onenote,omitempty"`
}

// BaseItem struct for BaseItem
type BaseItem struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Identity of the user, device, or application which created the item. Read-only.
	//CreatedBy NullableAnyOfIdentitySet `json:"createdBy,omitempty"`
	//// Date and time of item creation. Read-only.
	//CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	//// Provides a user-visible description of the item. Optional.
	//Description *string `json:"description,omitempty"`
	//// ETag for the item. Read-only.
	//ETag *string `json:"eTag,omitempty"`
	//// Identity of the user, device, and application which last modified the item. Read-only.
	//LastModifiedBy NullableAnyOfIdentitySet `json:"lastModifiedBy,omitempty"`
	//// Date and time the item was last modified. Read-only.
	//LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	//// The name of the item. Read-write.
	//Name *string `json:"name,omitempty"`
	//// Parent information, if the item has a parent. Read-write.
	//ParentReference NullableAnyOfItemReference `json:"parentReference,omitempty"`
	//// URL that displays the resource in the browser. Read-only.
	//WebUrl *string `json:"webUrl,omitempty"`
	//// Identity of the user who created the item. Read-only.
	//CreatedByUser NullableAnyOfUser `json:"createdByUser,omitempty"`
	//// Identity of the user who last modified the item. Read-only.
	//LastModifiedByUser NullableAnyOfUser `json:"lastModifiedByUser,omitempty"`
}

// SiteCollection struct for SiteCollection
type SiteCollection struct {
	//// The geographic region code for where this site collection resides. Read-only.
	//DataLocationCode *string `json:"dataLocationCode,omitempty"`
	//// The hostname for the site collection. Read-only.
	//Hostname *string `json:"hostname,omitempty"`
	//// If present, indicates that this is a root site collection in SharePoint. Read-only.
	//Root NullableAnyOfobject `json:"root,omitempty"`
}

// PublicError struct for PublicError
type PublicError struct {
	//Code       *string                              `json:"code,omitempty"`
	//Details    *[]AnyOfPublicErrorDetail     `json:"details,omitempty"`
	//InnerError NullableAnyOfPublicInnerError `json:"innerError,omitempty"`
	//Message    *string                              `json:"message,omitempty"`
	//Target     *string                              `json:"target,omitempty"`
}
