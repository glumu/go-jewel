package haihe

import (
	"github.com/go-jewel/haihe/autoscaling"
	"github.com/go-jewel/haihe/configmanager"
	"github.com/go-jewel/haihe/identity"
	"github.com/go-jewel/haihe/lbaas"
	"github.com/go-jewel/haihe/monitor"
	"github.com/go-jewel/haihe/notification"
	"github.com/go-jewel/haihe/orchestration"
	"github.com/go-jewel/haihe/rds_trove"
	"github.com/go-jewel/haihe/resource"
	"github.com/go-jewel/haihe/storage"
	"github.com/go-jewel/haihe/taskcube"
	"github.com/go-jewel/haihe/uitls"
)

func NewClient(server string, endpoint string, isMock bool, sourceService string, logCB func(logger utils.Logger)) interface{} {
	switch server {
	case "monitor":
		return &monitor.MonitorClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "orchestration":
		return &orchestration.OrchestrationClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "identity":
		return &identity.IdentityClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "resource":
		return &resource.ResourceClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "storage":
		return &storage.StorageClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "lbaas":
		return &lbaas.LbaasClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "autoscaling":
		return &autoscaling.AutoscalingClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "notification":
		return &notification.NotificationClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "taskcube":
		return &taskcube.TaskcubeClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "configmanager":
		return &configmanager.ConfigClient{
			BaseClient: utils.BaseClient{
				Endpoint:       endpoint,
				IsMock:         isMock,
				XServiceSource: sourceService,
				LogCallback:    logCB,
			},
		}
	case "rds_trove":
		return &rds_trove.RdsTroveClient{
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
