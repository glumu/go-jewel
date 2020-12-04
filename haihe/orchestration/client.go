package orchestration

import "github.com/go-jewel/haihe/uitls"

type OrchestrationClient struct {
	utils.BaseClient
}

var (
	GetDeploymentUrl    string = "/getDeployment"
	CreateDeploymentUrl string = "/createDeployment"
	DeleteDeploymentUrl string = "/deleteDeployment"
	ListDeploymentUrl   string = "/listDeployment"
	UpdateDeploymentUrl string = "/updateDeployment"
	ListNodeInstanceUrl string = "/listNodeInstance"
	StartExecutionUrl   string = "/startExecution"
	ListEventUrl        string = "/listEvent"
)
