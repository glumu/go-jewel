package resource

// VM =====================================================
type GetVMRequest struct {
	Id string `json:"id"`
}

// Flavor =================================================
type ResizeFlavorRequest struct {
	Id     string `json:"id"`
	Cpu    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

// Disk ===================================================
type CreateDiksRequest struct {
	Name   string `json:"name"`
	ZoneId string `json:"zoneId"`
	Size   int    `json:"size"`
	Count  int    `json:"count"`
}

type ResizeDiskRequest struct {
	Id         string `json:"id"`
	ExtendSize int    `json:"extendSize"`
}

type GetDiskRequest struct {
	Id string `json:"id"`
}

type UpdateDiskQosRequest struct {
	Id              string `json:"id"`
	MaxReadBytes    int    `json:"maxReadBytes"`
	MaxReadRequest  int    `json:"maxReadRequest"`
	MaxWriteBytes   int    `json:"maxWriteBytes"`
	MaxWriteRequest int    `json:"maxWriteRequest"`
}

type CopyDiskRequest struct {
	OwnerId string `json:"ownerId"`
}

type AttachDiskRequest struct {
	VirtualMachineId string `json:"virtualMachineId"`
	DiskId           string `json:"diskId"`
}

type DetachDiskRequest struct {
	VirtualMachineId string `json:"virtualMachineId"`
	DiskId           string `json:"diskId"`
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
