package configmanager

import "github.com/go-jewel/haihe/uitls"

type ConfigClient struct {
	utils.BaseClient
	MockApi []string
}

var (
	GetConfigUrl   string = "/getConfig"
)
