package resource

import "github.com/go-jewel/archeros/uitls"

// ou
type ResourceClient struct {
	utils.BaseClient
}

var (
	// vm
	GetVMUrl        = "/getVirtualMachine"
	ResizeFlavorUrl = "/resizeVirtualMachine"
	//disk
	CreateDiskUrl         = "/createDisk"
	DeleteDiskUrl         = "/deleteDisk"
	GetDiskUrl            = "/getDisk"
	ResizeDiskUrl         = "/resizeDisk"
	AttachDiskUrl         = "/attachDisk"
	DetachDiskUrl         = "/detachDisk"
	UpdateDiskQosUrl      = "/updateDiskQos"
	CreateDiskSnapshotUrl = "/createDiskSnapshot"
	DeleteDiskSnapshotUrl = "/deleteDiskSnapshot"
	ListDiskTypeUrl       = "/listDiskType"
)
