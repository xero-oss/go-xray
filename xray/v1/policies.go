package v1

import (
	"context"
	"fmt"
	"net/http"
)

type PoliciesService Service

type PolicyCVSSRange struct {
	To   *string `json:"to,omitempty"`
	From *string `json:"from,omitempty"`
}

type PolicyRuleCriteria struct {
	// Security Criteria
	MinimumSeverity *string          `json:"min_severity,omitempty"`
	CVSSRange       *PolicyCVSSRange `json:"cvss_range,omitempty"`

	// License Criteria
	AllowUnkown     *bool     `json:"allow_unknown,omitempty"`
	BannedLicenses  *[]string `json:"banned_licenses,omitempty"`
	AllowedLicenses *[]string `json:"allowed_licenses,omitempty"`
}

type BlockDownloadSettings struct {
	Unscanned *bool `json:"unscanned,omitempty"`
	Active    *bool `json:"active,omitempty"`
}

type PolicyRuleActions struct {
	Mails          *[]string              `json:"mails,omitempty"`
	FailBuild      *bool                  `json:"fail_build,omitempty"`
	BlockDownload  *BlockDownloadSettings `json:"block_download,omitempty"`
	Webhooks       *[]string              `json:"webhooks,omitempty"`
	CustomSeverity *string                `json:"custom_severity,omitempty"`
}

type PolicyRule struct {
	Name     *string             `json:"name,omitempty"`
	Priority *int                `json:"priority,omitempty"`
	Criteria *PolicyRuleCriteria `json:"criteria,omitempty"`
	Actions  *PolicyRuleActions  `json:"actions,omitempty"`
}

type Policy struct {
	Name        *string       `json:"name,omitempty"`
	Type        *string       `json:"type,omitempty"`
	Author      *string       `json:"author,omitempty"`
	Description *string       `json:"description,omitempty"`
	Rules       *[]PolicyRule `json:"rules,omitempty"`
	Created     *string       `json:"created,omitempty"`
	Modified    *string       `json:"modified,omitempty"`
}

// Description:  Gets a list of all policies in the system
// Security:  Requires a user with "View Watches" permission
// Usage: client.V1.Policies.ListPolicies(ctx)
func (s *PoliciesService) ListPolicies(ctx context.Context) (*[]Policy, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/policies", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	policies := new([]Policy)
	resp, err := s.client.Do(ctx, req, &policies)
	return policies, resp, err
}

// Description:  Gets a specific policy
// Security:  Requires a user with "View Watches" permission
// Usage: client.V1.Policies.GetPolicy(ctx, "name")
func (s *PoliciesService) GetPolicy(ctx context.Context, name string) (*Policy, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/policies/%s", name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	policy := new(Policy)
	resp, err := s.client.Do(ctx, req, &policy)
	return policy, resp, err
}

// Description: Creates a new policy
// Security:  Requires a user with Manage Policies permission
// Usage: client.V1.Users.CreateUser(ctx, user)
func (s *PoliciesService) CreatePolicy(ctx context.Context, policy *Policy) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/policies", policy)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Updates a policy.
// Security:  Requires a user with Manage Policies permission
// Usage: client.V1.Policies.UpdatePolicy(ctx, "name", user)
func (s *PoliciesService) UpdatePolicy(ctx context.Context, name string, policy *Policy) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/policies/%s", name)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, policy)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Description: Deletes a Policy
// Security:  Requires a user with Manage Policies permission
// Usage: client.V1.Policies.DeletePolicy(ctx, "name")
func (s *PoliciesService) DeletePolicy(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/policies/%s", name)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

type AssignPolicyInput struct {
	Watches *[]string `json:"watches,omitempty"`
}

//  Description: Assign a policy to watches
//  Security: Requires a valid user with "Manage Watches" permission
//  Usage: client.V1.Policies.AssignPolicy(ctx, "name", assignPolicyInput)
func (s *PoliciesService) AssignPolicy(ctx context.Context, name string, assignPolicyInput *AssignPolicyInput) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/policies/%s/assign", name)
	req, err := s.client.NewJSONEncodedRequest("POST", path, assignPolicyInput)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
