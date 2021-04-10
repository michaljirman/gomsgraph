package models

type Quota struct {
	// Total space consumed by files in the recycle bin, in bytes. Read-only.
	Deleted *int64 `json:"deleted,omitempty"`
	// Total space remaining before reaching the quota limit, in bytes. Read-only.
	Remaining *int64 `json:"remaining,omitempty"`
	// Enumeration value that indicates the state of the storage space. Read-only.
	State *string `json:"state,omitempty"`
	// Information about the drive's storage quota plans. Only in Personal OneDrive.
	StoragePlanInformation *StoragePlanInformation `json:"storagePlanInformation,omitempty"`
	// Total allowed storage space, in bytes. Read-only.
	Total *int64 `json:"total,omitempty"`
	// Total space used, in bytes. Read-only.
	Used *int64 `json:"used,omitempty"`
}

type StoragePlanInformation struct {
	// Indicates whether there are higher storage quota plans available. Read-only.
	UpgradeAvailable *bool `json:"upgradeAvailable,omitempty"`
}
