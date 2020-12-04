package resource_scheduler

import "github.com/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	ListQuotaUrl    = "/listQuota"
	ReserveQuotaUrl = "/reserveQuota"
)
