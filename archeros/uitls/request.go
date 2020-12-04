package utils

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
