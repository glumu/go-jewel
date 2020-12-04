package identity

import "github.com/glumu/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	GetVdcUrl  = "/getVdc"
	ListVdcUrl = "/listVdc"

	GetProjectUrl    = "/getProject"
	ListMyProjectUrl = "/listMyProject"
)
