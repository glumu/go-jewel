package rds_trove

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *RdsTroveClient) GetDatabases(tenantId string, instanceId string,
	requestId string, timeout time.Duration) (*GetDatabasesResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+DatabaseUrl, tenantId, instanceId)

	var response GetDatabasesResponse
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

func (c *RdsTroveClient) CreateDatabase(request CreateDatabaseRequest, tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+DatabaseUrl, tenantId, instanceId)

	_, err := c.TroveRequest(fullUrl, "POST", nil, request, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) DeleteDatabase(tenantId string, instanceId string, database string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+DeleteDatabaseUrl, tenantId, instanceId, database)
	_, err := c.TroveRequest(fullUrl, "DELETE", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}
