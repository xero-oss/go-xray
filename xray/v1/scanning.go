package v1

import (
	"context"
	"net/http"
)

type ScanningService Service

type ScanArtifactInput struct {
	ComponentId *string `json:"componentId,omitempty"`
}

type ScanArtifactOutput struct {
	Info *string `json:"info,omitempty"`
}

type ScanBuildInput struct {
	ArtifactoryId *string `json:"artifactoryId,omitempty"`
	BuildName     *string `json:"buildName,omitempty"`
	BuildNumber   *string `json:"buildNumber,omitempty"`
}

type BuildScanSummary struct {
	FailBuild      *bool   `json:"fail_build,omitempty"`
	Message        *string `json:"message,omitempty"`
	MoreDetailsUrl *string `json:"more_details_url,omitempty"`
	TotalAlerts    *int    `json:"total_alerts,omitempty"`
}

type BuildScanInfectedFileBannedLicense struct {
	AlertType   *string `json:"alert_type,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Severity    *string `json:"severity,omitempty"`
	Summary     *string `json:"summary,omitempty"`
}

type BuildScanInfectedFileVulnerability struct {
	AlertType   *string `json:"alert_type,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Severity    *string `json:"severity,omitempty"`
	Summary     *string `json:"summary,omitempty"`
}

type BuildScanInfectedFileDetails struct {
	BannedLicenses  *[]BuildScanInfectedFileBannedLicense `json:"banned_licenses,omitempty"`
	Child           *string                               `json:"child,omitempty"`
	Vulnerabilities *[]BuildScanInfectedFileVulnerability `json:"vulnerabilities,omitempty"`
}

type BuildScanInfectedFile struct {
	ComponentId *string                       `json:"component_id,omitempty"`
	Depth       *string                       `json:"depth,omitempty"`
	Details     *BuildScanInfectedFileDetails `json:"details,omitempty"`
	DisplayName *string                       `json:"display_name,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	ParentSHA   *string                       `json:"parent_sha,omitempty"`
	Path        *string                       `json:"path,omitempty"`
	PackageType *string                       `json:"pkg_type,omitempty"`
	SHA1        *string                       `json:"sha1,omitempty"`
	SHA256      *string                       `json:"sha256,omitempty"`
}

type BuildScanArtifact struct {
	Depth         *string                  `json:"depth,omitempty"`
	DisplayName   *string                  `json:"display_name,omitempty"`
	InfectedFiles *[]BuildScanInfectedFile `json:"infected_files,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	ParentSHA     *string                  `json:"parent_sha,omitempty"`
	Path          *string                  `json:"path,omitempty"`
	PackageType   *string                  `json:"pkg_type,omitempty"`
	SHA1          *string                  `json:"sha1,omitempty"`
	SHA256        *string                  `json:"sha256,omitempty"`
}

type BuildScanIssue struct {
	Created           *string              `json:"created,omitempty"`
	CVE               *string              `json:"cve,omitempty"`
	Description       *string              `json:"description,omitempty"`
	ImpactedArtifacts *[]BuildScanArtifact `json:"impacted_artifacts,omitempty"`
	Provider          *string              `json:"provider,omitempty"`
	Severity          *string              `json:"severity,omitempty"`
	Summary           *string              `json:"summary,omitempty"`
	Type              *string              `json:"type,omitempty"`
}

type BuileScanAlert struct {
	Created     *string           `json:"created,omitempty"`
	Issues      *[]BuildScanIssue `json:"issues,omitempty"`
	TopSeverity *string           `json:"top_severity,omitempty"`
	WatchName   *string           `json:"watch_name,omitempty"`
	SHA1        *string           `json:"sha1,omitempty"`
}

type BuildScanLicense struct {
	Name        *string   `json:"name,omitempty"`
	Components  *[]string `json:"components,omitempty"`
	FullName    *string   `json:"full_name,omitempty"`
	MoreInfoURL *[]string `json:"more_info_url,omitempty"`
}

type ScanBuildOutput struct {
	Summary  *BuildScanSummary   `json:"summary,omitempty"`
	Alerts   *[]BuileScanAlert   `json:"alerts,omitempty"`
	Licenses *[]BuildScanLicense `json:"path,omitempty"`
}

// Description:  Invokes scanning of an artifact
// Security:  Requires a valid user with "Manage Components" permission
// Usage: client.V1.Scanning.ScanArtifact(ctx, scanArtifactInput)
func (s *ScanningService) ScanArtifact(ctx context.Context, scanArtifactInput *ScanArtifactInput) (*ScanArtifactOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/scanArtifact", scanArtifactInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(ScanArtifactOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

// Description:  Invokes scanning of a build that was uploaded to Artifactory as requested by a CI server
// Security:  Requires a valid user with "Manage Components" permission
// Usage: client.V1.Scanning.ScanBuild(ctx, scanBuildInput)
func (s *ScanningService) ScanBuild(ctx context.Context, scanBuildInput *ScanBuildInput) (*ScanBuildOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/scanBuild", scanBuildInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(ScanBuildOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}
