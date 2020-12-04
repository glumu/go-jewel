package configmanager

type Detail struct {
	Name 			string `json:"name"`
	Type 			string `json:"type"`
	Value 			string `json:"value"`
	Description 	string `json:"description"`
}

type GetConfigResponse struct {
	RequestId string `json:"requestId"`
	Data      Detail `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

