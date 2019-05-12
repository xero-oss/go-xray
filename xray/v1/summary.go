package v1

import (
	"context"
	"fmt"
	"net/http"
)

type SummaryService Service

type SummaryArtifactGeneral struct {
	ComponentId *string `json:"component_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	PackageType *string `json:"pkg_type,omitempty"`
	Sha256      *string `json:"sha256,omitempty"`
}

type SummaryArtifactIssue struct {
	Created     *string   `json:"created,omitempty"`
	Description *string   `json:"description,omitempty"`
	ImpactPath  *[]string `json:"impact_path,omitempty"`
	IssueType   *string   `json:"issue_type,omitempty"`
	Provider    *string   `json:"provider,omitempty"`
	Severity    *string   `json:"severity,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
}

type SummaryArtifactLicense struct {
	Components  *[]string `json:"components,omitempty"`
	FullName    *string   `json:"full_name,omitempty"`
	MoreInfoUrl *[]string `json:"more_info_url,omitempty"`
	Name        *string   `json:"name,omitempty"`
}

type SummaryArtifact struct {
	General  *SummaryArtifactGeneral   `json:"general,omitempty"`
	Issues   *[]SummaryArtifactIssue   `json:"issues,omitempty"`
	Licenses *[]SummaryArtifactLicense `json:"licenses,omitempty"`
}

type SummaryError struct {
	Error      *string `json:"error,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

type Summary struct {
	Artifacts *[]SummaryArtifact `json:"artifacts,omitempty"`
	Errors    *[]SummaryError    `json:"errors,omitempty"`
}

// Description: Provides details about any build specified by build identifier (name + number)
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Summary.GetBuildSummary(ctx, buildName, buildNumber)
func (s *SummaryService) GetBuildSummary(ctx context.Context, buildName string, buildNumber string) (*Summary, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/summary/build?build_name=%s&build_number=%s", buildName, buildNumber)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	summary := new(Summary)
	resp, err := s.client.Do(ctx, req, &summary)
	return summary, resp, err
}

type GetArtifactSummaryInput struct {
	Checksums *[]string `json:"checksums,omitempty"`
	Paths     *[]string `json:"paths,omitempty"`
}

// Description: Provides details about any artifact specified by path identifiers or checksum
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Summary.GetArtifactSummary(ctx, getArtifactSummaryInput)
func (s *SummaryService) GetArtifactSummary(ctx context.Context, getArtifactSummaryInput *GetArtifactSummaryInput) (*Summary, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/summary/artifact", getArtifactSummaryInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	summary := new(Summary)
	resp, err := s.client.Do(ctx, req, &summary)
	return summary, resp, err
}
