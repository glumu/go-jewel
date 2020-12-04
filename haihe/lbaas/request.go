package lbaas

type UpdateLbRelatedResourceRequest struct {
	Id              string `json:"id"`
	RelatedResource string `json:"relatedResource"`
}

type ListLbBackendRequest struct {
	LoadbalancerId string `json:"loadbalancerId"`
}

type GetLbBackendRequest struct {
	Id string `json:"id"`
}

type CreateLbRequest struct {
	EnvironmentId string `json:"environmentId"`
	Ip            string `json:"ip"`
	Name          string `json:"name"`
	NetworkId     string `json:"networkId"`
	OwnerId       string `json:"ownerId"`
	SubnetId      string `json:"subnetId"`
	Type          string `json:"Type"`
	ZoneId        string `json:"zoneId"`
}

type CreateListenerRequest struct {
	CertificateId  string `json:"certificateId"`
	Connection     int    `json:"connection"`
	LoadbalancerId string `json:"loadbalancerId"`
	Name           string `json:"name"`
	Port           int    `json:"port"`
	Protocol       string `json:"protocol"`
}

type HealthMonitor struct {
	Delay             int    `json:"delay"`
	ExpectedHttpCodes string `json:"expectedHttpCodes"`
	HttpMethod        string `json:"httpMethod"`
	Retries           int    `json:"retries"`
	Timeout           int    `json:"timeout"`
	Type              string `json:"type"`
	UrlPath           string `json:"urlPath"`
}

type CreateLbPollRequest struct {
	Algorithm      string        `json:"algorithm"`
	HealthMonitor  HealthMonitor `json:"healthMonitor"`
	ListenerId     string        `json:"listenerId"`
	LoadbalancerId string        `json:"loadbalancerId"`
	Name           string        `json:"name"`
	Protocol       string        `json:"protocol"`
}
