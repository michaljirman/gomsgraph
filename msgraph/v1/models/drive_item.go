package models

// DriveItem struct for DriveItem
type DriveItem struct {
	//// Read-only.
	//Id *string `json:"id,omitempty"`
	//// Identity of the user, device, or application which created the item. Read-only.
	//CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	//// Date and time of item creation. Read-only.
	//CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	//// Provides a user-visible description of the item. Optional.
	//Description *string `json:"description,omitempty"`
	//// ETag for the item. Read-only.
	//ETag *string `json:"eTag,omitempty"`
	//// Identity of the user, device, and application which last modified the item. Read-only.
	//LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	//// Date and time the item was last modified. Read-only.
	//LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	//// The name of the item. Read-write.
	//Name *string `json:"name,omitempty"`
	//// Parent information, if the item has a parent. Read-write.
	//ParentReference *ItemReference `json:"parentReference,omitempty"`
	//// URL that displays the resource in the browser. Read-only.
	//WebUrl *string `json:"webUrl,omitempty"`
	//// Identity of the user who created the item. Read-only.
	//CreatedByUser *User `json:"createdByUser,omitempty"`
	//// Identity of the user who last modified the item. Read-only.
	//LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
	//// Audio metadata, if the item is an audio file. Read-only.
	//Audio *Audio `json:"audio,omitempty"`
	//// The content stream, if the item represents a file.
	//Content *string `json:"content,omitempty"`
	//// An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
	//CTag *string `json:"cTag,omitempty"`
	//// Information about the deleted state of the item. Read-only.
	//Deleted *Deleted `json:"deleted,omitempty"`
	//// File metadata, if the item is a file. Read-only.
	//File *File `json:"file,omitempty"`
	//// File system information on client. Read-write.
	//FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	//// Folder metadata, if the item is a folder. Read-only.
	//Folder *Folder `json:"folder,omitempty"`
	//// Image metadata, if the item is an image. Read-only.
	//Image *Image `json:"image,omitempty"`
	//// Location metadata, if the item has location data. Read-only.
	//Location *GeoCoordinates `json:"location,omitempty"`
	//// If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
	//Package *Package `json:"package,omitempty"`
	//// If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
	//PendingOperations *PendingOperations `json:"pendingOperations,omitempty"`
	//// Photo metadata, if the item is a photo. Read-only.
	//Photo *Photo `json:"photo,omitempty"`
	//// Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
	//Publication *PublicationFacet `json:"publication,omitempty"`
	//// Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
	//RemoteItem *RemoteItem `json:"remoteItem,omitempty"`
	//// If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
	//Root NullableAnyOfobject `json:"root,omitempty"`
	//// Search metadata, if the item is from a search result. Read-only.
	//SearchResult *SearchResult `json:"searchResult,omitempty"`
	//// Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
	//Shared *Shared `json:"shared,omitempty"`
	//// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	//SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	//// Size of the item in bytes. Read-only.
	//Size *int64 `json:"size,omitempty"`
	//// If the current item is also available as a special folder, this facet is returned. Read-only.
	//SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`
	//// Video metadata, if the item is a video. Read-only.
	//Video *Video `json:"video,omitempty"`
	//// WebDAV compatible URL for the item.
	//WebDavUrl *string `json:"webDavUrl,omitempty"`
	//// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
	//Workbook *Workbook `json:"workbook,omitempty"`
	//// Analytics about the view activities that took place on this item.
	//Analytics *ItemAnalytics `json:"analytics,omitempty"`
	//// Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
	//Children *[]DriveItem `json:"children,omitempty"`
	//// For drives in SharePoint, the associated document library list item. Read-only. Nullable.
	//ListItem *ListItem `json:"listItem,omitempty"`
	//// The set of permissions for the item. Read-only. Nullable.
	//Permissions *[]Permission `json:"permissions,omitempty"`
	//// The set of subscriptions on the item. Only supported on the root of a drive.
	//Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
	//// Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
	//Thumbnails *[]ThumbnailSet `json:"thumbnails,omitempty"`
	//// The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
	//Versions *[]DriveItemVersion `json:"versions,omitempty"`
}
