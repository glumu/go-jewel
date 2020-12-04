package composer

import (
	"encoding/json"
	"time"
)

func (c *Client) ListAvailableZone(ids []string, requestId string, timeout time.Duration) (*ListAvailableZoneResponse, error) {
	fullUrl := c.Endpoint + ListAvailableZoneUrl

	var response ListAvailableZoneResponse
	head := c.GenHeaders(requestId, nil)
	request := ListAvailableZoneRequest{
		Ids: ids,
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
