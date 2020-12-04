package responseData

type AttachDiskResponse struct {
	BaseResponse
}

type diskId struct {
	Id string `json:"id"`
}

type CreateDiskResponse struct {
	BaseResponse
	Data []diskId `json:"data"`
}

type detailDiskQos struct {
	MaxReadBytes    int `json:"maxReadBytes"`
	MaxReadRequest  int `json:"maxReadRequest"`
	MaxWriteBytes   int `json:"maxWriteBytes"`
	MaxWriteRequest int `json:"maxWriteRequest"`
}

type detailDisk struct {
	Id         string        `json:"id"`
	Ref        string        `json:"ref"`
	Name       string        `json:"name"`
	DiskType   string        `json:"diskType"`
	DiskTypeId string        `json:"diskTypeId"`
	OwnerId    string        `json:"ownerId"`
	OwnerName  string        `json:"ownerName"`
	Size       int           `json:"size"`
	SizeUsed   int           `json:"sizeUsed"`
	Status     string        `json:"status"`
	TaskStatus string        `json:"taskStatus"`
	UseStatus  string        `json:"useStatus"`
	Qos        detailDiskQos `json:"qos"`
	Path       string        `json:"path"`
	IsNew      bool          `json:"isNew"`
	IsSystem   bool          `json:"isSystem"`
	IsShared   bool          `json:"isShared"`
}

type GetDiskResponse struct {
	BaseResponse
	Data detailDisk `json:"data"`
}

type ResizeDiskResponse struct {
	BaseResponse
}

type UpdateDiskQosResponse struct {
	BaseResponse
}

type CreateDiskSnapshotResponse struct {
	BaseResponse
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}

type ListDiskTypeResponse struct {
	BaseResponse
	Data []struct {
		DatastoreBacken string `json:"datastoreBacken"`
		Description     string `json:"description"`
		Enabled         bool   `json:"enabled"`
		EnvironmentId   string `json:"environmentId"`
		Id              string `json:"id"`
		Name            string `json:"name"`
		Ref             string `json:"ref"`
	} `json:"data"`
}

type DeleteDisksResponse struct {
	BaseResponse
}
