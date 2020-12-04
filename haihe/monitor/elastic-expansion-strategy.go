package monitor

import (
	"encoding/json"
	"math/rand"
	"time"
)

func (c *MonitorClient) CreateStrategy(request CreateStrategyRequest, requestId string, timeout time.Duration) (*CreateStrategyResponse, error) {
	fullUrl := c.Endpoint + CreateStrategyUrl

	var response CreateStrategyResponse
	if c.IsMock {
		response = CreateStrategyResponse{
			Code: 0,
			Data: StrategyData{
				Id: rand.Intn(10000),
			},
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

func (c *MonitorClient) UpdateStrategy(request UpdateStrategyRequest, requestId string, timeout time.Duration) (*UpdateStrategyResponse, error) {
	fullUrl := c.Endpoint + UpdateStrategyUrl

	var response UpdateStrategyResponse
	if c.IsMock {
		response = UpdateStrategyResponse{
			Code: 0,
			Data: StrategyData{
				Id: request.Id,
			},
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

func (c *MonitorClient) DeleteStrategy(request DeleteStrategyRequest, requestId string, timeout time.Duration) (*DeleteStrategyResponse, error) {
	fullUrl := c.Endpoint + DeleteStrategyUrl

	var response DeleteStrategyResponse
	if c.IsMock {
		response = DeleteStrategyResponse{
			Code: 0,
			Id:   request.Id,
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
