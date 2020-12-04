package rds_trove

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *RdsTroveClient) GetUsers(tenantId string, instanceId string,
	requestId string, timeout time.Duration) (*GetUsersResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+UserUrl, tenantId, instanceId)

	var response GetUsersResponse
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

func (c *RdsTroveClient) CreateUser(request CreateUserRequest, tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+UserUrl, tenantId, instanceId)

	_, err := c.TroveRequest(fullUrl, "POST", nil, request, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) DeleteUser(tenantId string, instanceId string, user string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+DeleteUserUrl, tenantId, instanceId, user)
	_, err := c.TroveRequest(fullUrl, "DELETE", nil, nil, GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}
