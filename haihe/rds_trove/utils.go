package rds_trove

import "strings"

func GenTroveHeader(tenantId string, requestId string) map[string]string {
	head := map[string]string{
		"X-Tenant-Id":            tenantId,
		"X-Openstack-Request-Id": "req-" + strings.ToLower(requestId),
		"X-Request-Id":           requestId,
	}
	return head
}
