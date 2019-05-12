package v1

import (
	"context"
	"fmt"
	"net/http"
)

type BinaryManagersService Service

type BinaryManager struct {
	User         *string `json:"user,omitempty"`
	Password     *string `json:"password,omitempty"`
	Url          *string `json:"binMgrUrl,omitempty"`
	Id           *string `json:"binMgrId,omitempty"`
	Description  *string `json:"binMgrDesc,omitempty"`
	ProxyEnabled *bool   `json:"proxy_enabled,omitempty"`
}

type BinaryManagerRepository struct {
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	PackageType *string `json:"pkg_type,omitempty"`
}

type BinaryManagerRepoIndexingConfiguration struct {
	IndexedRepos    *[]BinaryManagerRepository `json:"indexed_repos,omitempty"`
	NonIndexedRepos *[]BinaryManagerRepository `json:"non_indexed_repos,omitempty"`
}

type BinaryManagerBuildIndexingConfiguration struct {
	IndexedBuilds    *[]string `json:"indexed_builds,omitempty"`
	NonIndexedBuilds *[]string `json:"non_indexed_builds,omitempty"`
}

// Description: Gets a list of all binary managers in the system
// Security:  Requires a valid user
// Usage: client.V1.BinaryManagers.ListBinaryManagers(ctx)
func (s *BinaryManagersService) ListBinaryManagers(ctx context.Context) (*[]BinaryManager, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/binMgr", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	binMgrs := new([]BinaryManager)
	resp, err := s.client.Do(ctx, req, &binMgrs)
	return binMgrs, resp, err
}

// Description: Gets a named binary manager
// Security:  Requires a valid user
// Usage: client.V1.BinaryManagers.GetBinaryManager(ctx, "name")
func (s *BinaryManagersService) GetBinaryManager(ctx context.Context, name string) (*BinaryManager, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/binMgr/%s", name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	binMgr := new(BinaryManager)
	resp, err := s.client.Do(ctx, req, &binMgr)
	return binMgr, resp, err
}

// Description: Creates a new Binary Manager
// Security:  Requires an admin user
// Usage: client.V1.BinaryManagers.CreateBinaryManager(ctx, binMgr)
func (s *BinaryManagersService) CreateBinaryManager(ctx context.Context, binMgr *BinaryManager) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/binMgr", binMgr)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Updates a binary manager.
// Security:  Requires an admin user
// Usage: client.V1.BinaryManagers.UpdateBinaryManager(ctx, "name", binMgr)
func (s *BinaryManagersService) UpdateBinaryManager(ctx context.Context, name string, binMgr *BinaryManager) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/binMgr/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, binMgr)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Deletes a Binary Manager
// Security:  Requires an admin users
// Usage: client.V1.BinaryManagers.DeleteBinaryManager(ctx, "name")
func (s *BinaryManagersService) DeleteBinaryManager(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/binMgr/%s", name)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
