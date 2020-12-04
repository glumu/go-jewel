package resource

import (
	"encoding/json"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

func (c *ResourceClient) ListVM(request ListVMRequest, requestId string, timeout time.Duration) (*ListVMResponse, error) {
	fullUrl := c.Endpoint + ListVMUrl

	var response ListVMResponse
	if c.IsMock {
		response = ListVMResponse{
			RequestId: requestId,
			Data:      []VM{},
			TimeStamp: 123,
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

func (c *ResourceClient) GetVM(request GetVMRequest, requestId string, timeout time.Duration) (*GetVMResponse, error) {
	fullUrl := c.Endpoint + GetVMUrl

	var response GetVMResponse
	if c.IsMock {
		response = GetVMResponse{
			RequestId: requestId,
			Data: Detail{
				Name:      "TestVM",
				Id:        utils.GenMockId(),
				Ref:       utils.GenMockId(),
				OwnerId:   utils.GenMockId(),
				OwnerName: "TestProject",
				Status:    "START",
			},
			TimeStamp: 123,
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

func (c *ResourceClient) ResizeVM(request ResizeRequest, requestId string, timeout time.Duration) (*ResizeResponse, error) {
	fullUrl := c.Endpoint + ResizeVMUrl

	var response ResizeResponse
	if c.IsMock {
		response = ResizeResponse{
			RequestId: requestId,
			TimeStamp: 123,
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
