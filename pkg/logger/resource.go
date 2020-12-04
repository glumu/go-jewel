package logger

import (
	"fmt"
	"time"

	"github.com/glumu/go-jewel/haihe/uitls"

	"github.com/glumu/go-jewel/haihe/identity"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// 资源日志 - 开始
type ResourceLogHeadData struct {
	CreateTime   string `json:"createTime"`                  // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType    string `json:"auditType"`                   // 业务类型：resource（资源类）
	LogType      string `json:"logType" default:"LOG_START"` // 日志类型（LOG_START）
	Action       string `json:"action"`                      // 操作action
	ServiceName  string `json:"serviceName"`
	RequestId    string `json:"requestId"`             // 操作请求ID
	Version      string `json:"version" default:"1.0"` // 版本号
	VisitorId    string `json:"visitorId"`             // 操作者ID
	CloudId      string `json:"cloudId"`               // 云环境ID
	DataCenterId string `json:"dataCenterId"`          // 数据中心ID
	OwnerId      string `json:"ownerId"`               // 项目ID
	OwnerName    string `json:"ownerName"`             // 项目名称
	ClientIP     string `json:"clientIP"`
	Level        string `json:"level" default:"Normal"`
	SourceAgent  string `json:"sourceAgent" default:"CloudSuite"`
}

// 资源日志 - 结束
type ResourceLogFootData struct {
	CreateTime   string `json:"createTime"`                // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	AuditType    string `json:"auditType"`                 // 业务类型：resource（资源类）
	LogType      string `json:"logType" default:"LOG_END"` // 日志类型（LOG_END）
	Action       string `json:"action"`                    // 操作action
	ServiceName  string `json:"serviceName"`
	RequestId    string `json:"requestId"`             // 操作请求ID
	Version      string `json:"version" default:"1.0"` // 版本号
	VisitorId    string `json:"visitorId"`             // 操作者ID
	CloudId      string `json:"cloudId"`               // 云环境ID
	DataCenterId string `json:"dataCenterId"`          // 数据中心ID
	OwnerId      string `json:"ownerId"`               // 项目ID
	OwnerName    string `json:"ownerName"`             // 项目名称
	Resources    string `json:"resources"`             // 资源列表（json字符串）
	Result       string `json:"result"`                // 操作结果（failure、success）
	ClientIP     string `json:"clientIP"`
	Level        string `json:"level" default:"Normal"`
	SourceAgent  string `json:"sourceAgent" default:"CloudSuite"`
}

// 日志 - 详情（资源、业务）
type ResourceLogContentData struct {
	CreateTime   string `json:"createTime"`                   // 操作时间（精确到毫秒） yyyy-MM-DD HH:mm:ss.S
	LogType      string `json:"logType" default:"LOG_DETAIL"` // 日志类型（LOG_DETAIL）
	Action       string `json:"action"`                       // 每一步详情的操作动作（由各服务自定义，语义清晰即可， 如create user、get project information等）
	ServiceName  string `json:"serviceName"`                  // 日志来源服务（注册中心的服务名），谁记录的日志
	RequestId    string `json:"requestId"`                    // 操作请求ID
	Version      string `json:"version" default:"1.0"`        // 版本号
	Params       string `json:"params"`                       // 当前详情操作涉及的参数（json字符串，由各服务自定义，数据要方便定位错误）
	ErrorCode    string `json:"errorCode"`                    // 错误码
	ErrorMessage string `json:"errorMessage"`                 // 错误信息
	Result       string `json:"result"`                       // 操作结果（failure、success）
}

type ResourceDetail struct {
	Id   string `json:"id"`   // 项目ID
	Name string `json:"name"` // 资源名称
	Type string `json:"type"` // 资源类型
}

type ProjectName struct {
	Id        string
	Name      string
	CreatedAt time.Time
}

type ResourceLogger struct {
	logger         *logrus.Logger
	fl             *rotatelogs.RotateLogs // 日志切割
	identityClient *identity.IdentityClient
	projectNames   []ProjectName
}

func NewResourceLogger(path string, identityEndpoint string) (*ResourceLogger, error) {
	fl, err := rotatelogs.New(path)
	//// 日志切割(每天切、保留7天)
	//fl, err := rotatelogs.New(
	//	fmt.Sprintf("%s.%%Y%%m%%d", path),
	//	rotatelogs.WithLinkName(path),
	//	rotatelogs.WithRotationTime(24*time.Hour), // default: 86400 sec
	//	rotatelogs.WithMaxAge(7*24*time.Hour),     // default: 7 days
	//)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.SetOutput(fl)
	logger.SetFormatter(&logrus.JSONFormatter{})

	var identityClient *identity.IdentityClient
	if identityEndpoint != "" {
		identityClient = &identity.IdentityClient{
			BaseClient: utils.BaseClient{
				Endpoint:    identityEndpoint,
				IsMock:      false,
				LogCallback: nil,
			},
		}
	}

	return &ResourceLogger{
		logger:         logger,
		fl:             fl,
		identityClient: identityClient,
		projectNames:   []ProjectName{},
	}, nil
}

func (a *ResourceLogger) Close() {
	if a != nil && a.fl != nil {
		a.fl.Close()
	}
}

func (a *ResourceLogger) Record(logInfo interface{}) {
	logObj := a.logger.WithFields(Struct2Map(logInfo))
	logObj.Info("")
}

func (a *ResourceLogger) RecordHead(head ResourceLogHeadData) {
	if a.identityClient != nil && head.OwnerName == "" && head.OwnerId != "" {
		name, err := a.getProjectName(head.OwnerId, head.RequestId)
		if err == nil {
			head.OwnerName = name
		}
	}
	a.Record(head)
}

func (a *ResourceLogger) RecordFoot(foot ResourceLogFootData) {
	if a.identityClient != nil && foot.OwnerName == "" && foot.OwnerId != "" {
		name, err := a.getProjectName(foot.OwnerId, foot.RequestId)
		if err == nil {
			foot.OwnerName = name
		}
	}
	a.Record(foot)
}

func (a *ResourceLogger) getProjectName(id string, requestId string) (string, error) {
	if a.identityClient != nil {
		index := -1
		for idx, val := range a.projectNames {
			if val.Id == id {
				index = idx
				s, err := time.ParseDuration(fmt.Sprintf("%dm", 60)) // cache 60 min
				if err != nil {
					return "", err
				}
				time1 := val.CreatedAt.Add(s)
				if time1.After(time.Now()) {
					return val.Name, nil
				}
				break
			}
		}
		getProjectRequest := identity.GetProjectRequest{
			Id: id,
		}
		getResponse, err := a.identityClient.GetProject(getProjectRequest, requestId, 14*time.Second)
		if err == nil {
			newPN := ProjectName{
				Id:        id,
				Name:      getResponse.Data.Name,
				CreatedAt: time.Now(),
			}
			if index != -1 {
				a.projectNames = append(a.projectNames[:index], a.projectNames[index+1:]...)
			}
			a.projectNames = append(a.projectNames, newPN)
			return getResponse.Data.Name, nil
		}
	}

	return "", nil
}
