package utils

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
