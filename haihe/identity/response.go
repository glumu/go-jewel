package identity

type Manager struct {
	Id 			string		`json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Telephone	string		`json:"telephone"`
}

type Detail struct {
	Name 		string 		`json:"name"`
	Managers	[]Manager	`json:"managers"`
}

type GetProjectResponse struct {
	RequestId string `json:"requestId"`
	Data      Detail `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type ReserveQuotaData struct {
	ReserveId string `json:"reserveId"`
}

type ReserveQuotaResponse struct {
	RequestId string           `json:"requestId"`
	Data      ReserveQuotaData `json:"data"`
	TimeStamp int              `json:"timestamp"`
}

type ListQuotaData struct {
	Id       string `json:"id"`
	Limit    int    `json:"limit"`
	Reserved int    `json:"reserved"`
	OwnerId  string `json:"ownerId"`
	Used     int    `json:"used"`
	Type     string `json:"type"`
}

type ListQuotaResponse struct {
	RequestId string          `json:"requestId"`
	Data      []ListQuotaData `json:"data"`
	TimeStamp int             `json:"timestamp"`
}

type Login struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

type LoginResponse struct {
	RequestId string `json:"requestId"`
	Data      Login  `json:"data"`
	TimeStamp int    `json:"timestamp"`
}

type LogoutResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	RequestId string  `json:"requestId"`
	Data      Project `json:"data"`
	TimeStamp int     `json:"timestamp"`
}

type DeleteProjectResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type UpdateProjectExpirationTimeResponse struct {
	RequestId string `json:"requestId"`
	TimeStamp int    `json:"timestamp"`
}

type UserDetail struct {
	Id   		string 		`json:"id"`
	Name		string 		`json:"name"`
	Type		string		`json:"type"`
	Status		string		`json:"status"`
	Email		string		`json:"email"`
	Telephone	string		`json:"telephone"`
}

type GetUserResponse struct {
	RequestId string          `json:"requestId"`
	Data      UserDetail	  `json:"data"`
	TimeStamp int             `json:"timestamp"`
}
