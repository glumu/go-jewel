package resource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

func (c *ResourceClient) CreateInterface(request CreateInterfaceRequest, requestId string, timeout time.Duration) (*CreateInterfaceResponse, error) {
	fullUrl := c.Endpoint + CreateInterfaceUrl

	var response CreateInterfaceResponse
	if c.IsMock {
		port := Port{
			Id:         utils.GenMockId(),
			Ref:        utils.GenMockId(),
			NetworkId:  utils.GenMockId(),
			SubnetId:   utils.GenMockId(),
			IpAddress:  "192.168.0.1",
			MacAddress: "00:e0:ed:57:18:20",
		}
		response = CreateInterfaceResponse{
			RequestId: requestId,
			Data:      port,
			TimeStamp: 1111111111,
			Ref:       "xxxxxx",
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &response)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (c *ResourceClient) DeleteInterface(request DeleteInterfaceRequest, requestId string, timeout time.Duration) (*DeleteInterfaceResponse, error) {
	fullUrl := c.Endpoint + DeleteInterfaceUrl

	var response DeleteInterfaceResponse
	if c.IsMock {
		response = DeleteInterfaceResponse{}
	} else {
		head := map[string]string{
			"X-Request-Id": requestId,
		}
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &response)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (c *ResourceClient) ListAllPort(request ListAllPortRequest, requestId string, timeout time.Duration) (*ListAllPortResponse, error) {
	fullUrl := c.Endpoint + ListAllPortUrl

	var response ListAllPortResponse
	if c.IsMock {
		data := InterfaceDetail{
			Id:        utils.GenMockId(),
			Ip:        "192.168.0.1",
			NetworkId: utils.GenMockId(),
		}
		response = ListAllPortResponse{
			RequestId: requestId,
			Data:      []InterfaceDetail{data},
			TimeStamp: 1111111111,
		}
	} else {
		head := map[string]string{
			"X-Request-Id": requestId,
		}
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &response)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (c *ResourceClient) CreateFireWall(request CreateFirewallRequest, requestId string, sessionId string, timeout time.Duration) (*CreateFirewallResponse, error) {
	fullUrl := c.Endpoint + CreateFirewallUrl

	var response CreateFirewallResponse
	if c.IsMock {
		response = CreateFirewallResponse{
			RequestId: requestId,
			Data:      "",
			TimeStamp: 123,
		}
	} else {
		head := map[string]string{
			"X-Request-Id": requestId,
			"Cookie":       fmt.Sprintf("sessionId=" + sessionId),
			"X-Scope":      "System",
		}
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &response)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}

func (c *ResourceClient) DeleteFirewall(request DeleteFirewallRequest, requestId string, sessionId string, timeout time.Duration) (*DeleteFirewallResponse, error) {
	fullUrl := c.Endpoint + DeleteFirewallUrl

	var response DeleteFirewallResponse
	if c.IsMock {
		response = DeleteFirewallResponse{
			RequestId: requestId,
			TimeStamp: 123,
		}
	} else {
		head := map[string]string{
			"X-Request-Id": requestId,
			"Cookie":       fmt.Sprintf("sessionId=" + sessionId),
			"X-Scope":      "System",
		}
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &response)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}
