package requestData

type AttachDiskRequest struct {
	VirtualMachineId string `json:"virtualMachineId"`
	DiskId           string `json:"diskId"`
}

type DetachDiskRequest struct {
	VirtualMachineId string `json:"virtualMachineId"`
	DiskId           string `json:"diskId"`
}

type CreateDiskRequest struct {
	Name          string `json:"name"`
	OwnerId       string `json:"ownerId"`
	ZoneId        string `json:"zoneId"`
	EnvironmentId string `json:"environmentId"`
	Size          int    `json:"size"`
	Count         int    `json:"count"`
	DiskType      string `json:"diskType,omitempty"`
}

type GetDiskRequest struct {
	Id string `json:"id"`
}

type ResizeDiskRequest struct {
	Id         string `json:"id"`
	ExtendSize int    `json:"extendSize"`
}

type UpdateDiskQosRequest struct {
	Id              string `json:"id"`
	MaxReadBytes    int    `json:"maxReadBytes"`
	MaxReadRequest  int    `json:"maxReadRequest"`
	MaxWriteBytes   int    `json:"maxWriteBytes"`
	MaxWriteRequest int    `json:"maxWriteRequest"`
}

type CreateDiskSnapshotRequest struct {
	Name        string `json:"name"`
	DiskId      string `json:"diskId"`
	Description string `json:"description"`
}

type DeleteDiskSnapshotRequest struct {
	Id string `json:"id"`
}

type ListDiskTypeRequest struct {
	ZoneId        string `json:"zoneId"`
	EnvironmentId string `json:"environmentId"`
}

type DeleteDisksRequest struct {
	Ids []string `json:"ids"`
}
