package orchestration

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

const VisitorIdHeader string = "visitorId"

func (c *OrchestrationClient) GetDeployment(visitorId string, id string, requestId string, timeout time.Duration) (*GetDeploymentResponse, error) {
	fullUrl := c.Endpoint + GetDeploymentUrl

	var getDeploymentResponse GetDeploymentResponse
	head := c.GenHeaders(requestId, [][2]string{
		{VisitorIdHeader, visitorId},
	})

	request := GetDeploymentRequest{
		DeploymentId: id,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &getDeploymentResponse)
	if err != nil {
		return nil, err
	}

	return &getDeploymentResponse, nil
}

func (c *OrchestrationClient) CreateKvmDeployment(visitorId string, request CreateKvmDeploymentRequest, requestId string, timeout time.Duration) (*CreateDeploymentResponse, error) {
	fullUrl := c.Endpoint + CreateDeploymentUrl

	var createDeploymentResponse CreateDeploymentResponse
	if c.IsMock {
		data := DeploymentInfo{
			Id:                 utils.GenMockId(),
			Status:             "wait_execution_created",
			BlueprintVersionId: "test",
		}
		createDeploymentResponse = CreateDeploymentResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &createDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createDeploymentResponse, nil
}

func (c *OrchestrationClient) CreateLbDeployment(visitorId string, request CreateLbDeploymentRequest, requestId string, timeout time.Duration) (*CreateDeploymentResponse, error) {
	fullUrl := c.Endpoint + CreateDeploymentUrl

	var createDeploymentResponse CreateDeploymentResponse
	if c.IsMock {
		data := DeploymentInfo{
			Id:                 utils.GenMockId(),
			Status:             "wait_execution_created",
			BlueprintVersionId: "test",
		}
		createDeploymentResponse = CreateDeploymentResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &createDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createDeploymentResponse, nil
}

func (c *OrchestrationClient) CreateRdsDeployment(visitorId string, request CreateRdsDeploymentRequest,
	requestId string, timeout time.Duration) (*CreateDeploymentResponse, error) {
	fullUrl := c.Endpoint + CreateDeploymentUrl
	var createDeploymentResponse CreateDeploymentResponse
	if c.IsMock {
		data := DeploymentInfo{
			Id:                 utils.GenMockId(),
			Status:             "wait_execution_created",
			BlueprintVersionId: "test",
		}
		createDeploymentResponse = CreateDeploymentResponse{
			RequestId: requestId,
			Data:      data,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})

		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &createDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &createDeploymentResponse, nil
}

func (c *OrchestrationClient) DeleteDeployment(visitorId string, request DeleteDeploymentRequest, requestId string, timeout time.Duration) (*DeleteDeploymentResponse, error) {
	fullUrl := c.Endpoint + DeleteDeploymentUrl

	var deleteDeploymentResponse DeleteDeploymentResponse
	if c.IsMock {
		deleteDeploymentResponse = DeleteDeploymentResponse{
			RequestId: requestId,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resp, &deleteDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &deleteDeploymentResponse, nil
}

func (c *OrchestrationClient) ListNodeInstance(request ListNodeInstanceRequest, requestId string, timeout time.Duration) (*ListNodeInstanceResponse, error) {
	fullUrl := c.Endpoint + ListNodeInstanceUrl

	var listNodeInstanceResponse ListNodeInstanceResponse
	if c.IsMock {
		nodeInstance := NodeInstance{
			Id:             utils.GenMockId(),
			EnvironmentId:  "1",
			ZoneId:         utils.GenMockId(),
			ProjectName:    "testProject",
			OwnerId:        utils.GenMockId(),
			Name:           "test",
			IsNew:          true,
			Type:           "VM",
			State:          "up",
			Ref:            "test",
			NodeInstanceId: utils.GenMockId(),
			ResourceId:     utils.GenMockId(),
		}
		listNodeInstanceResponse = ListNodeInstanceResponse{
			Data: []NodeInstance{nodeInstance},
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resp, &listNodeInstanceResponse)
		if err != nil {
			return nil, err
		}
	}

	return &listNodeInstanceResponse, nil
}

func (c *OrchestrationClient) ListDeployment(request ListDeploymentRequest, requestId string, timeout time.Duration) (*ListDeploymentResponse, error) {
	fullUrl := c.Endpoint + ListDeploymentUrl

	var listDeploymentResponse ListDeploymentResponse
	if c.IsMock {
		d := Deployment{
			Id:                   utils.GenMockId(),
			Name:                 "test",
			OwnerId:              utils.GenMockId(),
			BlueprintId:          utils.GenMockId(),
			BlueprintName:        "test",
			BlueprintVersionId:   utils.GenMockId(),
			BlueprintVersionName: "test",
			Status:               "ok",
		}
		listDeploymentResponse = ListDeploymentResponse{
			Data: []Deployment{d},
		}
	} else {
		head := c.GenHeaders(requestId, nil)
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &listDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &listDeploymentResponse, nil
}

func (c *OrchestrationClient) UpdateDeployment(visitorId string, request UpdateDeploymentRequest, requestId string, timeout time.Duration) (*UpdateDeploymentResponse, error) {
	fullUrl := c.Endpoint + UpdateDeploymentUrl

	var updateDeploymentResponse UpdateDeploymentResponse
	if c.IsMock {
		updateDeploymentResponse = UpdateDeploymentResponse{
			Id:           utils.GenMockId(),
			State:        "on",
			DeploymentId: utils.GenMockId(),
			ExecutionId:  utils.GenMockId(),
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resp, &updateDeploymentResponse)
		if err != nil {
			return nil, err
		}
	}

	return &updateDeploymentResponse, nil
}

func (c *OrchestrationClient) StartExecution(visitorId string, request interface{}, requestId string, timeout time.Duration) (*StartExecutionResponse, error) {
	fullUrl := c.Endpoint + StartExecutionUrl

	var startExecutionResponse StartExecutionResponse
	if c.IsMock {
		d := ExecutionData{
			Id:            utils.GenMockId(),
			DeploymentId:  utils.GenMockId(),
			WorkflowId:    utils.GenMockId(),
			Error:         "",
			BlueprintId:   utils.GenMockId(),
			StatusDisplay: "test",
			StartedAt:     "2019-04-11 11:22:14",
			EndedAt:       "2019-04-11 11:22:14",
		}
		startExecutionResponse = StartExecutionResponse{
			Data: d,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{VisitorIdHeader, visitorId},
		})
		resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(resp, &startExecutionResponse)
		if err != nil {
			return nil, err
		}
	}

	return &startExecutionResponse, nil
}

func (c *OrchestrationClient) ListEvent(request ListEventRequest, requestId string, sessionId string, timeout time.Duration) (*ListEventResponse, error) {
	fullUrl := c.Endpoint + ListEventUrl

	var listEventResponse ListEventResponse
	if c.IsMock {
		d := Event{
			BlueprintId:             utils.GenMockId(),
			DeploymentId:            utils.GenMockId(),
			ExecutionId:             utils.GenMockId(),
			NodeInstanceId:          utils.GenMockId(),
			ResourceId:              utils.GenMockId(),
			ResourceName:            "test_a0ce37ef",
			ResourceOperation:       "wait_created",
			ResourceOperationResult: "succeed",
			ResourceType:            "vm",
			WorkflowId:              "heal",
		}
		listEventResponse = ListEventResponse{
			Data: []Event{d},
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
		err = json.Unmarshal(resp, &listEventResponse)
		if err != nil {
			return nil, err
		}
	}

	return &listEventResponse, nil
}
