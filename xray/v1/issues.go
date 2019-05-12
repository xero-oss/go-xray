package v1

import (
	"context"
	"fmt"
	"net/http"
)

type IssuesService Service

type IssueEventComponent struct {
	Id                 *string   `json:"id,omitempty"`
	VulnerableVersions *[]string `json:"vulnerable_versions,omitempty"`
}

type IssueEventCVE struct {
	CVE    *string `json:"cve,omitempty"`
	CVSSV2 *string `json:"cvss_v2,omitempty"`
}

type IssueEventSource struct {
	SourceId *string `json:"source_id,omitempty"`
}

type CustomIssueEvent struct {
	Type        *string                `json:"type,omitempty"`
	Provider    *string                `json:"provider,omitempty"`
	PackageType *string                `json:"package_type,omitempty"`
	Severity    *string                `json:"severity,omitempty"`
	Components  *[]IssueEventComponent `json:"components,omitempty"`
	CVES        *[]IssueEventCVE       `json:"cves,omitempty"`
	Summary     *string                `json:"summary,omitempty"`
	Description *string                `json:"description,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Sources     *[]IssueEventSource    `json:"sources,omitempty"`
}

type GetIssueEventComponent struct {
	ComponentId *string `json:"component_id,omitempty"`
}

type GetIssueEventOutput struct {
	Provider    *string                   `json:"provider,omitempty"`
	Type        *string                   `json:"type,omitempty"`
	SourceId    *string                   `json:"source_id,omitempty"`
	Url         *string                   `json:"url,omitempty"`
	Created     *string                   `json:"created,omitempty"`
	Modified    *string                   `json:"modified,omitempty"`
	Updated     *string                   `json:"updated,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Summary     *string                   `json:"summary,omitempty"`
	Severity    *string                   `json:"severity,omitempty"`
	Components  *[]GetIssueEventComponent `json:"components,omitempty"`
	Properties  *map[string]string        `json:"properties,omitempty"`
}

// Description: Gets an issue created by a vendor
// Security:  Requires a valid user with the "View Components" permission
// Usage: client.V1.Issues.GetIssue(ctx, "name")
func (s *IssuesService) GetIssue(ctx context.Context, issueId string) (*GetIssueEventOutput, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/events/%s", issueId)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	issue := new(GetIssueEventOutput)
	resp, err := s.client.Do(ctx, req, &issue)
	return issue, resp, err
}

// Description: Creates a custom issue
// Security:  Requires a valid user with "Manage Components" permission
// Usage: client.V1.Issues.CreateIssue(ctx, issue)
func (s *IssuesService) CreateIssue(ctx context.Context, issue *CustomIssueEvent) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/events", issue)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// Description: Updates a custom event
// Security:  Requires a valid user with "Manage Components" permission
// Usage: client.V1.Issues.UpdateIssue(ctx, "issueId", permission)
func (s *IssuesService) UpdateIssue(ctx context.Context, issueId string, issue *CustomIssueEvent) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/events/%s", issueId)
	req, err := s.client.NewJSONEncodedRequest("PUT", path, issue)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
