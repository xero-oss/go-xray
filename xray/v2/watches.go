package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WatchesService Service

type WatchGeneralData struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type WatchFilterValue struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// WatchFilterValueWrapper is a wrapper around WatchFilterValue which handles the API returning both a string and an object for the watch filter value
type WatchFilterValueWrapper struct {
	WatchFilterValue
	IsPropertyFilter bool `json:”-”`
}

type WatchFilter struct {
	Type  *string                  `json:"type,omitempty"`
	Value *WatchFilterValueWrapper `json:"value,omitempty"`
}

type WatchProjectResource struct {
	Type            *string        `json:"type,omitempty"`
	RepoType        *string        `json:"repo_type,omitempty"`
	BinaryManagerId *string        `json:"bin_mgr_id,omitempty"`
	Name            *string        `json:"name,omitempty"`
	Filters         *[]WatchFilter `json:"filters,omitempty"`
}

type WatchProjectResources struct {
	Resources *[]WatchProjectResource `json:"resources,omitempty"`
}

type WatchAssignedPolicy struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type Watch struct {
	GeneralData      *WatchGeneralData      `json:"general_data,omitempty"`
	ProjectResources *WatchProjectResources `json:"project_resources,omitempty"`
	AssignedPolicies *[]WatchAssignedPolicy `json:"assigned_policies,omitempty"`
}

// UnmarshalJSON converts JSON data into a WatchFilterValueWrapper object
// It returns any errors that occured during the function
func (wf *WatchFilterValueWrapper) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		var v WatchFilterValue

		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		wf.WatchFilterValue = v
		wf.IsPropertyFilter = true

		return nil
	}

	wf.Value = &v
	wf.IsPropertyFilter = false

	return nil
}

// MarshalJSON coverts the WatchFilterValueWrapper into JSON data
// It returns the JSON data and any errors that occured during the function
func (wf WatchFilterValueWrapper) MarshalJSON() ([]byte, error) {
	if wf.IsPropertyFilter {
		return json.Marshal(WatchFilterValue{
			Key:   wf.Key,
			Value: wf.Value,
		})
	}

	return json.Marshal(&wf.Value)
}

// Description: Gets a list of all watches in the system
// Security:  Requires a valid user with "View Watches" permission
// Usage: client.V2.Watches.ListWatches(ctx)
func (s *WatchesService) ListWatches(ctx context.Context) (*[]Watch, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v2/watches", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	watches := new([]Watch)
	resp, err := s.client.Do(ctx, req, &watches)
	return watches, resp, err
}

// Description: Gets a named watch
// Security:  Requires a valid user with "View Watches" permission
// Usage: client.V2.Watches.GetWatch(ctx, "name")
func (s *WatchesService) GetWatch(ctx context.Context, name string) (*Watch, *http.Response, error) {
	path := fmt.Sprintf("/api/v2/watches/%s", name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	watch := new(Watch)
	resp, err := s.client.Do(ctx, req, &watch)
	return watch, resp, err
}

// Description: Creates a new Watch
// Security:  Requires a valid user with "Manage Watches" permission
// Usage: client.V2.Watches.CreateWatch(ctx, watch)
func (s *WatchesService) CreateWatch(ctx context.Context, watch *Watch) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v2/watches", watch)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Updates a Watch.
// Security:  Requires a valid user with "Manage Watches" permission
// Usage: client.V2.Watches.UpdateWatch(ctx, "name", watch)
func (s *WatchesService) UpdateWatch(ctx context.Context, name string, watch *Watch) (*http.Response, error) {
	path := fmt.Sprintf("/api/v2/watches/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, watch)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Deletes a Watch
// Security:  Requires a valid user with "Manage Watches" permissions
// Usage: client.V2.Watches.DeleteWatch(ctx, "name")
func (s *WatchesService) DeleteWatch(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v2/watches/%s", name)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
