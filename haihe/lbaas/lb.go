package lbaas

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *LbaasClient) CreateLb(request CreateLbRequest, requestId string, sessionId string, timeout time.Duration) (*CreateLbResponse, error) {
	fullUrl := c.Endpoint + CreateLbUrl

	var response CreateLbResponse
	if c.IsMock {
		response = CreateLbResponse{
			RequestId: requestId,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *LbaasClient) CreateListener(request CreateListenerRequest, requestId string, sessionId string, timeout time.Duration) (*CreateListenerResponse, error) {
	fullUrl := c.Endpoint + CreateListenerUrl

	var response CreateListenerResponse
	if c.IsMock {
		response = CreateListenerResponse{
			RequestId: requestId,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *LbaasClient) CreateLbPoll(request CreateLbPollRequest, requestId string, sessionId string, timeout time.Duration) (*CreateLbPollResponse, error) {
	fullUrl := c.Endpoint + CreateLbPollUrl

	var response CreateLbPollResponse
	if c.IsMock {
		response = CreateLbPollResponse{
			RequestId: requestId,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *LbaasClient) UpdateLbRelatedResource(request UpdateLbRelatedResourceRequest,
	requestId string, timeout time.Duration) (*UpdateLbRelatedResourceResponse, error) {
	fullUrl := c.Endpoint + UpdateLbRelatedResourceUrl

	var response UpdateLbRelatedResourceResponse
	if c.IsMock {
		response = UpdateLbRelatedResourceResponse{
			RequestId: requestId,
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

func (c *LbaasClient) ListLbBackend(request ListLbBackendRequest, requestId string, timeout time.Duration) (*ListLbBackendResponse, error) {
	fullUrl := c.Endpoint + ListLbBackendUrl

	var response ListLbBackendResponse
	if c.IsMock {
		response = ListLbBackendResponse{
			RequestId: requestId,
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

func (c *LbaasClient) GetLbBackend(request GetLbBackendRequest, requestId string, timeout time.Duration) (*GetLbBackendResponse, error) {
	fullUrl := c.Endpoint + GetLbBackendUrl

	var response GetLbBackendResponse
	if c.IsMock {
		response = GetLbBackendResponse{
			RequestId: requestId,
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
