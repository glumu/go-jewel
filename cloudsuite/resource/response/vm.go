package responseData

type vmDetailEnv struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type vmDetailZone struct {
	Id     string `json:"id"`
	Ref    string `json:"ref"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type vmDetailFlavor struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Cpu      int    `json:"cpu"`
	DiskSize int    `json:"diskSize"`
	Memory   int    `json:"memory"`
}

type vmDetailDisk struct {
	Id               string `json:"id"`
	Ref              string `json:"ref"`
	Name             string `json:"name"`
	DiskType         string `json:"diskType"`
	DiskTypeId       string `json:"diskTypeId"`
	Size             int    `json:"size"`
	SizeUsed         int    `json:"sizeUsed"`
	Status           string `json:"status"`
	TaskStatus       string `json:"taskStatus"`
	UseStatus        string `json:"useStatus"`
	ReadBytesSecond  int    `json:"readBytesSecond"`
	ReadIopsSecond   int    `json:"readIopsSecond"`
	WriteBytesSecond int    `json:"writeBytesSecond"`
	WriteIopsSecond  int    `json:"writeIopsSecond"`
	Path             string `json:"path"`
	IsNew            bool   `json:"isNew"`
	IsSystem         bool   `json:"isSystem"`
	IsShared         bool   `json:"isShared"`
}

type vmDetailInterfaces struct {
	Id         string `json:"id"`
	Ip         string `json:"ip"`
	MacAddress string `json:"macAddress"`
	NetworkId  string `json:"networkId"`
	SubnetId   string `json:"subnetId"`
	Status     string `json:"status"`
}

type vmDetail struct {
	Id               string               `json:"id"`
	Ref              string               `json:"ref"`
	Name             string               `json:"name"`
	OwnerId          string               `json:"ownerId"`
	OwnerName        string               `json:"ownerName"`
	Status           string               `json:"status"`
	CreateTime       string               `json:"createTime"`
	Environment      vmDetailEnv          `json:"environment"`
	Zone             vmDetailZone         `json:"zone"`
	Flavor           vmDetailFlavor       `json:"flavor"`
	Disk             []vmDetailDisk       `json:"disk"`
	HostResize       bool                 `json:"hostResize"`
	CpuHostResize    bool                 `json:"cpuHostResize"`
	MemoryHostResize bool                 `json:"memoryHlostResize"`
	Interfaces       []vmDetailInterfaces `json:"interfaces"`
}

type GetVMResponse struct {
	BaseResponse
	Data vmDetail `json:"data"`
}

type ResizeFlavorResponse struct {
	BaseResponse
}
