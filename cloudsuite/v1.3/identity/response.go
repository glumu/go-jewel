package identity

type BaseResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type _vdcData struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	IsDeleted   int      `json:"isDeleted"`
	RegionIds   []string `json:"regionIds"`
	Aggregation struct {
		Project []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			IsDeleted int    `json:"isDeleted"`
		} `json:"project"`
		AvailableZoneIds []string `json:"availableZoneIds"`
	} `json:"aggregation"`
}

type GetVdcResponse struct {
	BaseResponse
	Data _vdcData `json:"data"`
}

type ListVdcResponse struct {
	BaseResponse
	Data []_vdcData `json:"data"`
}

type _projectData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted int    `json:"isDeleted"`
	VdcId     string `json:"vdcId"`
	VdcName   string `json:"vdcName"`
}

type GetProjectResponse struct {
	BaseResponse
	Data _projectData `json:"data"`
}

type ListMyProjectResponse struct {
	BaseResponse
	Data []_projectData `json:"data"`
}
