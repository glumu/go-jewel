package rds_trove

import "github.com/glumu/go-jewel/haihe/uitls"

type RdsTroveClient struct {
	utils.BaseClient
	MockApi []string
}

var (
	CreateInstanceUrl  string = "/v1.0/%s/instances"
	PutInstanceConfUrl string = "/v1.0/%s/instances/%s"
	GetInstanceConfUrl string = "/v1.0/%s/instances/%s/configuration"

	UserUrl       string = "/v1.0/%s/instances/%s/users"
	DeleteUserUrl string = "/v1.0/%s/instances/%s/users/%s"

	DatabaseUrl       string = "/v1.0/%s/instances/%s/databases"
	DeleteDatabaseUrl string = "/v1.0/%s/instances/%s/databases/%s"

	ConfUrl         string = "/v1.0/%s/configurations"
	ConfItemsUrl    string = "/v1.0/%s/configurations/%s"
	ConfInstanceUrl string = "/v1.0/%s/configurations/%s/instances"

	ActionUrl string = "/v1.0/%s/instances/%s/action"

	CreateClusterUrl string = "/v1.0/%s/clusters"
)
