package cloudsuite

import (
	"github.com/go-jewel/cloudsuite/v1.2/composer"
	"github.com/go-jewel/cloudsuite/v1.2/identity"
	"github.com/go-jewel/cloudsuite/v1.2/platform"
	composer13 "github.com/go-jewel/cloudsuite/v1.3/composer"
	identity13 "github.com/go-jewel/cloudsuite/v1.3/identity"
	platform13 "github.com/go-jewel/cloudsuite/v1.3/platform"
	"github.com/go-jewel/cloudsuite/v1.3/resource_scheduler"
	"github.com/go-jewel/pkg/request"
)

type Service interface {
	Setting(endpoint string, isMock bool, sourceService string, logBC func(logger request.Logger))
}

var services = map[string]map[string]Service{
	"v1.2": {
		"identity": &identity.Client{},
		"platform": &platform.Client{},
		"composer": &composer.Client{},
	},
	"v1.3": {
		"identity":           &identity13.Client{},
		"platform":           &platform13.Client{},
		"composer":           &composer13.Client{},
		"resource_scheduler": &resource_scheduler.Client{},
	},
}

func NewClient(server string, endpoint string, version string, isMock bool, sourceService string, logCB func(logger request.Logger)) interface{} {
	client := services[version][server]
	if client == nil {
		return nil
	}
	client.Setting(endpoint, isMock, sourceService, logCB)
	return client
}
