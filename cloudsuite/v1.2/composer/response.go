package composer

type BaseResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type ListAvailableZoneResponse struct {
	BaseResponse
	Data []struct {
		Id           string   `json:"id"`
		Name         string   `json:"name"`
		IsDeleted    bool     `json:"isDeleted"`
		Type         string   `json:"type"` // ArNet, ArSdn
		NetworkType  []string `json:"networkType"`
		LogicalZones []struct {
			Id            string `json:"id"`
			Name          string `json:"name"`
			ComponentCode string `json:"componentCode"`
			Arch          string `json:"arch"`
			Version       string `json:"version"`
			Description   string `json:"description"`
			Zones         []struct {
				Id              string `json:"id"`
				Ref             string `json:"ref"`
				Name            string `json:"name"`
				ComponentCode   string `json:"componentCode"` //ArcherOS, VMware, BareMetal
				Status          string `json:"status"`        // Down, Up
				availableZoneId string `json:"availableZoneId"`
				Version         string `json:"version"`
			} `json:"zones"`
		} `json:"logicalZones"`
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
