package v1

import (
	"context"
	"fmt"
	"net/http"
)

type PermissionsService Service

type PermissionReference struct {
	Name *string `json:"name,omitempty"`
	Uri  *string `json:"uri,omitempty"`
}

type PermissionResource struct {
	Name          *string `json:"name,omitempty"`
	ArtifactoryId *string `json:"artifactory_id,omitempty"`
	Type          *string `json:"type,omitempty"`
}

type PermissionUser struct {
	Name  *string   `json:"name,omitempty"`
	Roles *[]string `json:"roles,omitempty"`
}

type PermissionGroup struct {
	Name  *string   `json:"name,omitempty"`
	Roles *[]string `json:"roles,omitempty"`
}

type Permission struct {
	Name      *string               `json:"name,omitempty"`
	Scope     *string               `json:"scope,omitempty"`
	Resources *[]PermissionResource `json:"resources,omitempty"`
	Users     *[]PermissionUser     `json:"users,omitempty"`
	Groups    *[]PermissionGroup    `json:"groups,omitempty"`
}

// Description: Gets a list of all permissions in the system
// Security:  Requires an admin user
// Usage: client.V1.Permissions.ListPermissions(ctx)
func (s *PermissionsService) ListPermissions(ctx context.Context) (*[]PermissionReference, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("GET", "/api/v1/permissions", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	permissions := new([]PermissionReference)
	resp, err := s.client.Do(ctx, req, &permissions)
	return permissions, resp, err
}

// Description: Gets the details of a specific permission
// Security:  Requires an admin user
// Usage: client.V1.Permissions.GetPermission(ctx, "name")
func (s *PermissionsService) GetPermission(ctx context.Context, name string) (*Permission, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/permissions/%s", name)
	req, err := s.client.NewJSONEncodedRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	permission := new(Permission)
	resp, err := s.client.Do(ctx, req, &permission)
	return permission, resp, err
}

// Description: Creates a permission
// Security:  Requires an admin user
// Usage: client.V1.Permissions.CreatePermission(ctx, permission)
func (s *PermissionsService) CreatePermission(ctx context.Context, permission *Permission) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/permissions", permission)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// NOTE: This is an undocumented API and may not work
// Description: Updates a permission
// Security:  Requires an admin user
// Usage: client.V1.Permissions.UpdatePermission(ctx, "name", permission)
func (s *PermissionsService) UpdatePermission(ctx context.Context, name string, permission *Permission) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/permissions/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, permission)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// Description: Deletes a permission
// Security:  Requires an admin user
// Usage: client.V1.Permissions.DeletePermission(ctx, "name")
func (s *PermissionsService) DeletePermission(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/permissions/%s", name)
	req, err := s.client.NewJSONEncodedRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
