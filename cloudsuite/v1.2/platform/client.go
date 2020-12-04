package platform

import "github.com/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	ListAvailableZoneUrl = "/listAvailableZone"
)