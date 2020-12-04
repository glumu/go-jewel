package identity

import (
	"encoding/json"
	"time"

	"github.com/go-jewel/haihe/uitls"
)

func (c *IdentityClient) ReserveQuota(request ReserveQuotaRequest, requestId string, timeout time.Duration) (*ReserveQuotaResponse, error) {
	fullUrl := c.Endpoint + ReserveQuotaUrl

	var response ReserveQuotaResponse
	if c.IsMock || utils.IsMockApi(c.MockApi, ReserveQuotaUrl) {
		response = ReserveQuotaResponse{
			RequestId: requestId,
			Data: ReserveQuotaData{
				ReserveId: utils.GenMockId(),
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

func (c *IdentityClient) ListQuota(request ListQutaRequest, requestId string, timeout time.Duration) (*ListQuotaResponse, error) {
	fullUrl := c.Endpoint + ListQuotaUrl

	var response ListQuotaResponse
	if c.IsMock {
		data := ListQuotaData{
			Id:       utils.GenMockId(),
			Limit:    11,
			OwnerId:  request.OwnerIds[0],
			Reserved: 0,
			Used:     1,
			Type:     "BARE_MACHINE",
		}
		response = ListQuotaResponse{
			RequestId: requestId,
			Data: []ListQuotaData{
				data,
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
