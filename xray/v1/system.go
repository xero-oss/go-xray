package v1

import (
	"context"
	"net/http"
)

type SystemService Service

type SystemMonitoringProblem struct {
	Severity *string   `json:"severity,omitempty"`
	Services *[]string `json:"services,omitempty"`
	Problem  *string   `json:"problem,omitempty"`
}

type GetSystemMonitoringStatusOutput struct {
	Problems *[]SystemMonitoringProblem `json:"problems,omitempty"`
}

// Description: Gets system monitoring status
// Security:  Requires an admin user
// Usage: client.V1.System.GetMonitoringStatus(ctx)
func (s *SystemService) GetMonitoringStatus(ctx context.Context) (*GetSystemMonitoringStatusOutput, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/monitor", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(GetSystemMonitoringStatusOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

type PingRequestOutput struct {
	Status *string `json:"status,omitempty"`
}

// Description: Sends a ping request
// Security:  Requires a valid user
// Usage: client.V1.System.Ping(ctx)
func (s *SystemService) Ping(ctx context.Context) (*PingRequestOutput, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/system/ping", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(PingRequestOutput)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}

// NOTE: This doesn't seem to actually exist
// Description: Sends a ping request to external sources (Global Database, Bintray, etc.)
// Security:  Requires a valid user
// Usage: client.V1.System.ExternalPing(ctx)
// func (s *SystemService) ExternalPing(ctx context.Context) (*PingRequestOutput, *http.Response, error) {
// 	req, err := s.client.NewRequest("GET", "/api/v1/system/external/ping", nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	req.Header.Set("Accept", "application/json")

// 	output := new(PingRequestOutput)
// 	resp, err := s.client.Do(ctx, req, &output)
// 	return output, resp, err
// }

type XrayVersion struct {
	Version  *string `json:"xray_version,omitempty"`
	Revision *string `json:"xray_revision,omitempty"`
}

// Description:  Gets the Xray version and revision you are running
// Security:  Requires a valid user
// Usage: client.V1.System.Version(ctx)
func (s *SystemService) Version(ctx context.Context) (*XrayVersion, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/system/version", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	output := new(XrayVersion)
	resp, err := s.client.Do(ctx, req, &output)
	return output, resp, err
}
