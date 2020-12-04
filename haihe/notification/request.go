package notification


type EmailReceiver struct {
	ReceiverId			string		`json:"receiverId"`
	Email				string		`json:"email"`
}

type SendEmailRequest struct {
	Service 			string				`json:"service"`
	TemplateCode		string           	`json:"templateCode"`
	TemplateParams		[]string       		`json:"templateParams"`
	Receivers			[]EmailReceiver 	`json:"receivers"`
	Title				string          	`json:"title"`
}

type SMSReceiver struct {
	UserId				string				`json:"userId"`
	Mobile				string				`json:"mobile"`
}

type SendSMSRequest struct {
	Service 			string				`json:"service"`
	ContentType			string        		`json:"contentType"`
	TemplateCode		string           	`json:"templateCode"`
	TemplateParams		[]string       		`json:"templateParams"`
	ReceiverType		string				`json:"receiverType"`
	Receivers			[]SMSReceiver		`json:"receivers"`
	ResourceInfo		string           	`json:"resourceInfo"`
	ReceiverIP			string         		`json:"receiverIP"`
	FrequencyLimit		bool           		`json:"frequencyLimit"`
	Sign				string				`json:"sign"`
}

type CheckEmailRequest struct {
	SmtpAdress 			string				`json:"smtpAdress"`
	SmtpUser 			string				`json:"smtpUser"`
	SmtpPwd 			string				`json:"smtpPwd"`
	ReceiverEmail 		string				`json:"receiverEmail"`
}

type CheckSMSRequest struct {
	AppId				string				`json:"appId"`
	AppSecurity			string				`json:"appSecurity"`
	Url					string				`json:"url"`
	ReceiverNo			string				`json:"receiverNo"`
}

type ListRecordRequest struct {
	Title				string				`json:"title"`
	SendWay				string				`json:"sendWay"`
	ReceiverId			string				`json:"receiverId"`
}

