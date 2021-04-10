package beta

import (
	"context"
	"errors"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

type SchemaExtensionsService service

type SchemaExtensionsResponse struct {
	OData
	SchemaExtensions []*SchemaExtension `json:"value"`
}

// ListAll retrieves a list of schemaExtension objects created by any apps you own in the current tenant
// (that can be InDevelopment, Available, or Deprecated), and all other schema extensions owned by other apps that are marked as Available.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-list?view=graph-rest-beta&tabs=http
func (s *SchemaExtensionsService) ListAll(ctx context.Context, opts *ListOptions) (*SchemaExtensionsResponse, error) {
	var u string
	if opts.NextLink != "" {
		u = opts.NextLink
	} else {
		u = msgraph.URL(SchemaExtensions).Options(opts).Build()
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
	OData
	SchemaExtension
}

// GetSchemaExtension retrieves the properties of the specified schemaExtension definition.
//
// GET /schemaExtensions/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-get?view=graph-rest-beta&tabs=http
func (s *SchemaExtensionsService) GetSchemaExtension(ctx context.Context, schemaExtensionID string, opts *ListOptions) (*SchemaExtensionResponse, error) {
	u := msgraph.URL(SchemaExtensions, schemaExtensionID).Options(opts).Build()
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

// CreateSchemaExtension creates a new schemaExtension definition to extend a supporting resource type.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-post-schemaextensions?view=graph-rest-beta&tabs=http
func (s *SchemaExtensionsService) CreateSchemaExtension(ctx context.Context, r SchemaExtension) (*SchemaExtensionResponse, error) {
	u := msgraph.URL(SchemaExtensions).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, r)
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

// DeleteSchemaExtension deletes the definition of a schema extension.
// Only the app that created the schema extension (owner app) can delete the schema extension definition,
// and only when the extension is in the InDevelopment state. Deleting a schema extension definition does not affect
// accessing custom data that has been added to resource instances based on that definition.
//
//
// DELETE /schemaExtensions/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/schemaextension-delete?view=graph-rest-beta&tabs=http
func (s *SchemaExtensionsService) DeleteSchemaExtension(ctx context.Context, schemaExtensionID string) error {
	u := msgraph.URL(SchemaExtensions, schemaExtensionID).Build()
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSchemaExtension updates properties in the definition of the specified schemaExtension.
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
// https://docs.microsoft.com/en-us/graph/api/schemaextension-update?view=graph-rest-beta&tabs=http
func (s *SchemaExtensionsService) UpdateSchemaExtension(ctx context.Context, r SchemaExtension) error {
	if r.Id == nil {
		return errors.New("schema extension id is required")
	}
	u := msgraph.URL(SchemaExtensions, *r.Id).Build()
	req, err := s.client.NewRequest(http.MethodPatch, u, r)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
