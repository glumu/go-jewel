package monitor

import (
	"encoding/json"
	"time"

	"github.com/glumu/go-jewel/haihe/uitls"
)

func (c *MonitorClient) CreateGroup(request CreateGroupRequest, requestId string, timeout time.Duration) (*CreateGroupResponse, error) {
	fullUrl := c.Endpoint + CreateGroupUrl

	var response CreateGroupResponse
	if c.IsMock {
		data := GroupDataResponse{
			Id:   utils.GenerateRandnum(),
			Name: "test",
		}
		response = CreateGroupResponse{
			Code: 0,
			Data: data,
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

func (c *MonitorClient) DeleteGroup(request DeleteGroupRequest, requestId string, timeout time.Duration) (*DeleteGroupResponse, error) {
	fullUrl := c.Endpoint + DeleteGroupUrl

	var response DeleteGroupResponse
	if c.IsMock {
		data := GroupDataResponse{
			Id:   utils.GenerateRandnum(),
			Name: "test",
		}
		response = DeleteGroupResponse{
			Code: 0,
			Data: data,
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

func (c *MonitorClient) UpdateGroup(request UpdateGroupRequest, requestId string, timeout time.Duration) (*UpdateGroupResponse, error) {
	fullUrl := c.Endpoint + UpdateGroupUrl

	var response UpdateGroupResponse
	if c.IsMock {
		data := GroupDataResponse{
			Id:   utils.GenerateRandnum(),
			Name: "test",
		}
		response = UpdateGroupResponse{
			Code: 0,
			Data: data,
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

func (c *MonitorClient) BindGroup(request BindGroupRequest, requestId string, timeout time.Duration) (*BindGroupResponse, error) {
	fullUrl := c.Endpoint + BindHostToGroupUrl

	var response BindGroupResponse
	if c.IsMock {
		response = BindGroupResponse{
			Code: 0,
			Data: "bind success",
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

func (c *MonitorClient) UnbindGroup(request UnbindGroupRequest, requestId string, timeout time.Duration) (*UnbindGroupResponse, error) {
	fullUrl := c.Endpoint + UnBindHostToGroupUrl

	var response UnbindGroupResponse
	if c.IsMock {
		response = UnbindGroupResponse{
			Code: 0,
			Data: "unbind success",
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

func (c *MonitorClient) SyncHostToGroup(request SyncHostToGroupRequest, requestId string, timeout time.Duration) (*SyncHostToGroupResponse, error) {
	fullUrl := c.Endpoint + SyncHostToGroupUrl

	var response SyncHostToGroupResponse
	if c.IsMock {
		response = SyncHostToGroupResponse{
			Code: 0,
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
