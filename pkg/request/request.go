package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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

func TRequest(url string, method string, parameters map[string]string, data interface{},
	head map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := TNewRequest(url, method, parameters, data, head)
	if err != nil {
		return nil, err
	}

	resp, err := doTRequest(req, timeout)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func TNewRequest(fullUrl string, method string, parameters map[string]string, data interface{}, head map[string]string) (*THTTPRequest, error) {
	var q []string
	for k, v := range parameters {
		if v == "" {
			continue
		}
		q = append(q, url.QueryEscape(k))
		q = append(q, "=")
		q = append(q, url.QueryEscape(v))
		q = append(q, "&")
	}
	if len(q) > 0 {
		q = q[:(len(q) - 1)]
	}
	fullUrl += strings.Join(q, "")

	req := &THTTPRequest{
		URL:    fullUrl,
		Method: method,
		Data:   data,
		Header: head,
	}
	return req, nil
}

func doTRequest(request *THTTPRequest, timeout time.Duration) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	var req *http.Request
	var err error
	if request.Data != nil {
		rendered, err := json.Marshal(request.Data)
		if err != nil {
			return nil, err
		}

		body := bytes.NewReader(rendered)
		req, err = http.NewRequest(request.Method, request.URL, body)
	} else {
		req, err = http.NewRequest(request.Method, request.URL, nil)
	}

	if err != nil {
		return nil, err
	}
	for k, v := range request.Header {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var by []byte
	if res.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("internal server error")
	} else if res.StatusCode == http.StatusBadRequest {
		by, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("status bad request, read body error.")
		}
		var response BadResponse
		err = json.Unmarshal(by, &response)
		if err != nil {
			return nil, errors.New("status bad request, json unmarshal error.")
		}
		if response.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprintf("bad request, %s", response.ErrorMessage))
		}
		return nil, errors.New("status bad request")
	} else if res.StatusCode == http.StatusNotFound {
		by, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("status 404 bad request, read body error.")
		}
		var response BadResponse
		err = json.Unmarshal(by, &response)
		if err != nil {
			return nil, errors.New("status 404 bad request, json unmarshal error.")
		}
		if response.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprintf("status 404 bad request, %s", response.ErrorMessage))
		}
		return nil, errors.New("status 404 bad request")
	}
	defer res.Body.Close()

	by, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return by, nil
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