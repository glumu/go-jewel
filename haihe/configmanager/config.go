package configmanager

import (
	"encoding/json"
	"time"
)

func (c *ConfigClient) GetConfig(request GetConfigRequest, requestId string, timeout time.Duration) (*GetConfigResponse, error) {
	fullUrl := c.Endpoint + GetConfigUrl

	var response GetConfigResponse
	if c.IsMock {
		response = GetConfigResponse{
			RequestId: requestId,
			Data: Detail{
				Name: "TestConfig",
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
