package resource

import (
	"encoding/json"
	"time"

	"github.com/glumu/go-jewel/haihe/uitls"
)

func (c *ResourceClient) ListZoneInfo(request ListZoneInfoRequest, requestId string, timeout time.Duration) (*ListZoneInfoResponse, error) {
	fullUrl := c.Endpoint + ListZoneInfoUrl

	var response ListZoneInfoResponse
	if c.IsMock {
		zone := Zone{
			Id:               utils.GenMockId(),
			Uuid:             "",
			Name:             "RegionOne",
			Description:      "RegionOne-178.100",
			EnvironmentId:    "1",
			CreateTime:       "2019-04-11 11:22:14",
			UpdateTime:       "2019-04-11 11:22:14",
			Status:           "UP",
			IdentityEndpoint: "http://172.16.4.131:5000/v3",
			UserName:         "admin",
			Password:         "passw0rd",
			TenantID:         utils.GenMockId(),
			TenantName:       "admin",
			DomainName:       "default",
			IsBaremetal:      true,
		}
		env := Environment{
			Id:               "1",
			Name:             "KVM",
			SyncInterval:     2.0,
			LastSyncTime:     "2019-04-24 15:50:08",
			CloudGatewayName: "openstack-gateway:11404",
			CreateTime:       "2019-04-11 11:22:14",
			UpdateTime:       "2019-04-11 11:22:14",
			HyperVisorType:   "KVM",
			Type:             "PRIVATE_CLOUD",
			IsBaremetal:      true,
			Zones:            []Zone{zone},
		}
		response = ListZoneInfoResponse{
			Environments: []Environment{env},
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
