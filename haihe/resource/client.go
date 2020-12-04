package resource

import "github.com/glumu/go-jewel/haihe/uitls"

// ou
type ResourceClient struct {
	utils.BaseClient
}

var (
	GetVMUrl           string = "/getVirtualMachine"
	ListVMUrl          string = "/listVirtualMachine"
	ListZoneInfoUrl    string = "/listZoneInfo"
	CreateInterfaceUrl string = "/createInterface"
	DeleteInterfaceUrl string = "/deleteInterface"
	ListAllPortUrl     string = "/listAllPort"
	GetSubnetUrl       string = "/getSubnet"
	ListNetworkUrl     string = "/listNetwork"
	CreateNetworkUrl   string = "/createNetwork"
	CreateSubnetUrl    string = "/createSubnet"
	DeleteNetworkUrl   string = "/deleteNetwork"
	ResizeVMUrl        string = "/resizeVirtualMachine"
	//firewall
	CreateFirewallUrl string = "/createFirewall"
	DeleteFirewallUrl string = "/deleteFirewall"
	CallbackUrl       string = "/callback"
	//disk
	UpdateDiskQosUrl string = "/updateDiskQos"
)
