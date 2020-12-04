package identity

import (
	"encoding/json"
	"time"
)

func (c *IdentityClient) GetUser(request GetUserRequest, requestId string, timeout time.Duration) (*GetUserResponse, error) {
	fullUrl := c.Endpoint + GetUserUrl

	var response GetUserResponse
	if c.IsMock {
		response = GetUserResponse{
			RequestId: requestId,
			Data: UserDetail{
				Name: "TestUser",
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
