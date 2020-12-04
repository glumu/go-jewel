package notification

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *NotificationClient) SendEmail(request SendEmailRequest, requestId string, timeout time.Duration) (*SendEmailResponse, error) {
	fullUrl := c.Endpoint + SendEmailUrl

	var sendEmailResponse SendEmailResponse
	if c.IsMock {
		data := []EmailResponseData{
			{ReceiverId: "id1", Email: "admin@huayun.com"},
		}
		sendEmailResponse = SendEmailResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &sendEmailResponse)
		if err != nil {
			return nil, err
		}
	}

	return &sendEmailResponse, nil
}

func (c *NotificationClient) SendSMS(request SendSMSRequest, requestId string, timeout time.Duration) (*SendSMSResponse, error) {
	fullUrl := c.Endpoint + SendSMSUrl

	var sendSMSResponse SendSMSResponse
	if c.IsMock {
		data := []string{
			"mock",
		}
		sendSMSResponse = SendSMSResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &sendSMSResponse)
		if err != nil {
			return nil, err
		}
	}

	return &sendSMSResponse, nil
}

func (c *NotificationClient) CheckEmail(request CheckEmailRequest, requestId string, timeout time.Duration) (*CheckEmailResponse, error) {
	fullUrl := c.Endpoint + CheckEmailUrl

	var checkEmailResponse CheckEmailResponse
	if c.IsMock {
		data := CheckResult{
			Result: "success",
		}
		checkEmailResponse = CheckEmailResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &checkEmailResponse)
		if err != nil {
			return nil, err
		}
	}

	return &checkEmailResponse, nil
}

func (c *NotificationClient) CheckSMS(request CheckSMSRequest, requestId string, timeout time.Duration) (*CheckSMSResponse, error) {
	fullUrl := c.Endpoint + CheckSMSUrl

	var checkSMSResponse CheckSMSResponse
	if c.IsMock {
		data := CheckResult{
			Result: "success",
		}
		checkSMSResponse = CheckSMSResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &checkSMSResponse)
		if err != nil {
			return nil, err
		}
	}

	return &checkSMSResponse, nil
}

func (c *NotificationClient) ListRecord(request ListRecordRequest, requestId string, sessionId string, timeout time.Duration) (*ListRecordResponse, error) {
	fullUrl := c.Endpoint + ListRecordUrl

	var listRecordResponse ListRecordResponse
	if c.IsMock {
		listRecordResponse = ListRecordResponse{
			RequestId: requestId,
			Data: []RecordResponse{
				{
					ReceiverId:  "receiver-id",
					SendWay:     "Email",
					ContentType: "",
					SendState:   "success",
				},
			},
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &listRecordResponse)
		if err != nil {
			return nil, err
		}
	}
	return &listRecordResponse, nil
}
