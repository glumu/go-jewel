package resource_scheduler

type ListQuotaRequest struct {
	OwnerIds []string `json:"ownerIds"`
}

type ReserveQuotaRequest struct {
	OwnerId   string           `json:"ownerId"`
	Quotas    []map[string]int `json:"quotas"` // type: CPU, Memory, Disk, NetworkIp, ElasticIp，GPU
	action    string           `json:"action"`
	component string           `json:"component"`
}
