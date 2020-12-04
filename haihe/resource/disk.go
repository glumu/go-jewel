package resource

import (
	"encoding/json"
	"time"
)

func (c *ResourceClient) UpdateDiskQos(request UpdateDiskQosRequest, requestId string, timeout time.Duration) (*UpdateDiskQosResponse, error) {
	fullUrl := c.Endpoint + UpdateDiskQosUrl

	var response UpdateDiskQosResponse
	if c.IsMock {
		response = UpdateDiskQosResponse{
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
