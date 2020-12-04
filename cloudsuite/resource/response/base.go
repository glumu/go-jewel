package responseData

type BaseResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}
