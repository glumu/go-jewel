package identity

import (
	"encoding/json"
	"errors"
	"time"
)

func (c *Client) GetVdc(id string, requestId string, timeout time.Duration) (*GetVdcResponse, error) {
	fullUrl := c.Endpoint + GetVdcUrl

	var response GetVdcResponse
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
		return nil, errors.New("vdc is deleted")
	}

	return &response, nil
}

func (c *Client) ListVdc(request ListVdcRequest, requestId string, timeout time.Duration) (*ListVdcResponse, error) {
	fullUrl := c.Endpoint + ListVdcUrl

	var response ListVdcResponse
	head := c.GenHeaders(requestId, nil)
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
