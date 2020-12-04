package notification

import "github.com/glumu/go-jewel/haihe/uitls"

type NotificationClient struct {
	utils.BaseClient
}

var (
	SendEmailUrl 	string = "/sendEmailBatch"
	SendSMSUrl		string = "/sendSMSBatch"
	CheckEmailUrl	string = "/checkEmail"
	CheckSMSUrl		string = "/checkSMS"
	ListRecordUrl	string = "/listRecord"
)