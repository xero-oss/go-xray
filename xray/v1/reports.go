package v1

import (
	"context"
	"fmt"
	"net/http"
)

type ReportsService Service

type GenerateReportOutput struct {
	Info *string `json:"info,omitempty"`
}

// Description: Generates a new license report
// Security:  Requires a valid user with "Generate Reports" permission
// Usage: client.V1.Reports.GenerateLicenseReport(ctx)
func (s *ReportsService) GenerateLicenseReport(ctx context.Context) (*GenerateReportOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/licensesReport/generate", nil)
	if err != nil {
		return nil, nil, err
	}

	output := new(GenerateReportOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type LicenseCompliance struct {
	Banned  *int `json:"banned,omitempty"`
	Unknown *int `json:"unknown,omitempty"`
	Valid   *int `json:"valid,omitempty"`
}

type LicenseReport struct {
	Distribution *map[string]int    `json:"distribution,omitempty"`
	Compliance   *LicenseCompliance `json:"compliance,omitempty"`
	LastUpdate   *string            `json:"lastUpdate,omitempty"`
}

// Description: Gets the last generated license report
// Security:  Requires a valid user with "Generate Reports" permission
// Usage: client.V1.Reports.GetLicenseReport(ctx)
func (s *ReportsService) GetLicenseReport(ctx context.Context) (*LicenseReport, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/licensesReport", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	report := new(LicenseReport)
	resp, err := s.client.Do(ctx, req, &report)
	return report, resp, err
}

type GetLicenseReportComponentsInput struct {
	License      *string `json:"license,omitempty"`
	Compliance   *string `json:"compliance,omitempty"`
	NumberOfRows *int    `json:"num_of_rows,omitempty"`
	OrderBy      *string `json:"order_by,omitempty"`
	PageNumber   *int    `json:"page_num,omitempty"`
}

func (s *GetLicenseReportComponentsInput) toQueryString() (string, error) {
	var queryString string
	if s.Compliance != nil {
		queryString = fmt.Sprintf("compliance=%s", s.Compliance)
	} else {
		if s.License != nil {
			queryString = fmt.Sprintf("license=%s", s.License)
		}
	}
	// An initial filter is required
	if queryString == "" {
		return "", fmt.Errorf("Either License or Compliance must be set")
	}

	if s.NumberOfRows != nil {
		queryString = fmt.Sprintf("%s&num_of_rows=%s", queryString, s.NumberOfRows)
	}
	if s.OrderBy != nil {
		queryString = fmt.Sprintf("%s&order_by=%s", queryString, s.OrderBy)
	}
	if s.PageNumber != nil {
		queryString = fmt.Sprintf("%s&page_num=%s", queryString, s.PageNumber)
	}

	return queryString, nil
}

type LicenseReportComponent struct {
	Id          *string   `json:"component_id,omitempty"`
	Name        *string   `json:"component_name,omitempty"`
	PackageType *string   `json:"pkg_type,omitempty"`
	IsRoot      *bool     `json:"is_root,omitempty"`
	Licenses    *[]string `json:"licenses,omitempty"`
}

type GetLicenseReportComponentsOutput struct {
	Data       *[]LicenseReportComponent `json:"data,omitempty"`
	TotalCount *int                      `json:"total_count,omitempty"`
}

// Description: Get license report.  either `License` or `Compliance` query parameter are required together with NumberOfRows, OrderBy and PageNumber.
// Security:  Requires a "Generate Reports" permission
// Usage: client.V1.Reports.GetLicenseReportComponents(ctx, getViolationsInput)
func (s *ReportsService) GetLicenseReportComponents(ctx context.Context, getLicenseReportComponentsInput *GetLicenseReportComponentsInput) (*GetLicenseReportComponentsOutput, *http.Response, error) {
	query, err := getLicenseReportComponentsInput.toQueryString()
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("/api/v1/licensesReport/components?%s", query)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetLicenseReportComponentsOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

// Description: Generates a new security report
// Security:  Requires a valid user with "Generate Reports" permission
// Usage: client.V1.Reports.GenerateSecurityReport(ctx)
func (s *ReportsService) GenerateSecurityReport(ctx context.Context) (*GenerateReportOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/securityReport/generate", nil)
	if err != nil {
		return nil, nil, err
	}

	output := new(GenerateReportOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type SecurityReportTopVulnerability struct {
	Summary                *string `json:"summary,omitempty"`
	TotalAffectedArtifacts *int    `json:"total_affected_artifacts,omitempty"`
}

type SecurityReportTopArtifact struct {
	ComponentId          *string `json:"component_id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Version              *string `json:"version,omitempty"`
	PackageType          *string `json:"package_type,omitempty"`
	VulnerabilitiesCount *int    `json:"vulnerabilities_count,omitempty"`
}

type SecurityReport struct {
	RecentVulnerabilities *map[string]int                   `json:"recent_vulnerabilities,omitempty"`
	RecentComponents      *map[string]int                   `json:"recent_components,omitempty"`
	TopVulnerabilities    *[]SecurityReportTopVulnerability `json:"top_vulnerabilities,omitempty"`
	TopArtifacts          *[]SecurityReportTopArtifact      `json:"top_artifacts,omitempty"`
	LastUpdate            *string                           `json:"lastUpdate,omitempty"`
}

// Description: Gets the last generated security report
// Security:  Requires a valid user with "Generate Reports" permission
// Usage: client.V1.Reports.GetSecurityReport(ctx)
func (s *ReportsService) GetSecurityReport(ctx context.Context) (*SecurityReport, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/securityReport", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	report := new(SecurityReport)
	resp, err := s.client.Do(ctx, req, &report)
	return report, resp, err
}

type TopVulnerabilityReportCve struct {
	Cve  *string `json:"cve,omitempty"`
	Cvss *string `json:"cvss,omitempty"`
}

type TopVulnerabilityReportAffectedComponent struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Version     *string `json:"version,omitempty"`
	PackageType *string `json:"package_type,omitempty"`
}

type TopVulnerabilityReport struct {
	Summary            *string                                    `json:"summary,omitempty"`
	Description        *string                                    `json:"description,omitempty"`
	Severity           *string                                    `json:"severity,omitempty"`
	Created            *string                                    `json:"created,omitempty"`
	Cves               *[]TopVulnerabilityReportCve               `json:"cves,omitempty"`
	AffectedComponents *[]TopVulnerabilityReportAffectedComponent `json:"affected_components,omitempty"`
}

// Description: Gets the last generated top vulnerabilities security report
// Security:  Requires a valid user with "Generate Reports" permission
// Usage: client.V1.Reports.GetTopVulnerabilitiesSecurityReport(ctx)
func (s *ReportsService) GetTopVulnerabilitiesSecurityReport(ctx context.Context) (*[]TopVulnerabilityReport, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/securityReport/topVulnerabilities", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	report := new([]TopVulnerabilityReport)
	resp, err := s.client.Do(ctx, req, &report)
	return report, resp, err
}
