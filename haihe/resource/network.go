package resource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

func (c *ResourceClient) GetSubnet(request GetSubnetRequest, requestId string, timeout time.Duration) (*GetSubnetResponse, error) {
	fullUrl := c.Endpoint + GetSubnetUrl

	var response GetSubnetResponse
	if c.IsMock {
		data := Subnet{
			Id:      utils.GenMockId(),
			Name:    "test",
			Gateway: "10.10.10.1",
			Ref:     "xx-dsf-eer-ttt",
			Cidr:    "10.10.10.0/24",
		}
		response = GetSubnetResponse{
			RequestId: requestId,
			TimeStamp: 1111111111,
			Data:      data,
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

func (c *ResourceClient) CreateSubnet(request CreateSubnetRequest, requestId string, sessionId string, timeout time.Duration) (*CreateSubnetResponse, error) {
	fullUrl := c.Endpoint + CreateSubnetUrl

	var response CreateSubnetResponse
	if c.IsMock {
		response = CreateSubnetResponse{
			RequestId: requestId,
			TimeStamp: 1111111111,
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

func (c *ResourceClient) ListNetwork(request ListNetworkRequest, requestId string, timeout time.Duration) (*ListNetworkResponse, error) {
	fullUrl := c.Endpoint + ListNetworkUrl

	var response ListNetworkResponse
	if c.IsMock {
		subnet := Subnet{
			Id:      utils.GenMockId(),
			Name:    "test",
			Gateway: "10.10.10.1",
			Ref:     "xx-dsf-eer-ttt",
			Cidr:    "10.10.10.0/24",
		}
		data := Network{
			Id:                   utils.GenMockId(),
			Name:                 "test",
			OwnerId:              utils.GenMockId(),
			Ref:                  utils.GenMockId(),
			BaremetalNetworkType: "manage",
			Subnets:              []Subnet{subnet},
		}
		response = ListNetworkResponse{
			RequestId: requestId,
			TimeStamp: 1111111111,
			Data:      []Network{data},
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

func (c *ResourceClient) CreateNetwork(request CreateNetworkRequest, requestId string, sessionId string, timeout time.Duration) (*CreateNetworkResponse, error) {
	fullUrl := c.Endpoint + CreateNetworkUrl

	var response CreateNetworkResponse
	if c.IsMock {
		response = CreateNetworkResponse{
			RequestId: requestId,
			TimeStamp: 1111111111,
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

func (c *ResourceClient) DeleteNetwork(request DeleteNetworkRequest, requestId string, sessionId string, timeout time.Duration) (*DeleteNetworkResponse, error) {
	fullUrl := c.Endpoint + DeleteNetworkUrl

	var response DeleteNetworkResponse
	if c.IsMock {
		response = DeleteNetworkResponse{
			RequestId: requestId,
			TimeStamp: 1111111111,
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
