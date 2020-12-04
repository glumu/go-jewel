package configmanager

import "github.com/glumu/go-jewel/haihe/uitls"

type ConfigClient struct {
	utils.BaseClient
	MockApi []string
}

var (
	GetConfigUrl   string = "/getConfig"
)
