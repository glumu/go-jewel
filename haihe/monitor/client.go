package monitor

import "github.com/go-jewel/haihe/uitls"

// ou
type MonitorClient struct {
	utils.BaseClient
}

var (
	CreateGroupUrl string = "/createElasticExpansionGroup"
	DeleteGroupUrl string = "/deleteElasticExpansionGroup"
	UpdateGroupUrl string = "/updateElasticExpansionGroup"
	SearchGroupUrl string = "/searchElasticExpansionGroup"

	BindHostToGroupUrl   string = "/bindHostToElasticExpansionGroup"
	UnBindHostToGroupUrl string = "/unbindHostToElasticExpansionGroup"

	CreateStrategyUrl string = "/createElasticExpansionStrategy"
	UpdateStrategyUrl string = "/updateElasticExpansionStrategy"
	DeleteStrategyUrl string = "/deleteElasticExpansionStrategy"

	SyncHostToGroupUrl string = "/syncHostToElasticExpansionGroup"
)
