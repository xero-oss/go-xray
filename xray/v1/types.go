package v1

import "github.com/atlassian/go-artifactory/v2/artifactory/client"

type Service struct {
	client *client.Client
}

type V1 struct {
	common Service

	// Services used for talking to different parts of the Xray API.
	BinaryManagers *BinaryManagersService
	Components     *ComponentsService
	Configuration  *ConfigurationService
	Integrations   *IntegrationsService
	Issues         *IssuesService
	Permissions    *PermissionsService
	Policies       *PoliciesService
	Reports        *ReportsService
	Scanning       *ScanningService
	Summary        *SummaryService
	System         *SystemService
	Users          *UsersService
	Violations     *ViolationsService
}
