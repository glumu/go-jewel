package resource_scheduler

type BaseResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type ListQuotaResponse struct {
	BaseResponse
	Data struct {
		Limit    int    `json:"limit"`
		Max      int    `json:"max"`
		Min      int    `json:"min"`
		Reserved int    `json:"reserved"`
		Type     string `json:"type"`
		Used     int    `json:"used"`
	} `json:"data"`
}

type ReserveQuotaResponse struct {
	BaseResponse
	Data struct {
		ReserveId string `json:"reserveId"`
	} `json:"data"`
}
