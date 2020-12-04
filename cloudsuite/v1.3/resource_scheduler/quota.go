package resource_scheduler

import (
	"encoding/json"
	"time"
)

func (c *Client) ListQuota(ownerId string, requestId string, timeout time.Duration) (*ListQuotaResponse, error) {
	fullUrl := c.Endpoint + ListQuotaUrl

	var response ListQuotaResponse
	head := c.GenHeaders(requestId, nil)
	request := ListQuotaRequest{
		OwnerIds: []string{ownerId},
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

func (c *Client) ReserveQuota(request *ReserveQuotaRequest, requestId string, timeout time.Duration) (*ReserveQuotaResponse, error) {
	fullUrl := c.Endpoint + ReserveQuotaUrl

	var response ReserveQuotaResponse
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
