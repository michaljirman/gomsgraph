package models

// OfficeGraphInsights struct for OfficeGraphInsights
type OfficeGraphInsights struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
	Shared *[]SharedInsight `json:"shared,omitempty"`
	// Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user's closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before.
	Trending *[]Trending `json:"trending,omitempty"`
	// Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use.
	Used *[]UsedInsight `json:"used,omitempty"`
}

// SharedInsight struct for SharedInsight
type SharedInsight struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Details about the shared item. Read only.
	//LastShared *SharingDetail `json:"lastShared,omitempty"`
	//// Reference properties of the shared document, such as the url and type of the document. Read-only
	//ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	//// Properties that you can use to visualize the document in your experience. Read-only
	//ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	//SharingHistory        *[]AnyOfSharingDetail  `json:"sharingHistory,omitempty"`
	//LastSharedMethod      *Entity                `json:"lastSharedMethod,omitempty"`
	//// Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
	//Resource *Entity `json:"resource,omitempty"`
}

// Trending struct for Trending
type Trending struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	//// Reference properties of the trending document, such as the url and type of the document.
	//ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	//// Properties that you can use to visualize the document in your experience.
	//ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	//// Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
	//Weight *AnyOfnumberstringstring `json:"weight,omitempty"`
	//// Used for navigating to the trending document.
	//Resource *Entity `json:"resource,omitempty"`
}

// UsedInsight struct for UsedInsight
type UsedInsight struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Information about when the item was last viewed or modified by the user. Read only.
	//LastUsed *UsageDetails `json:"lastUsed,omitempty"`
	//// Reference properties of the used document, such as the url and type of the document. Read-only
	//ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	//// Properties that you can use to visualize the document in your experience. Read-only
	//ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	//// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
	//Resource *Entity `json:"resource,omitempty"`
}
