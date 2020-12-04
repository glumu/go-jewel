package identity

import "github.com/go-jewel/haihe/uitls"

// ou
type IdentityClient struct {
	utils.BaseClient
	MockApi []string
}

var (
	GetProjectUrl                  string = "/getProject"
	CreateProjectUrl               string = "/createProject"
	DeleteProjectUrl               string = "/deleteProject"
	ReserveQuotaUrl                string = "/reserveQuota"
	ListQuotaUrl                   string = "/listQuota"
	LoginUrl                       string = "/login"
	LogoutUrl                      string = "/logout"
	UpdateProjectExpirationTimeUrl string = "/updateProjectExpirationTime"
	GetUserUrl					   string = "/getUser"
)
