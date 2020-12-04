package orchestration

type GetDeploymentRequest struct {
	DeploymentId string `json:"deploymentId"`
}

type CreateKvmDeploymentInputs struct {
	Name                string   `json:"name"`
	OwnerId             string   `json:"owner_id"`
	UserId              string   `json:"user_id"`
	Image               string   `json:"image"`
	EnvironmentId       string   `json:"environment_id"`
	ZoneId              string   `json:"zone_id"`
	UseExternalResource bool     `json:"use_external_resource"`
	PrivateResourceId   string   `json:"private_resource_id"`
	SubnetResourceId    string   `json:"subnet_resource_id"`
	Cpu                 int      `json:"cpu"`
	Memory              int      `json:"memory"`
	MinNum              int      `json:"min_num"`
	MaxNum              int      `json:"max_num"`
	AdminPassword       string   `json:"admin_password"`
	KeypairId           string   `json:"keypair_id"`
	Script              []string `json:"script"`
	FireWall            string   `json:"fire_wall"`
	DiskType            string   `json:"disk_type"`
}

type CreateKvmDeploymentRequest struct {
	Name               string                    `json:"name"`
	BlueprintVersionId string                    `json:"blueprintVersionId"`
	OwnerId            string                    `json:"ownerId"`
	Timeout            int                       `json:"timeout"`
	Inputs             CreateKvmDeploymentInputs `json:"inputs"`
}

type CreateLbDeploymentInputs struct {
	Name                string   `json:"name"`
	OwnerId             string   `json:"owner_id"`
	Image               string   `json:"image"`
	EnvironmentId       string   `json:"environment_id"`
	ZoneId              string   `json:"zone_id"`
	UseExternalResource bool     `json:"use_external_resource"`
	Cpu                 int      `json:"cpu"`
	Memory              int      `json:"memory"`
	PrivateResourceId   string   `json:"private_resource_id"`
	SubnetResourceId    string   `json:"subnet_resource_id"`
	LbId                string   `json:"lb_id"`
	ListenerId          string   `json:"listener_id"`
	ServicePort         int      `json:"service_port"`
	Weight              int      `json:"weight"`
	MinNum              int      `json:"min_num"`
	MaxNum              int      `json:"max_num"`
	AdminPassword       string   `json:"admin_password"`
	KeypairId           string   `json:"keypair_id"`
	Script              []string `json:"script"`
	FireWall            string   `json:"fire_wall"`
	DiskType            string   `json:"disk_type"`
}

type CreateLbDeploymentRequest struct {
	Name               string                   `json:"name"`
	BlueprintVersionId string                   `json:"blueprintVersionId"`
	OwnerId            string                   `json:"ownerId"`
	Timeout            int                      `json:"timeout"`
	Inputs             CreateLbDeploymentInputs `json:"inputs"`
}

type FirewallRule struct {
	Direction string `json:"direction"`
	Protocol  string `json:"protocol"`
	MaxPort   string `json:"max_port"`
	MinPort   string `json:"min_port"`
}

type CreateRdsDeploymentRunDataTrove struct {
	TenantId     string `json:"tenant_id"`
	TransportUrl string `json:"transport_url"`
}
type CreateRdsDeploymentRunData struct {
	VersionTag string                          `json:"version_tag"`
	SourceAddr string                          `json:"source_addr"`
	Dest       string                          `json:"dest"`
	Trove      CreateRdsDeploymentRunDataTrove `json:"trove"`
}
type CreateRdsDeploymentInputs struct {
	OwnerId       string `json:"owner_id"`
	EnvironmentId string `json:"environment_id"`
	ZoneId        string `json:"zone_id"`
	// vm
	VmName            string `json:"vm_name"`
	ImageId           string `json:"image"`
	Cpu               int    `json:"cpu"`
	Memory            int    `json:"memory"`
	AdminPassword     string `json:"admin_password"`
	NetworkResourceId string `json:"network_id"`
	SubnetResourceId  string `json:"subnet_id"`
	// firewall
	FirewallName             string         `json:"firewall_name"`
	FirewallExternalResource bool           `json:"firewall_external_resource"`
	FirewallId               string         `json:"firewall_id"`
	FirewallRules            []FirewallRule `json:"rules"`
	// volumn
	VmSize int `json:"vm_size"`
	// data
	RunData CreateRdsDeploymentRunData `json:"run_data"`
}

type CreateRdsDeploymentRequest struct {
	Name               string                    `json:"name"`
	BlueprintVersionId string                    `json:"blueprintVersionId"`
	OwnerId            string                    `json:"ownerId"`
	Timeout            int                       `json:"timeout"`
	Inputs             CreateRdsDeploymentInputs `json:"inputs"`
}

type DeleteDeploymentRequest struct {
	DeploymentIds []string `json:"deploymentIds"`
}

type ListNodeInstanceRequest struct {
	DeploymentId string `json:"deploymentId"`
}

type ListDeploymentRequest struct {
	DeploymentId string `json:"Id"`
}

type UpdateDeploymentInputs struct {
	MinNum int `json:"min_num"`
	MaxNum int `json:"max_num"`
}

type UpdateDeploymentRequest struct {
	Id                 string                 `json:"id" binding:"required"`
	BlueprintVersionId string                 `json:"blueprintVersionId"`
	Inputs             UpdateDeploymentInputs `json:"inputs" binding:"required"`
	SkipInstall        bool                   `json:"skipInstall"`
	SkipUninstall      bool                   `json:"skipUninstall"`
	SkipReinstall      bool                   `json:"skipReinstall"`
	Force              bool                   `json:"force"`
	IgnoreFailure      bool                   `json:"ignoreFailure"`
	InstallFirst       bool                   `json:"installFirst"`
	Preview            bool                   `json:"preview"`
	ReinstallList      []string               `json:"reinstallList"`
}

type ExecScaleParameters struct {
	ScalableEntityName string `json:"scalable_entity_name"`
	Delta              int    `json:"delta"`
}
type ExecHealParameters struct {
	NodeInstanceId string `json:"node_instance_id"`
}
type StartExecutionRequest struct {
	DeploymentId string `json:"deploymentId"`
	WorkflowId   string `json:"workflowId"`
}
type StartExecScaleRequest struct {
	StartExecutionRequest
	Parameters ExecScaleParameters
}
type StartExecHealRequest struct {
	StartExecutionRequest
	Parameters ExecHealParameters
}

type ListEventRequest struct {
	ExecutionId string `json:"executionId"`
}
