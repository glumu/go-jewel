package identity

import (
	"encoding/json"
	"errors"
	"time"
)

func (c *Client) GetProject(id string, requestId string, timeout time.Duration) (*GetProjectResponse, error) {
	fullUrl := c.Endpoint + GetProjectUrl

	var response GetProjectResponse
	head := c.GenHeaders(requestId, nil)
	resp, err := c.Request(fullUrl, "POST", nil, IdRequest{Id: id}, head, timeout)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	if response.Data.IsDeleted == 1 {
		return nil, errors.New("project is deleted")
	}

	return &response, nil
}

func (c *Client) ListMyProject(visitorId string, requestId string, timeout time.Duration) (*ListMyProjectResponse, error) {
	fullUrl := c.Endpoint + ListMyProjectUrl

	var response ListMyProjectResponse
	head := c.GenHeaders(requestId, [][2]string{
		{"visitorId", visitorId},
	})
	resp, err := c.Request(fullUrl, "POST", nil, struct {
	}{}, head, timeout)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
