package identity

type GetProjectRequest struct {
	Id string `json:"id"`
}

type CreateProjectRequest struct {
	CreatorId      string    `json:"creatorId"`
	Description    string    `json:"description"`
	ExpirationTime string    `json:"expirationTime"`
	Name           string    `json:"name"`
	Quota          QuotaData `json:"quota"`
	Users          []Users   `json:"users"`
	VdcId          string    `json:"vdcId"`
	ZoneIds        []string  `json:"zoneIds"`
	Status         string    `json:"status"`
}

type DeleteProjectRequest struct {
	Ids []string `json:"ids"`
}

type Quota struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type ReserveQuotaRequest struct {
	OwnerId string  `json:"ownerId"`
	Quotas  []Quota `json:"quotas"`
}

type ListQutaRequest struct {
	OwnerIds []string `json:"ownerIds"`
	Type     string   `json:"type"`
}

type LoginRequest struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type LogoutRequest struct {
	UserId string `json:"userId"`
}

type QuotaData struct {
	BareMachine  string `json:"bareMachine"`
	Cpu          string `json:"cpu"`
	Disk         string `json:"disk"`
	ElasticIp    string `json:"elasticIp"`
	Loadbalancer string `json:"loadbalancer"`
	Memory       string `json:"memory"`
	NetworkIp    string `json:"networkIp"`
}

type Users struct {
	UserId   string `json:"userId"`
	UserType string `json:"userType"`
}

type UpdateProjectExpirationTimeRequest struct {
	Id             string `json:"id"`
	ExpirationTime string `json:"expirationTime"`
}

type GetUserRequest struct {
	Id             string `json:"id"`
}
