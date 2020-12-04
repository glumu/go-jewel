package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type BaseClient struct {
	Endpoint       string
	XServiceSource string
	IsMock         bool
	LogCallback    func(logger Logger)
}

func (c *BaseClient) GenHeaders(requestId string, headers [][2]string) map[string]string {
	head := map[string]string{
		"Content-Type": `application/json;charset=UTF-8`,
		"X-Request-Id": requestId,
	}
	if c.XServiceSource != "" {
		head["X-Service-Source"] = c.XServiceSource
	}
	if headers != nil {
		for _, header := range headers {
			if header[1] != "" {
				head[header[0]] = header[1]
			}
		}
	}
	return head
}

func (c *BaseClient) Request(url string, method string, parameters map[string]string, data interface{},
	head map[string]string, timeout time.Duration) ([]byte, error) {

	if head == nil {
		head = map[string]string{
			"Content-Type": `application/json;charset=UTF-8`,
		}
	} else {
		if _, ok := head["Content-Type"]; ok == false {
			head["Content-Type"] = "application/json;charset=UTF-8"
		}
	}
	requestId := "none"
	if _, ok := head[REQUEST_ID]; ok == true {
		requestId = head[REQUEST_ID]
	}
	start := time.Now()
	reqStr := ""
	reqJson, err := json.Marshal(data)
	if err == nil {
		reqStr = string(reqJson)
	}
	c.Start(start, url, reqStr, requestId)
	resp, err := TRequest(url, method, parameters, data, head, timeout)
	c.End(start, url, reqStr, requestId, resp, err)
	return resp, err
}

func (c *BaseClient) TroveRequest(url string, method string, parameters map[string]string, data interface{},
	head map[string]string, timeout time.Duration, expect int) ([]byte, error) {

	if head == nil {
		head = map[string]string{
			"Content-Type": `application/json;charset=UTF-8`,
		}
	} else {
		if _, ok := head["Content-Type"]; ok == false {
			head["Content-Type"] = "application/json;charset=UTF-8"
		}
	}
	requestId := "none"
	if _, ok := head[REQUEST_ID]; ok == true {
		requestId = head[REQUEST_ID]
	}
	start := time.Now()
	reqStr := ""
	reqJson, err := json.Marshal(data)
	if err == nil {
		reqStr = string(reqJson)
	}
	c.Start(start, url, reqStr, requestId)
	resp, err := TroveRequest(url, method, parameters, data, head, timeout, expect)
	c.End(start, url, reqStr, requestId, resp, err)
	return resp, err
}

func (c *BaseClient) Start(t time.Time, url string, reqStr string, requestId string) {
	if c.LogCallback != nil {
		logger := Logger{
			RequestId:    requestId,
			Type:         "request",
			Url:          url,
			Request:      reqStr,
			Response:     "",
			ResponseCode: "",
			RequestTime:  t.Format("2006-01-02 15:04:05.000"),
			ResponseTime: "",
			Cost:         "",
		}
		c.LogCallback(logger)
	}
}

func (c *BaseClient) End(start time.Time, url string, reqStr string, requestId string, resp []byte, err error) {
	if c.LogCallback != nil {
		end := time.Now()
		latency := end.Sub(start)

		var respStr string
		var respCode string
		if err == nil {
			respCode = "200"
			respStr = string(resp)
		} else {
			respCode = "400"
			respStr = err.Error()
		}
		logger := Logger{
			RequestId:    requestId,
			Type:         "response",
			Url:          url,
			Request:      reqStr,
			Response:     respStr,
			ResponseCode: respCode,
			RequestTime:  start.Format("2006-01-02 15:04:05.000"),
			ResponseTime: end.Format("2006-01-02 15:04:05.000"),
			Cost:         fmt.Sprintf("%v", latency),
		}
		c.LogCallback(logger)
	}
}
