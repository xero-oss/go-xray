package v1

import "github.com/atlassian/go-artifactory/v2/artifactory/client"

func String(v string) *string { return &v }

func NewV1(client *client.Client) *V1 {
	v := &V1{}
	v.common.client = client

	v.BinaryManagers = (*BinaryManagersService)(&v.common)
	v.Components = (*ComponentsService)(&v.common)
	v.Configuration = (*ConfigurationService)(&v.common)
	v.Integrations = (*IntegrationsService)(&v.common)
	v.Issues = (*IssuesService)(&v.common)
	v.Permissions = (*PermissionsService)(&v.common)
	v.Policies = (*PoliciesService)(&v.common)
	v.Reports = (*ReportsService)(&v.common)
	v.Scanning = (*ScanningService)(&v.common)
	v.Summary = (*SummaryService)(&v.common)
	v.System = (*SystemService)(&v.common)
	v.Users = (*UsersService)(&v.common)
	v.Violations = (*ViolationsService)(&v.common)

	return v
}
