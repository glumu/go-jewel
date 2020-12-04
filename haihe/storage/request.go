package storage

type AddLunRequest struct {
	Wwpns  []string `json:"wwpns"`
	ZoneId string   `json:"zone_id"`
	//StorageGroup string   `json:"storage_group"`
	//LunId   string `json:"lun_id"`
	LunNumber string `json:"lun_number"`
	FcsanId   string `json:"fcsan_id"`
}

type RemoveLunRequest struct {
	//StorageGroup string `json:"storage_group"`
	Wwpns   []string `json:"wwpns"`
	LunId   string   `json:"lun_id"`
	FcsanId string   `json:"fcsan_id"`
	ZoneId  string   `json:"zone_id"`
	Delete  bool     `json:"delete"`
}

type ListLunRequest struct {
	LunIds []string `json:"lun_id"`
}
