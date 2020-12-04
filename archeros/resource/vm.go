package resource

import (
	"encoding/json"
	"time"
)

func (c *ResourceClient) GetVM(id string, requestId string, timeout time.Duration) (*GetVMResponse, error) {
	fullUrl := c.Endpoint + GetVMUrl

	var response GetVMResponse

	head := c.GenHeaders(requestId, nil)
	request := GetVMRequest{
		Id: id,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ResourceClient) ResizeFlavor(request ResizeFlavorRequest, requestId string, timeout time.Duration) (*ResizeFlavorResponse, error) {
	fullUrl := c.Endpoint + ResizeFlavorUrl

	var response ResizeFlavorResponse
	if c.IsMock {
		response = ResizeFlavorResponse{
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
