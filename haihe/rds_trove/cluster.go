package rds_trove

import (
	"fmt"
	"time"
)

func (c *RdsTroveClient) CreateCluster(request CreateClusterRequest,
	tenantId string, requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+CreateClusterUrl, tenantId)

	_, err := c.TroveRequest(fullUrl, "POST", nil, request, GenTroveHeader(tenantId, requestId), timeout, 200)
	return err
}
