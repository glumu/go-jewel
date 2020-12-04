package requestData

type GetVMRequest struct {
	Id string `json:"id"`
}

type ResizeFlavorRequest struct {
	Id     string `json:"id"`
	Cpu    int    `json:"cpu"`
	Memory int    `json:"memory"`
}
