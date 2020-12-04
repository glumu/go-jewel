package orchestration

type DeploymentInfo struct {
	Id                 string `json:"id"`
	Status             string `json:"status"`
	BlueprintVersionId string `json:"blueprintVersionId"`
}

type GetDeploymentResponse struct {
	RequestId string `json:"requestId"`
	Data      struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		OwnerId string `json:"ownerId"`
		Status  string `json:"status"`
	} `json:"data"`
}

type CreateDeploymentResponse struct {
	RequestId string         `json:"requestId"`
	Data      DeploymentInfo `json:"data"`
}

type DeleteDeploymentResponse struct {
	RequestId string `json:"requestId"`
}

type NodeInstance struct {
	Id             string `json:"id"`
	EnvironmentId  string `json:"environmentId"`
	ZoneId         string `json:"zoneId"`
	ProjectName    string `json:"projectName"`
	OwnerId        string `json:"ownerId"`
	Name           string `json:"name"`
	IsNew          bool   `json:"isNew"`
	Type           string `json:"resourceType"`
	State          string `json:"state"`
	Ref            string `json:"ref"`            // openstack node id
	NodeInstanceId string `json:"nodeInstanceId"` // cloudify id
	ResourceId     string `json:"resourceId"`     // resource id
}

type ListNodeInstanceResponse struct {
	Data []NodeInstance `json:"data"`
}

type Deployment struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	OwnerId              string `json:"ownerId"`
	BlueprintId          string `json:"blueprintId"`
	BlueprintName        string `json:"blueprintName"`
	BlueprintVersionId   string `json:"blueprintVersionId"`
	BlueprintVersionName string `json:"blueprintVersionName"`
	Status               string `json:"status"`
}

type ListDeploymentResponse struct {
	Data []Deployment `json:"data"`
}

type UpdateDeploymentResponse struct {
	Id           string `json:"id"`
	State        string `json:"state"`
	DeploymentId string `json:"deployment_id"`
	ExecutionId  string `json:"execution_id"`
}

type ExecutionData struct {
	Id            string `json:"id"`
	DeploymentId  string `json:"deployment_id"`
	WorkflowId    string `json:"workflow_id"`
	Error         string `json:"error"`
	BlueprintId   string `json:"blueprint_id"`
	StatusDisplay string `json:"status_display"`
	StartedAt     string `json:"started_at"`
	EndedAt       string `json:"ended_at"`
}

type StartExecutionResponse struct {
	Data ExecutionData
}

type Event struct {
	BlueprintId             string `json:"blueprintId"`
	DeploymentId            string `json:"deploymentId"`
	ExecutionId             string `json:"executionId"`
	NodeInstanceId          string `json:"nodeInstanceId"`
	ResourceId              string `json:"resourceId"`
	ResourceName            string `json:"resourceName"`
	ResourceOperation       string `json:"resourceOperation"`
	ResourceOperationResult string `json:"resourceOperationResult"`
	ResourceType            string `json:"resourceType"`
	WorkflowId              string `json:"workflowId"`
	ErrorMessage            string `json:"errorMessage"`
}

type ListEventResponse struct {
	Data []Event `json:"data"`
}
