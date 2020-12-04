package lbaas

import "github.com/glumu/go-jewel/haihe/uitls"

// ou
type LbaasClient struct {
	utils.BaseClient
}

var (
	CreateLbUrl                string = "/createLoadBalancer"
	CreateListenerUrl          string = "/createListener"
	CreateLbPollUrl            string = "/createLbPoll"
	UpdateLbRelatedResourceUrl string = "/updateLbRelatedResource"
	ListLbBackendUrl           string = "/listLbBackend"
	GetLbBackendUrl            string = "/getLbBackend"
)
