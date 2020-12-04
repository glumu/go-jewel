package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ErrorRequest struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Trove400Response struct {
	BadRequest ErrorRequest `json:"badRequest"`
}

type Trove401Response struct {
	Unauthorized ErrorRequest `json:"unauthorized"`
}

type Trove403Response struct {
	Forbidden ErrorRequest `json:"forbidden"`
}

type Trove404Response struct {
	ItemNotFound ErrorRequest `json:"itemNotFound"`
}

type Trove405Response struct {
	BadMethod ErrorRequest `json:"badMethod"`
}

type Trove413Response struct {
	OverLimit ErrorRequest `json:"overLimit"`
}

type Trove422Response struct {
	UnprocessableEntity ErrorRequest `json:"unprocessableEntity"`
}

type Trove500Response struct {
	InstanceFault ErrorRequest `json:"instanceFault"`
}

type Trove501Response struct {
	NotImplemented ErrorRequest `json:"notImplemented"`
}

type Trove503Response struct {
	ServiceUnavailable ErrorRequest `json:"serviceUnavailable"`
}

func TroveRequest(url string, method string, parameters map[string]string, data interface{},
	head map[string]string, timeout time.Duration, expect int) ([]byte, error) {
	req, err := TNewRequest(url, method, parameters, data, head)
	if err != nil {
		return nil, err
	}

	resp, err := doTroveRequest(req, timeout, expect)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func doTroveRequest(request *THTTPRequest, timeout time.Duration, expect int) ([]byte, error) {
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

	defer res.Body.Close()

	var by []byte
	by, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read body error, status code: %d", res.StatusCode))
	}
	if res.StatusCode != expect {
		switch res.StatusCode {
		case 400:
			var response Trove400Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.BadRequest.Message != "" {
				return nil, errors.New(fmt.Sprintf("bad request, %s", response.BadRequest.Message))
			}
		case 401:
			var response Trove401Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.Unauthorized.Message != "" {
				return nil, errors.New(fmt.Sprintf("unauthorized, %s", response.Unauthorized.Message))
			}
		case 403:
			var response Trove403Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.Forbidden.Message != "" {
				return nil, errors.New(fmt.Sprintf("forbidden, %s", response.Forbidden.Message))
			}
		case 404:
			var response Trove404Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.ItemNotFound.Message != "" {
				return nil, errors.New(fmt.Sprintf("item not found, %s", response.ItemNotFound.Message))
			}
		case 405:
			var response Trove405Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.BadMethod.Message != "" {
				return nil, errors.New(fmt.Sprintf("bad method, %s", response.BadMethod.Message))
			}
		case 413:
			var response Trove413Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.OverLimit.Message != "" {
				return nil, errors.New(fmt.Sprintf("over limit, %s", response.OverLimit.Message))
			}
		case 422:
			var response Trove422Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.UnprocessableEntity.Message != "" {
				return nil, errors.New(fmt.Sprintf("unporcessable entity, %s", response.UnprocessableEntity.Message))
			}
		case 500:
			var response Trove500Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.InstanceFault.Message != "" {
				return nil, errors.New(fmt.Sprintf("instance fault, %s", response.InstanceFault.Message))
			}
		case 501:
			var response Trove501Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.NotImplemented.Message != "" {
				return nil, errors.New(fmt.Sprintf("not implemented, %s", response.NotImplemented.Message))
			}
		case 503:
			var response Trove503Response
			err = json.Unmarshal(by, &response)
			if err != nil {
				return nil, errors.New("status bad request, json unmarshal error.")
			}
			if response.ServiceUnavailable.Message != "" {
				return nil, errors.New(fmt.Sprintf("service unavailable, %s", response.ServiceUnavailable.Message))
			}
		default:
			return nil, errors.New(fmt.Sprintf("unknown error, status code: %d", res.StatusCode))
		}
	}
	return by, nil
}
