package lbaas

type Data struct {
	Id string `json:"id"`
}

type UpdateLbRelatedResourceResponse struct {
	RequestId string `json:"requestId"`
}

type LbBackend struct {
	BackendId    string `json:"backendId"`
	BackendName  string `json:"backendName"`
	BackendType  string `json:"backendType"`
	HealthStatus string `json:"healthStatus"`
}

type ListLbBackendResponse struct {
	RequestId string      `json:"requestId"`
	Data      []LbBackend `json:"data"`
}

type GetLbBackendResponse struct {
	RequestId string    `json:"requestId"`
	Data      LbBackend `json:"data"`
}

type CreateLbResponse struct {
	RequestId string `json:"requestId"`
	Data      Data   `json:"data"`
}

type CreateListenerResponse struct {
	RequestId string `json:"requestId"`
	Data      Data   `json:"data"`
}

type CreateLbPollResponse struct {
	RequestId string `json:"requestId"`
	Data      Data   `json:"data"`
}
