package autoscaling

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateGroupResponse struct {
	RequestId string `json:"requestId"`
	Data      Group  `json:"data"`
}

type Configuration struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateConfigurationResponse struct {
	RequestId string        `json:"requestId"`
	Data      Configuration `json:"data"`
}

type Policy struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreatePolicyResponse struct {
	RequestId string `json:"requestId"`
	Data      Policy `json:"data"`
}

type DeleteGroupResponse struct {
	RequestId string `json:"requestId"`
}

type DeleteConfigurationResponse struct {
	RequestId string `json:"requestId"`
}

type DeletePolicyResponse struct {
	RequestId string `json:"requestId"`
}

type ScaleHostInstanceResponse struct {
	RequestId string `json:"requestId"`
	Data      Group  `json:"data"`
}

type ElasticLog struct {
	Id string `json:"id"`
}

type ListElasticLogResponse struct {
	RequestId string       `json:"requestId"`
	Data      []ElasticLog `json:"data"`
}

type GetElasticLogResponse struct {
	RequestId string     `json:"requestId"`
	Data      ElasticLog `json:"data"`
}
