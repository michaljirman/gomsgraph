package v1

import (
	"context"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	"github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

type SchemaExtensionsService service

const (
	SchemaExtensionsResourceName = "schemaExtensions"
)

type SchemaExtensionListOptions struct {
	// The $count and $search parameters are currently not available in Azure AD B2C tenants.
	//Search string `url:"$search,omitempty"`
	//Count  bool   `url:"$count,omitempty"`
	Filter   string `url:"$filter,omitempty"`
	Select   string `url:"$select,omitempty"`
	Top      int    `url:"$top,omitempty"`
	NextLink string `url:"-"`
}

type SchemaExtensionsResponse struct {
	models.OData
	SchemaExtensions []*models.SchemaExtension `json:"value"`
}

// Get a list of schemaExtension objects created by any apps you own in the current tenant
// (that can be InDevelopment, Available, or Deprecated), and all other schema extensions owned by other apps that are marked as Available.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-list?view=graph-rest-1.0&tabs=http
func (s *SchemaExtensionsService) ListAll(ctx context.Context, opts *SchemaExtensionListOptions) (*SchemaExtensionsResponse, error) {
	var u string
	if opts.NextLink != "" {
		u = opts.NextLink
	} else {
		var err error
		u, err = core.AddOptions(SchemaExtensionsResourceName, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var schemaExtensionsResp SchemaExtensionsResponse
	_, err = s.client.Do(ctx, req, &schemaExtensionsResp)
	if err != nil {
		return nil, err
	}

	return &schemaExtensionsResp, nil
}

type SchemaExtensionResponse struct {
	models.OData
	models.SchemaExtension
}

// Get the properties of the specified schemaExtension definition.
//
// GET /schemaExtensions/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-get?view=graph-rest-1.0&tabs=http
func (s *SchemaExtensionsService) GetSchemaExtension(ctx context.Context, schemaExtensionID string, opts *SchemaExtensionListOptions) (*SchemaExtensionResponse, error) {
	u, err := core.AddOptions(core.BuildURL(SchemaExtensionsResourceName, schemaExtensionID), opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	schemaExtensionResp := new(SchemaExtensionResponse)
	_, err = s.client.Do(ctx, req, schemaExtensionResp)
	if err != nil {
		return nil, err
	}

	return schemaExtensionResp, nil
}

// Create a new schemaExtension definition to extend a supporting resource type.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-post-schemaextensions?view=graph-rest-1.0&tabs=http
func (s *SchemaExtensionsService) CreateSchemaExtension(ctx context.Context, r *models.SchemaExtension) (*SchemaExtensionResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, SchemaExtensionsResourceName, r)
	if err != nil {
		return nil, err
	}

	schemaExtensionResp := new(SchemaExtensionResponse)
	_, err = s.client.Do(ctx, req, schemaExtensionResp)
	if err != nil {
		return nil, err
	}

	return schemaExtensionResp, nil
}

// Delete the definition of a schema extension.
// Only the app that created the schema extension (owner app) can delete the schema extension definition,
// and only when the extension is in the InDevelopment state. Deleting a schema extension definition does not affect
// accessing custom data that has been added to resource instances based on that definition.
//
//
// DELETE /schemaExtensions/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-delete?view=graph-rest-1.0&tabs=http
func (s *SchemaExtensionsService) DeleteSchemaExtension(ctx context.Context, schemaExtensionID string) error {
	req, err := s.client.NewRequest(http.MethodDelete, core.BuildURL(SchemaExtensionsResourceName, schemaExtensionID), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Update properties in the definition of the specified schemaExtension.
// This means custom properties or target resource types cannot be removed from the definition,
// but new custom properties can be added and the description of the extension changed.
// Additive updates to the extension can only be made when the extension is in the InDevelopment or Available status.
//
// PATCH /schemaExtensions/{id}
//
// Warning:
// * One application can only create up to 5 schema extensions. Keep that in mind!
// * Extension schema can be deleted only if the status is not set to "Available". (InDevelopment=>Available=>Deprecated)
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-update?view=graph-rest-1.0&tabs=http
func (s *SchemaExtensionsService) UpdateSchemaExtension(ctx context.Context, r *models.SchemaExtension) error {
	req, err := s.client.NewRequest(http.MethodPatch, core.BuildURL(SchemaExtensionsResourceName, *r.Id), r)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
