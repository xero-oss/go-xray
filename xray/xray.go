package xray

import (
	"github.com/atlassian/go-artifactory/v2/artifactory/client"
	"github.com/xero-oss/go-xray/xray/v1"
	"github.com/xero-oss/go-xray/xray/v2"

	"net/http"
)

// Xray is the container for all the api methods
type Xray struct {
	V1 *v1.V1
	V2 *v2.V2
}

// NewClient creates a Xray from a provided base url for an xray instance and a service Xray
func NewClient(baseURL string, httpClient *http.Client) (*Xray, error) {
	c, err := client.NewClient(baseURL, httpClient)

	if err != nil {
		return nil, err
	}

	rt := &Xray{
		V1: v1.NewV1(c),
		V2: v2.NewV2(c),
	}

	return rt, nil
}
