package platform

import "github.com/glumu/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	ListAvailableZoneUrl = "/listAvailableZone"
)
