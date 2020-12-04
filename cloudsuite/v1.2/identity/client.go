package identity

import "github.com/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	GetVdcUrl  = "/getVdc"
	ListVdcUrl = "/listVdc"

	GetProjectUrl    = "/getProject"
	ListMyProjectUrl = "/listMyProject"
)
