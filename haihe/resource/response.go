package resource

type DetailFlavor struct {
	Cpu      int    `json:"cpu"`
	DiskSize int    `json:"diskSize"`
	Id       string `json:"id"`
	Memory   int    `json:"memory"`
}
type DetailDisk struct {
	Id               string `json:"id"`
	DiskType         string `json:"diskType"`
	ReadBytesSecond  int    `json:"readBytesSecond"`
	ReadIopsSecond   int    `json:"readIopsSecond"`
	Size             int    `json:"size"`
	Status           string `json:"status"`
	TaskStatus       string `json:"taskStatus"`
	useStatus        string `json:"useStatus"`
	WriteBytesSecond int    `json:"writeBytesSecond"`
	WriteIopsSecond  int    `json:"writeIopsSecond"`
	Path             string `json:"path"`
}
type Detail struct {
	Id         string       `json:"id"`
	Ref        string       `json:"ref"`
	Name       string       `json:"name"`
	OwnerId    string       `json:"ownerId"`
	OwnerName  string       `json:"ownerName"`
	Status     string       `json:"status"`
	CreateTime string       `json:"createTime"`
	Flavor     DetailFlavor `json:"flavor"`
	Disk       []DetailDisk `json:"disk"`
}

type GetVMResponse struct {
	RequestId string `json:"requestId"`
	Data      Detail `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type VM struct {
	Id         string `json:"id"`
	Ref        string `json:"ref"`
	Name       string `json:"name"`
	OwnerId    string `json:"ownerId"`
	OwnerName  string `json:"ownerName"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type ListVMResponse struct {
	RequestId string `json:"requestId"`
	Data      []VM   `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type Zone struct {
	Id               string `json:"id"`
	Uuid             string `json:"uuid"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	EnvironmentId    string `json:"environmentId"`
	CreateTime       string `json:"createTime"`
	UpdateTime       string `json:"updateTime"`
	Status           string `json:"status"`
	IdentityEndpoint string `json:"identityEndpoint"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	TenantID         string `json:"tenantId"`
	TenantName       string `json:"tenantName"`
	DomainName       string `json:"domainName"`
	IsBaremetal      bool   `json:"isBaremetal"`
}

type Environment struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	SyncInterval     float32 `json:"syncInterval"`
	LastSyncTime     string  `json:"lastSyncTime"`
	CloudGatewayName string  `json:"cloudGatewayName"`
	CreateTime       string  `json:"createTime"`
	UpdateTime       string  `json:"updateTime"`
	HyperVisorType   string  `json:"hypervisorType"`
	Type             string  `json:"type"`
	Zones            []Zone  `json:"zones"`
	IsBaremetal      bool    `json:"isBaremetal"`
}

type ListZoneInfoResponse struct {
	Environments []Environment `json:"data"`
}

type Port struct {
	Id         string `json:"id"`
	Ref        string `json:"ref"`
	NetworkId  string `json:"networkId"`
	SubnetId   string `json:"subnetId"`
	IpAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
}

type Network struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Ref                  string   `json:"ref"`
	OwnerId              string   `json:"ownerId"`
	BaremetalNetworkType string   `json:"baremetalNetworkType"`
	Subnets              []Subnet `json:"subnets"`
}

type Subnet struct {
	Cidr    string `json:"cidr"`
	Gateway string `json:"gateway"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Ref     string `json:"ref"`
}

type CreateInterfaceResponse struct {
	RequestId string `json:"requestId"`
	Data      Port   `json:"data"`
	TimeStamp int    `json:"timestamp"`
	Ref       string `json:"ref"`
}

type DeleteInterfaceResponse struct {
}

type InterfaceDetail struct {
	Id        string `json:"id"`
	Ip        string `json:"ip"`
	NetworkId string `json:"networkId"`
}

type ListAllPortResponse struct {
	RequestId string            `json:"requestId"`
	Data      []InterfaceDetail `json:"data"`
	TimeStamp int               `json:"timestamp"`
}

type GetSubnetResponse struct {
	RequestId string `json:"requestId"`
	Data      Subnet `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type ListNetworkResponse struct {
	RequestId string    `json:"requestId"`
	Data      []Network `json:"data"`
	TimeStamp int       `json:"timestamp"`
}

type CreateNetworkResponse struct {
	RequestId string  `json:"requestId"`
	Data      Network `json:"data"`
	TimeStamp int     `json:"timestamp"`
}

type CreateSubnetResponse struct {
	RequestId string  `json:"requestId"`
	Data      Network `json:"data"`
	TimeStamp int     `json:"timestamp"`
}

type DeleteNetworkResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type Firewall struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateFirewallResponse struct {
	RequestId string `json:"requestId"`
	Data      string `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type DeleteFirewallResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type CallbackResponse struct {
	RequestId    string `json:"requestId"`
	Result       string `json:"result"`
	ErrorMessage string `json:"errorMessage"`
}

type ResizeResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type UpdateDiskQosResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}
