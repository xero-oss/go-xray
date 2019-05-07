package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type ComponentsService Service

type ComponentSource struct {
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	Updated *string `json:"updated,omitempty"`
}

type ComponentVersionFile struct {
	Name   *string `json:"name,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	Sha1   *string `json:"sha1,omitempty"`
	Md5    *string `json:"md5,omitempty"`
}

type ComponentVersion struct {
	Version  *string                 `json:"version,omitempty"`
	Released *string                 `json:"released,omitempty"`
	Licenses *[]string               `json:"licenses,omitempty"`
	Files    *[]ComponentVersionFile `json:"files,omitempty"`
}
type Component struct {
	Component   *string             `json:"component,omitempty"`
	PackageType *string             `json:"package_type,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Description *string             `json:"description,omitempty"`
	WebsiteUrl  *string             `json:"website_url,omitempty"`
	Downloads   *int                `json:"downloads,omitempty"`
	Created     *string             `json:"created,omitempty"`
	Modified    *string             `json:"modified,omitempty"`
	Sources     *[]ComponentSource  `json:"sources,omitempty"`
	Versions    *[]ComponentVersion `json:"versions,omitempty"`
}

// Description: Gets a named component
// Security:  Requires a valid user with "View Component" permission
// Usage: client.V1.Components.GetComponent(ctx, "name")
func (s *ComponentsService) GetComponent(ctx context.Context, name string) (*Component, *http.Response, error) {
	path := fmt.Sprintf("/api/v1/component/%s", name)
	req, err := s.client.NewJSONEncodedRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	component := new(Component)
	resp, err := s.client.Do(ctx, req, &component)
	return component, resp, err
}

type ListComponentByCveInput struct {
	Cves *[]string `json:"cves,omitempty"`
}

type ListComponent struct {
	Name        *string `json:"name,omitempty"`
	PackageType *string `json:"package_type,omitempty"`
	Version     *string `json:"version,omitempty"`
	Link        *string `json:"link,omitempty"`
}

type ListComponentByCveOutput struct {
	CveDetails *string          `json:"cve_details,omitempty"`
	Components *[]ListComponent `json:"components,omitempty"`
}

// Description: Gets a list of components affected by a set of CVEs
// Security:  Requires a valid user with "View Component" permission
// Usage: client.V1.Components.ListComponentsByCves(ctx, listComponentByCveInput)
func (s *ComponentsService) ListComponentsByCves(ctx context.Context, listComponentByCveInput *ListComponentByCveInput) (*[]ListComponentByCveOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/component/searchByCves", listComponentByCveInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new([]ListComponentByCveOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type ListCvesByComponentInput struct {
	ComponentsId *[]string `json:"components_id,omitempty"`
}

type ListCvesByComponentOutput struct {
	Component *string   `json:"component,omitempty"`
	Cves      *[]string `json:"cves,omitempty"`
}

// Description: Gets a list of Cves that affect a particular component
// Security:  Requires a valid user with "View Component" permission
// Usage: client.V1.Components.ListCveSByComponents(ctx, listCvesByComponentInput)
func (s *ComponentsService) ListCvesByComponents(ctx context.Context, listCveSByComponentInput *ListCvesByComponentInput) (*[]ListCvesByComponentOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/component/searchCvesByComponents", listCveSByComponentInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new([]ListCvesByComponentOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type Artifact struct {
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	PackageType *string `json:"pkg_type,omitempty"`
	Sha256      *string `json:"sha256,omitempty"`
	Sha1        *string `json:"sha1,omitempty"`
	ComponentId *string `json:"component_id,omitempty"`
}

type GraphComponent struct {
	ComponentName *string           `json:"component_name,omitempty"`
	ComponentId   *string           `json:"component_id,omitempty"`
	PackageType   *string           `json:"package_type,omitempty"`
	Version       *string           `json:"version,omitempty"`
	Created       *string           `json:"created,omitempty"`
	Modified      *string           `json:"modified,omitempty"`
	Components    *[]GraphComponent `json:"components,omitempty"`
}

type GetArtifactDependencyGraphInput struct {
	Path *string `json:"path,omitempty"`
}

type GetArtifactDependencyGraphOutput struct {
	Artifact   *Artifact         `json:"artifact,omitempty"`
	Components *[]GraphComponent `json:"components,omitempty"`
}

// Description: Gets a dependency graph for a given component
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Components.GetArtifactDependencyGraph(ctx, getArtifactDependencyGraphInput)
func (s *ComponentsService) GetArtifactDependencyGraph(ctx context.Context, getArtifactDependencyGraphInput *GetArtifactDependencyGraphInput) (*GetArtifactDependencyGraphOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/dependencyGraph/artifact", getArtifactDependencyGraphInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetArtifactDependencyGraphOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type Build struct {
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	PackageType *string `json:"pkg_type,omitempty"`
	Sha256      *string `json:"sha256,omitempty"`
	ComponentId *string `json:"component_id,omitempty"`
}

type GetBuildDependencyGraphInput struct {
	ArtifactoryId *string `json:"artifactory_id,omitempty"`
	BuildName     *string `json:"build_name,omitempty"`
	BuildNumber   *string `json:"build_number,omitempty"`
}

type GetBuildDependencyGraphOutput struct {
	Build      *Build            `json:"build,omitempty"`
	Components *[]GraphComponent `json:"components,omitempty"`
}

// Description: Gets a dependency graph for the components in a given build
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Components.GetBuildDependencyGraph(ctx, getBuildDependencyGraphInput)
func (s *ComponentsService) GetBuildDependencyGraph(ctx context.Context, getBuildDependencyGraphInput *GetBuildDependencyGraphInput) (*GetBuildDependencyGraphOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/dependencyGraph/build", getBuildDependencyGraphInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetBuildDependencyGraphOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type GetArtifactComparisonInput struct {
	SourceArtifactPath *string `json:"source_artifact_path,omitempty"`
	TargetArtifactPath *string `json:"target_artifact_path,omitempty"`
}

type ComparisonComponent struct {
	ComponentName *string `json:"component_name,omitempty"`
	ComponentId   *string `json:"component_id,omitempty"`
	PackageType   *string `json:"package_type,omitempty"`
	Version       *string `json:"version,omitempty"`
	Created       *string `json:"created,omitempty"`
	Modified      *string `json:"modified,omitempty"`
}

type GetArtifactComparisonOutput struct {
	SourceArtifact *Artifact              `json:"source_artifact,omitempty"`
	TargetArtifact *Artifact              `json:"target_artifact,omitempty"`
	Removed        *[]ComparisonComponent `json:"removed,omitempty"`
	Added          *[]ComparisonComponent `json:"added,omitempty"`
	Unchanged      *[]ComparisonComponent `json:"unchanged,omitempty"`
}

// Description: Compares two artifacts and produces the difference between them
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Components.GetArtifactComparison(ctx, getArtifactComparisonInput)
func (s *ComponentsService) GetArtifactComparison(ctx context.Context, getArtifactComparisonInput *GetArtifactComparisonInput) (*GetArtifactComparisonOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/dependencyGraph/artifactDelta", getArtifactComparisonInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetArtifactComparisonOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type GetBuildComparisonInput struct {
	SourceArtifactoryId *string `json:"source_artifactory_id,omitempty"`
	SourceBuildName     *string `json:"source_build_name,omitempty"`
	SourceBuildNumber   *string `json:"source_build_number,omitempty"`

	TargetArtifactoryId *string `json:"target_artifactory_id,omitempty"`
	TargetBuildName     *string `json:"target_build_name,omitempty"`
	TargetBuildNumber   *string `json:"target_build_number,omitempty"`
}

type GetBuildComparisonOutput struct {
	SourceBuild *Build                 `json:"source_build,omitempty"`
	TargetBuild *Build                 `json:"target_build,omitempty"`
	Removed     *[]ComparisonComponent `json:"removed,omitempty"`
	Added       *[]ComparisonComponent `json:"added,omitempty"`
	Unchanged   *[]ComparisonComponent `json:"unchanged,omitempty"`
}

// Description: Compares two builds and produces the difference between them
// Security:  Requires a valid user with "View Components" permission
// Usage: client.V1.Components.GetBuildComparison(ctx, getBuildComparisonInput)
func (s *ComponentsService) GetBuildComparison(ctx context.Context, getBuildComparisonInput *GetBuildComparisonInput) (*GetBuildComparisonOutput, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/dependencyGraph/artifactDelta", getBuildComparisonInput)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetBuildComparisonOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type ExportComponentDetailsInput struct {
	Violations     *bool   `json:"violations,omitempty"`
	License        *bool   `json:"license,omitempty"`
	Security       *bool   `json:"security,omitempty"`
	ExcludeUnknown *bool   `json:"exclude_unkown,omitempty"`
	ComponentName  *string `json:"component_name,omitempty"`
	PackageType    *string `json:"package_type,omitempty"`
	OutputFormat   *string `json:"output_format,omitempty"`
	Sha256         *string `json:"sha_256"`
}

// Description: Export component details.
// Security:  Requires a valid user with "Read Components" permission
// Usage: client.V1.Components.ExportComponentDetails(ctx, exportComponentDetailsInput, writer)
func (s *ComponentsService) ExportComponentDetails(ctx context.Context, exportComponentDetailsInput *ExportComponentDetailsInput, writer io.Writer) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("POST", "/api/v1/dependencyGraph/artifactDelta", exportComponentDetailsInput)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, writer)
	return resp, err
}
