package rds_trove

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *RdsTroveClient) GetConfGroups(tenantId string, requestId string, timeout time.Duration) (*GetConfGroupsResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfUrl, tenantId)

	var response GetConfGroupsResponse
	resp, err := c.TroveRequest(fullUrl, "GET", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 200)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *RdsTroveClient) GetConfInstances(tenantId string, confId string, requestId string, timeout time.Duration) (*GetConfInstancesResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfInstanceUrl, tenantId, confId)

	var response GetConfInstancesResponse
	resp, err := c.TroveRequest(fullUrl, "GET", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 200)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *RdsTroveClient) GetConfItems(tenantId string, confId string, requestId string, timeout time.Duration) (*GetConfItmesResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfItemsUrl, tenantId, confId)

	var response GetConfItmesResponse
	resp, err := c.TroveRequest(fullUrl, "GET", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 200)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *RdsTroveClient) CreateConfGroup(request CreateConfGroupRequest, tenantId string, requestId string, timeout time.Duration) (*CreateConfGroupResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfUrl, tenantId)

	var response CreateConfGroupResponse
	resp, err := c.TroveRequest(fullUrl, "POST", nil, request, GenTroveHeader(tenantId, requestId), timeout, 200)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *RdsTroveClient) DeleteConfGroup(tenantId string, confId string, requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfItemsUrl, tenantId, confId)

	_, err := c.TroveRequest(fullUrl, "DELETE", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) AddConfItems(request AddConfItemsRequest, tenantId string,
	confId string, requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfItemsUrl, tenantId, confId)

	_, err := c.TroveRequest(fullUrl, "PATCH", nil, request, GenTroveHeader(tenantId, requestId), timeout, 200)
	return err
}

func (c *RdsTroveClient) RestConf(request ResetConfRequest, tenantId string,
	confId string, requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ConfItemsUrl, tenantId, confId)

	_, err := c.TroveRequest(fullUrl, "PUT", nil, request, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}
