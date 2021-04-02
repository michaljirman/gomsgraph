package models

// Onenote struct for Onenote
type Onenote struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	//Notebooks *[]Notebook `json:"notebooks,omitempty"`
	//// The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
	//Operations *[]OnenoteOperation `json:"operations,omitempty"`
	//// The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	//Pages *[]OnenotePage `json:"pages,omitempty"`
	//// The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable.
	//Resources *[]OnenoteResource `json:"resources,omitempty"`
	//// The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	//SectionGroups *[]SectionGroup `json:"sectionGroups,omitempty"`
	//// The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	//Sections *[]OnenoteSection `json:"sections,omitempty"`
}
