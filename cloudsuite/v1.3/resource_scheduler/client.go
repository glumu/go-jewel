package resource_scheduler

import "github.com/glumu/go-jewel/pkg/request"

type Client struct {
	request.BaseClient
}

const (
	ListQuotaUrl    = "/listQuota"
	ReserveQuotaUrl = "/reserveQuota"
)
