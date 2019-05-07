package v2

import "github.com/atlassian/go-artifactory/v2/artifactory/client"

type Service struct {
	client *client.Client
}

type V2 struct {
	common Service

	// Services used for talking to different parts of the Xray API.
	Watches *WatchesService
}
