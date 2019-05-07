package v1

import (
	"context"
	"net/http"
)

type ConfigurationService Service

type SystemParameters struct {
	SslInsecure                   *bool `json:"ssl_insecure,omitempty"`
	MaxDiskDataUsage              *int  `json:"maxDiskDataUsage,omitempty"`
	MonitorSamplingInterval       *int  `json:"monitorSamplingInterval,omitempty"`
	MailNoSsl                     *bool `json:"mailNoSsl,omitempty"`
	MessageMaxTTL                 *int  `json:"messageMaxTTL,omitempty"`
	JobInterval                   *int  `json:"jobInterval,omitempty"`
	AllowSendingAnalytics         *bool `json:"allowSendingAnalytics,omitempty"`
	HTTPSPorts                    *int  `json:"httpsPort,omitempty"`
	EnableTlsConnectionToRabbitMQ *bool `json:"enableTlsConnectionToRabbitMQ,omitempty"`
}

// Description: Gets the current system configuration
// Security:  Requires an admin user
// Usage: client.V1.Configuration.GetSystemParameters(ctx)
func (s *ConfigurationService) GetSystemParameters(ctx context.Context) (*SystemParameters, *http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("GET", "/api/v1/configuration/systemParameters", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	parameters := new(SystemParameters)
	resp, err := s.client.Do(ctx, req, &parameters)
	return parameters, resp, err
}

// Description: Updates the current system configuration
// Security:  Requires an admin user
// Usage: client.V1.Configuration.UpdateSystemParameters(ctx, parameters)
func (s *ConfigurationService) UpdateSystemParameters(ctx context.Context, parameters *SystemParameters) (*http.Response, error) {
	req, err := s.client.NewJSONEncodedRequest("PUT", "/api/v1/configuration/systemParameters", parameters)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
