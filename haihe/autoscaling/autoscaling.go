package autoscaling

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

const VisitorIdHeader string = "visitorId"

func (c *AutoscalingClient) CreateGroup(request CreateGroupRequest, requestId string, sessionId string, timeout time.Duration) (*CreateGroupResponse, error) {
	fullUrl := c.Endpoint + CreateGroupUrl

	var createGroupResponse CreateGroupResponse
	if c.IsMock {
		data := Group{
			Id:   utils.GenMockId(),
			Name: "wait_execution_created",
		}
		createGroupResponse = CreateGroupResponse{
			RequestId: requestId,
			Data:      data,
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
		err = json.Unmarshal(resp, &createGroupResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createGroupResponse, nil
}

func (c *AutoscalingClient) CreateConfiguration(request CreateConfigurationRequest, requestId string, sessionId string, timeout time.Duration) (*CreateConfigurationResponse, error) {
	fullUrl := c.Endpoint + CreateConfigurationUrl

	var createConfigurationResponse CreateConfigurationResponse
	if c.IsMock {
		data := Configuration{
			Id:   utils.GenMockId(),
			Name: "wait_execution_created",
		}
		createConfigurationResponse = CreateConfigurationResponse{
			RequestId: requestId,
			Data:      data,
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
		err = json.Unmarshal(resp, &createConfigurationResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createConfigurationResponse, nil
}

func (c *AutoscalingClient) CreatePolicy(request CreatePolicyRequest, requestId string, sessionId string, timeout time.Duration) (*CreatePolicyResponse, error) {
	fullUrl := c.Endpoint + CreatePolicyUrl

	var createPolicyResponse CreatePolicyResponse
	if c.IsMock {
		data := Policy{
			Id:   utils.GenMockId(),
			Name: "wait_execution_created",
		}
		createPolicyResponse = CreatePolicyResponse{
			RequestId: requestId,
			Data:      data,
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
		err = json.Unmarshal(resp, &createPolicyResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createPolicyResponse, nil
}

func (c *AutoscalingClient) DeleteGroup(request DeleteGroupRequest, requestId string, sessionId string, timeout time.Duration) (*DeleteGroupResponse, error) {
	fullUrl := c.Endpoint + DeleteGroupUrl

	var deleteGroupResponse DeleteGroupResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &deleteGroupResponse)
		if err != nil {
			return nil, err
		}
	}

	return &deleteGroupResponse, nil
}

func (c *AutoscalingClient) DeleteConfiguration(request DeleteConfigurationRequest, requestId string, sessionId string, timeout time.Duration) (*DeleteConfigurationResponse, error) {
	fullUrl := c.Endpoint + DeleteConfigurationUrl

	var deleteConfigurationResponse DeleteConfigurationResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &deleteConfigurationResponse)
		if err != nil {
			return nil, err
		}
	}

	return &deleteConfigurationResponse, nil
}

func (c *AutoscalingClient) DeletePolicy(request DeletePolicyRequest, requestId string, sessionId string, timeout time.Duration) (*DeletePolicyResponse, error) {
	fullUrl := c.Endpoint + DeletePolicyUrl

	var deletePolicyResponse DeletePolicyResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &deletePolicyResponse)
		if err != nil {
			return nil, err
		}
	}

	return &deletePolicyResponse, nil
}

func (c *AutoscalingClient) ScaleHostInstance(request ScaleHostInstanceRequest, requestId string, sessionId string, timeout time.Duration) (*ScaleHostInstanceResponse, error) {
	fullUrl := c.Endpoint + ScaleHostInstanceUrl

	var scaleHostInstanceResponse ScaleHostInstanceResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &scaleHostInstanceResponse)
		if err != nil {
			return nil, err
		}
	}

	return &scaleHostInstanceResponse, nil
}

func (c *AutoscalingClient) ListElasticLog(request ListElasticLogRequest, requestId string, sessionId string, timeout time.Duration) (*ListElasticLogResponse, error) {
	fullUrl := c.Endpoint + ListElasticLogUrl

	var listElasticLogResponse ListElasticLogResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &listElasticLogResponse)
		if err != nil {
			return nil, err
		}
	}

	return &listElasticLogResponse, nil
}

func (c *AutoscalingClient) GetElasticLog(request GetElasticLogRequest, requestId string, sessionId string, timeout time.Duration) (*GetElasticLogResponse, error) {
	fullUrl := c.Endpoint + GetElasticLogUrl

	var getElasticLogResponse GetElasticLogResponse
	if c.IsMock {

	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &getElasticLogResponse)
		if err != nil {
			return nil, err
		}
	}

	return &getElasticLogResponse, nil
}
