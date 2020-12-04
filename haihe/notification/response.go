package notification

type EmailResponseData struct {
	ReceiverId		string					`json:"receiverId"`
	Email 			string					`json:"email"`
}

type SendEmailResponse struct {
	RequestId 		string 					`json:"requestId"`
	Data      		[]EmailResponseData  	`json:"data"`
}

type SendSMSResponse struct {
	RequestId 		string 		`json:"requestId"`
	Data      		[]string  	`json:"data"`
}

type CheckResult struct {
	Result			string		`json:"result"`
}

type CheckEmailResponse struct {
	RequestId 		string 			`json:"requestId"`
	Data      		CheckResult  	`json:"data"`
}

type CheckSMSResponse struct {
	RequestId 		string 			`json:"requestId"`
	Data      		CheckResult  	`json:"data"`
}

type RecordResponse struct {
	ReceiverId			string		`json:"receiverId"`
	SendWay				string		`json:"sendWay"`
	ContentType			string		`json:"contentType"`
	SendState			string		`json:"sendState"`
}

type ListRecordResponse struct {
	RequestId 		string 				`json:"requestId"`
	Data      		[]RecordResponse  	`json:"data"`
}
