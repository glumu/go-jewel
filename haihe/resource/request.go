package resource

type GetVMRequest struct {
	Id string `json:"id"`
}

type ListVMRequest struct {
	Ids []string `json:"ids"`
}

type ListZoneInfoRequest struct {
}

type CreateInterfaceRequest struct {
	NetworkId     string      `json:"networkId"`
	SubnetId      string      `json:"subnetId"`
	SwitchIp      string      `json:"switchIp,omitempty"`
	PortNum       string      `json:"portNum,omitempty"`
	MacAddress    string      `json:"macAddress"`
	EnvironmentId string      `json:"environmentId"`
	ZoneId        string      `json:"zoneId"`
	OwnerId       string      `json:"ownerId"`
	DeviceOwner   string      `json:"deviceOwner"`
	IpAddress     interface{} `json:"ipAddress"`
}

type DeleteInterfaceRequest struct {
	Id string `json:"id"`
}

type ListAllPortRequest struct {
	Macs []string `json:"macs"`
}

type GetSubnetRequest struct {
	Id string `json:"id"`
}

type ListNetworkRequest struct {
	IsBaremetal bool   `json:"isBaremetal"`
	ZoneId      string `json:"zoneId"`
}

type CreateNetworkRequest struct {
	EnvironmentId   string `json:"environmentId"`
	Name            string `json:"name"`
	NetworkType     string `json:"networkType"`
	OwnerId         string `json:"ownerId"`
	PhysicalNetwork string `json:"physicalNetwork"`
	Shared          int    `json:"shared"`
	Type            string `json:"type"`
	VlanId          int    `json:"vlanId"`
	ZoneId          string `json:"zoneId"`
}

type CreateSubnetRequest struct {
	Dhcp                 string `json:"dhcp"`
	Dns                  string `json:"dns"`
	HostRoute            string `json:"hostRoute"`
	IpPool               string `json:"ipPool"`
	IpVersion            int    `json:"ipVersion"`
	NetworkId            string `json:"networkId"`
	SubnetCidr           string `json:"subnetCidr"`
	SubnetGateway        string `json:"subnetGateway"`
	SubnetGatewayEnabled string `json:"subnetGatewayEnabled"`
	SubnetName           string `json:"subnetName"`
}

type DeleteNetworkRequest struct {
	Ids []string `json:"ids"`
}

type CreateFirewallRequest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	OwnerId     string `json:"ownerId"`
	ZoneId      string `json:"zoneId"`
}

type DeleteFirewallRequest struct {
	Id string `json:"id"`
}

type CallbackRequest struct {
	RequestId      string      `json:"requestId"`
	ResourceTaskId string      `json:"resourceTaskId"`
	Status         string      `json:"status"`
	ServiceName    string      `json:"serviceName"`
	Action         string      `json:"action"`
	VisitorId      string      `json:"visitorId"`
	ZoneId         string      `json:"zoneId"`
	ErrorMessage   string      `json:"errorMessage"`
	ErrorCode      string      `json:"errorCode"`
	Data           interface{} `json:"data"`
}

type ResizeRequest struct {
	Id     string `json:"id"`
	Cpu    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

type UpdateDiskQosRequest struct {
	Id              string `json:"id"`
	MaxReadBytes    int    `json:"maxReadBytes"`
	MaxReadRequest  int    `json:"maxReadRequest"`
	MaxWriteBytes   int    `json:"maxWriteBytes"`
	MaxWriteRequest int    `json:"maxWriteRequest"`
}
