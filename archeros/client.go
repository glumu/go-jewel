package archeros

import (
	"github.com/go-jewel/archeros/resource"
	"github.com/go-jewel/archeros/uitls"
)

func NewClient(server string, endpoint string, isMock bool, sourceService string, logCB func(logger utils.Logger)) interface{} {
	switch server {
	case "resource":
		return &resource.ResourceClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	}
	return nil
}
