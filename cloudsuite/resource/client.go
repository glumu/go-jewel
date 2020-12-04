package resource

import "github.com/go-jewel/haihe/uitls"

type ResourceClient struct {
	utils.BaseClient
}

var (
	// vm
	GetVMUrl        = "/getVirtualMachine"
	ResizeFlavorUrl = "/resizeVirtualMachine"
	//disk
	AttachDiskUrl    = "/attachDisk"
	DetachDiskUrl    = "/detachDisk"
	CreateDiskUrl    = "/createDisk"
	DeleteDiskUrl    = "/deleteDisk"
	GetDiskUrl       = "/getDisk"
	ResizeDiskUrl    = "/resizeDisk"
	UpdateDiskQosUrl = "/updateDiskQos"
	ListDiskTypeUrl  = "/listDiskType"

	CreateDiskSnapshotUrl = "/createDiskSnapshot"
	DeleteDiskSnapshotUrl = "/deleteDiskSnapshot"
)
