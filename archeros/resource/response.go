package resource

type detailFlavor struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Cpu      int    `json:"cpu"`
	DiskSize int    `json:"diskSize"`
	Memory   int    `json:"memory"`
}
type detailDisk struct {
	Id               string `json:"id"`
	Ref              string `json:"ref"`
	Name             string `json:"name"`
	DiskType         string `json:"diskType"`
	Size             int    `json:"size"`
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

type diskId struct {
	Id string `json:"id"`
}

type detailEnv struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type detailZone struct {
	Id     string `json:"id"`
	Ref    string `json:"ref"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type detailInterfaces struct {
	Id          string `json:"id"`
	Ip          string `json:"ip"`
	MacAddress  string `json:"macAddress"`
	NetworkId   string `json:"networkId"`
	Status      string `json:"status"`
	NetworkName string `json:"networkName"`
}

type detail struct {
	Id               string             `json:"id"`
	Ref              string             `json:"ref"`
	Name             string             `json:"name"`
	OwnerId          string             `json:"ownerId"`
	OwnerName        string             `json:"ownerName"`
	Status           string             `json:"status"`
	CreateTime       string             `json:"createTime"`
	Environment      detailEnv          `json:"environment"`
	Zone             detailZone         `json:"zone"`
	Flavor           detailFlavor       `json:"flavor"`
	Disk             []detailDisk       `json:"disk"`
	HostResize       bool               `json:"hostResize"`
	CpuHostResize    bool               `json:"cpuHostResize"`
	MemoryHostResize bool               `json:"memoryHlostResize"`
	Interfaces       []detailInterfaces `json:"interfaces"`
}

// ============================================================
// VM =========================================================
type GetVMResponse struct {
	RequestId string `json:"requestId"`
	Data      detail `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type AttachDiskResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

// Flavor =====================================================
type ResizeFlavorResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

// Disk =======================================================
type CreateDiskResponse struct {
	RequestId string   `json:"requestId"`
	TimeStamp int      `json:"timestamp"`
	Data      []diskId `json:"data"`
}

type ResizeDiskResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type GetDiskResponse struct {
	RequestId string     `json:"requestId"`
	Data      detailDisk `json:"data"`
	TimeStamp int        `json:"timestamp"`
}

type UpdateDiskQosResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type CreateDiskSnapshotResponse struct {
	RequestId string `json:"requestId"`
	Data      struct {
		Id string `json:"id"`
	} `json:"data"`
	TimeStamp int `json:"timestamp"`
}

type ListDiskTypeResponse struct {
	RequestId string `json:"requestId"`
	Data      []struct {
		DatastoreBacken string `json:"datastoreBacken"`
		Description     string `json:"description"`
		Id              string `json:"id"`
		Name            string `json:"name"`
	} `json:"data"`
	TimeStamp int `json:"timestamp"`
}
