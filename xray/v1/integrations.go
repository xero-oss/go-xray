package v1

import (
	"context"
	"fmt"
	"net/http"
)

type IntegrationsService Service

type Integration struct {
	Vendor      *string `json:"vendor,omitempty"`
	ApiKey      *string `json:"api_key,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Context     *string `json:"context,omitempty"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	TestUrl     *string `json:"test_url,omitempty"`
}

// Description: Gets a list of all integrations in the system
// Security:  Requires an admin user
// Usage: client.V1.Integrations.ListIntegrations(ctx)
func (s *IntegrationsService) ListIntegrations(ctx context.Context) (*[]Integration, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/integration", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	integrations := new([]Integration)
	resp, err := s.client.Do(ctx, req, &integrations)
	return integrations, resp, err
}

// Description: Creates a new Integration configuration
// Security:  Requires an admin user
// Usage: client.V1.Integrations.CreateIntegration(ctx, integration)
func (s *IntegrationsService) CreateIntegration(ctx context.Context, integration *Integration) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/integration", integration)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Updates an integration.
// Security:  Requires an admin user
// Usage: client.V1.Integrations.UpdateIntegration(ctx, "name", integration)
func (s *IntegrationsService) UpdateIntegration(ctx context.Context, name string, integration *Integration) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/integration/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, integration)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Deletes an Integration
// Security:  Requires an admin users
// Usage: client.V1.Integrations.DeleteIntegration(ctx, "name")
func (s *IntegrationsService) DeleteIntegration(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/integration/%s", name)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
