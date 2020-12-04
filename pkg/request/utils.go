package request

type Logger struct {
	RequestId    string `json:"requestId"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	Request      string `json:"request"`
	Response     string `json:"response"`
	ResponseCode string `json:"responseCode"`
	RequestTime  string `json:"requestTime"`
	ResponseTime string `json:"responseTime"`
	Cost         string `json:"cost"`
}

type BaseClient struct {
	Endpoint       string
	XServiceSource string
	IsMock         bool
	LogCallback    func(logger Logger)
}

func (c *BaseClient) Setting(endpoint string, isMock bool, sourceService string, logBC func(logger Logger)) {
	c.Endpoint = endpoint
	c.IsMock = isMock
	c.XServiceSource = sourceService
	c.LogCallback = logBC
}

const (
	REQUEST_ID string = "X-Request-Id"
)

type HTTPRequest struct {
	URL    string
	Method string
	Header map[string]string
	Data   map[string]string
}

type THTTPRequest struct {
	URL    string
	Method string
	Header map[string]string
	Data   interface{}
}

type BadResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
