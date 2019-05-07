package v1

import (
	"context"
	"net/http"
)

type ViolationsService Service

type GetViolationsFilters struct {
	NameContains    *string `json:"name_contains,omitempty"`
	ViolationType   *string `json:"violation_type,omitempty"`
	WatchName       *string `json:"name_contains,omitempty"`
	MinimumSeverity *string `json:"min_severity,omitempty"`
	CreatedFrom     *string `json:"created_from,omitempty"`
}

type GetViolationsPagination struct {
	OrderBy *string `json:"order_by,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Offset  *string `json:"offset,omitempty"`
}

type GetViolationsInput struct {
	Filters    *GetViolationsFilters    `json:"filters,omitempty"`
	Pagination *GetViolationsPagination `json:"pagination,omitempty"`
}

type Violation struct {
	Description         *string   `json:"description,omitempty"`
	Severity            *string   `json:"severity,omitempty"`
	Type                *string   `json:"type,omitempty"`
	InfectedComponent   *[]string `json:"infected_component,omitempty"`
	Created             *string   `json:"created,omitempty"`
	WatchName           *string   `json:"watch_name,omitempty"`
	IssueId             *string   `json:"issue_id,omitempty"`
	ViolationDetailsUrl *string   `json:"violations_details_url,omitempty"`
	ImpactedArtifacts   *[]string `json:"impacted_artifacts,omitempty"`
}

type GetViolationsOutput struct {
	TotalViolations *int         `json:"total_violations,omitempty"`
	Violations      *[]Violation `json:"violations,omitempty"`
}

// Description:  Gets the Xray violations based on a set of search criteria
// Security:  Requires a "View Watches" permission
// Usage: client.V1.Violations.GetViolations(ctx, getViolationsInput)
func (s *ViolationsService) GetViolations(ctx context.Context, getViolationsInput *GetViolationsInput) (*GetViolationsOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/violations", getViolationsInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetViolationsOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}
