package models

// SchemaExtension struct for SchemaExtension
type SchemaExtension struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Description for the schema extension.
	Description *string `json:"description,omitempty"`
	// The appId of the application that is the owner of the schema extension. This property can be supplied on creation, to set the owner.  If not supplied, then the calling application's appId will be set as the owner. In either case, the signed-in user must be the owner of the application. So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner property. Once set, this property is read-only and cannot be changed.
	Owner *string `json:"owner,omitempty"`
	// The collection of property names and types that make up the schema extension definition.
	Properties *[]ExtensionSchemaProperty `json:"properties,omitempty"`
	// The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated. Automatically set to InDevelopment on creation. Schema extensions provides more information on the possible state transitions and behaviors.
	Status *string `json:"status,omitempty"`
	// Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from contact, device, event, group, message, organization, post, or user.
	TargetTypes *[]string `json:"targetTypes,omitempty"`
}

// ExtensionSchemaProperty struct for ExtensionSchemaProperty
type ExtensionSchemaProperty struct {
	// The name of the strongly-typed property defined as part of a schema extension.
	Name *string `json:"name,omitempty"`
	// The type of the property that is defined as part of a schema extension.  Allowed values are Binary, Boolean, DateTime, Integer or String.  See the table below for more details.
	Type *string `json:"type,omitempty"`
}

type SchemaExtensionRawDataType map[string]map[string]string
