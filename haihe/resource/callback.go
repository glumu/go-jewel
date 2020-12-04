package resource

import (
	"encoding/json"
	"time"
)

func (c *ResourceClient) Callback(request CallbackRequest, requestId string, timeout time.Duration) (*CallbackResponse, error) {
	fullUrl := c.Endpoint + CallbackUrl

	var response CallbackResponse
	if c.IsMock {
		response = CallbackResponse{
			RequestId: requestId,
			Result:    "",
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
