package platform

type BaseResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type ListAvailableZoneResponse struct {
	BaseResponse
	Data []struct {
		Id    string `json:"id"`
		Ref   string `json:"ref"`
		Name  string `json:"name"`
		Type  string `json:"type"` // ArNet, ArSdn
		Zones []struct {
			Id              string `json:"id"`
			Ref             string `json:"ref"`
			Name            string `json:"name"`
			ComponentCode   string `json:"componentCode"` //ArcherOS, VMware, BareMetal
			Status          string `json:"status"`        // Down, Up
			availableZoneId string `json:"availableZoneId"`
			Version         string `json:"version"`
		} `json:"zones"`
	} `json:"data"`
}
