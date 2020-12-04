package storage

import "github.com/go-jewel/haihe/uitls"

// ou
type StorageClient struct {
	utils.BaseClient
}

var (
	AddLunUrl    string = "/storage/fcsan/storagegroup/addLun"
	RemoveLunUrl string = "/storage/fcsan/storagegroup/removeLun"
	ListLunUrl   string = "/storage/fcsan/lun/listLun"
)
