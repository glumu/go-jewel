package autoscaling

type CreateGroupRequest struct {
	BackendMax               int    `json:"backendMax"`
	BackendMin               int    `json:"backendMin"`
	CoolTime                 int    `json:"coolTime"`
	EnableLb                 bool   `json:"enableLb"`
	EnvironmentId            string `json:"environmentId"`
	HealthCheckInterval      int    `json:"healthCheckInterval"`
	HealthCheckType          string `json:"healthCheckType"`
	HealthCheckWaitingPeriod int    `json:"healthCheckWaitingPeriod"`
	InstanceConfId           string `json:"instanceConfId"`
	InstanceRemovePolicy     string `json:"instanceRemovePolicy"`
	LbId                     string `json:"lbId"`
	MonitorId                string `json:"monitorId"`
	Name                     string `json:"name"`
	PrivateNetworkId         string `json:"privateNetworkId"`
	ProjectId                string `json:"projectId"`
	RegionId                 string `json:"regionId"`
	ServicePort              int    `json:"servicePort"`
	Status                   string `json:"status"`
	SubnetId                 string `json:"subnetId"`
	Weight                   int    `json:"weight"`
	ZoneId                   string `json:"zoneId"`
}

type Volume struct {
	IsSystem bool   `json:"isSystem"`
	Type     string `json:"type"`
	Size     int    `json:"size"`
}

type CreateConfigurationRequest struct {
	Cpu             int      `json:"cpu"`
	ImageId         string   `json:"imageId"`
	KeypairId       string   `json:"keypairId"`
	Memory          int      `json:"memory"`
	Name            string   `json:"name"`
	OwnerId         string   `json:"ownerId"`
	Password        string   `json:"password"`
	PasswordLogin   bool     `json:"passwordLogin"`
	RegionId        string   `json:"regionId"`
	SecurityGroupId string   `json:"securityGroupId"`
	Type            string   `json:"type"`
	UserData        string   `json:"userData"`
	Volumes         []Volume `json:"volumes"`
	ZoneId          string   `json:"zoneId"`
	FireWall        string   `json:"fireWall"`
}

type CreatePolicyRequest struct {
	ActionNum       int    `json:"actionNum"`
	ActionType      string `json:"actionType"`
	Cd              int    `json:"cd"`
	Count           int    `json:"count"`
	GroupId         string `json:"groupId"`
	MonitorType     string `json:"monitorType"`
	Name            string `json:"name"`
	Period          int    `json:"period"`
	TriggerCacl     string `json:"triggerCacl"`
	TriggerItem     string `json:"triggerItem"`
	TriggerOperator string `json:"triggerOperator"`
	TriggerPercent  int    `json:"triggerPercent"`
	Type            string `json:"type"`
}

type DeleteGroupRequest struct {
	Id string `json:"id"`
}

type DeleteConfigurationRequest struct {
	Ids []string `json:"ids"`
}

type DeletePolicyRequest struct {
	Id string `json:"id"`
}

type ScaleHostInstanceRequest struct {
	GroupId     string `json:"groupId"`
	ScaleType   string `json:"scaleType"`
	ScaleNumber int    `json:"scaleNumber"`
}

type ListElasticLogRequest struct {
	GroupId string `json:"groupId"`
}

type GetElasticLogRequest struct {
	Id string `json:"id"`
}
