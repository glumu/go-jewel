package resource

import (
	"encoding/json"
	"time"

	"github.com/go-jewel/cloudsuite/resource/request"
	"github.com/go-jewel/cloudsuite/resource/response"
)

func (c *ResourceClient) GetVM(id string, requestId string, timeout time.Duration) (*responseData.GetVMResponse, error) {
	fullUrl := c.Endpoint + GetVMUrl

	var response responseData.GetVMResponse

	head := c.GenHeaders(requestId, nil)
	request := requestData.GetVMRequest{
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

func (c *ResourceClient) ResizeFlavor(request requestData.ResizeFlavorRequest,
	requestId string, timeout time.Duration) (*responseData.ResizeFlavorResponse, error) {
	fullUrl := c.Endpoint + ResizeFlavorUrl

	var response responseData.ResizeFlavorResponse
	if c.IsMock {
		response = responseData.ResizeFlavorResponse{
			responseData.BaseResponse{
				RequestId: requestId,
				TimeStamp: 123,
			},
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
