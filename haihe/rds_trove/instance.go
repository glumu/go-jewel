package rds_trove

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *RdsTroveClient) CreateInstance(request CreateInstanceRequest,
	tenantId string, requestId string, timeout time.Duration) (*CreateInstanceResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+CreateInstanceUrl, tenantId)

	var response CreateInstanceResponse
	if c.IsMock {
		response = CreateInstanceResponse{
			Instance: ResInstance{
				Id:       "cab267db-003d-4a38-ad84-5185a715eb2d",
				Status:   "BUILD",
				Created:  "2019-10-22T07:28:58",
				Updated:  "2019-10-22T07:28:58",
				Name:     "DBaas_test",
				TenantId: "42c2103fc6174ddd93d35412d5bb1d67",
				Datastore: ResDatastore{
					Type:    "mysql",
					Version: "Mysql-5.6",
				},
			},
		}
	} else {
		resp, err := c.TroveRequest(fullUrl, "POST", nil, request,
			GenTroveHeader(tenantId, requestId), timeout, 200)
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

func (c *RdsTroveClient) PutInstanceConf(request PutInstanceConfRequest, tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+PutInstanceConfUrl, tenantId, instanceId)

	_, err := c.TroveRequest(fullUrl, "PUT", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) GetInstanceConf(tenantId string, instanceId string,
	requestId string, timeout time.Duration) (*GetInstanceConfResponse, error) {
	fullUrl := fmt.Sprintf(c.Endpoint+GetInstanceConfUrl, tenantId, instanceId)

	var response GetInstanceConfResponse
	resp, err := c.TroveRequest(fullUrl, "GET", nil, nil,
		GenTroveHeader(tenantId, requestId), timeout, 200)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *RdsTroveClient) Buid(request BuildRequest, tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) StopService(tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)

	var request StopRequest
	request.Stop = make(map[string]interface{})
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) StartService(tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)

	var request StartRequest
	request.Start = make(map[string]interface{})
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) RestartService(tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)

	var request RestartRequest
	request.Restart = make(map[string]interface{})
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) ResizeLvm(request ResizeLvmRequest, tenantId string, instanceId string,
	requestId string, timeout time.Duration) error {
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)

	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) Mount(tenantId string, instanceId string, disk string, path string,
	requestId string, timeout time.Duration) error {
	request := MountRequest{
		DiskBody{
			Disk: disk,
			Path: path,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) UMount(tenantId string, instanceId string, disk string, path string,
	requestId string, timeout time.Duration) error {
	request := UmountRequest{
		DiskBody{
			Disk: disk,
			Path: path,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) GetDiskUsage(tenantId string, instanceId string, disk string,
	requestId string, timeout time.Duration) error {
	request := GetDiskUsageRequest{
		DiskBody{
			Disk: disk,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) Resize2FS(tenantId string, instanceId string, disk string,
	requestId string, timeout time.Duration) error {
	request := Resize2FSRequest{
		DiskBody{
			Disk: disk,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) Backup(tenantId string, instanceId string, info string, logPath string,
	backupPath string, requestId string, timeout time.Duration) error {
	request := BackupRequest{
		Backup: BackupBody{
			Info: info,
			LogPath: logPath,
			BackupPath: backupPath,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) MergeInc2Full(tenantId string, instanceId string,
	backups []string, logPath string, backupPath string, requestId string, timeout time.Duration) error {
	request := MergeIncRequest{
		MergeInc2Full: MergeIncBody{
			Backups: backups,
			LogPath: logPath,
			BackupPath: backupPath,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}

func (c *RdsTroveClient) DeleteInc(tenantId string, instanceId string,
	backups []string, logPath string, requestId string, timeout time.Duration) error {
	request := DeleteIncRequest{
		DeleteInc: DeleteIncBody{
			Backups: backups,
			LogPath: logPath,
		},
	}
	fullUrl := fmt.Sprintf(c.Endpoint+ActionUrl, tenantId, instanceId)
	_, err := c.TroveRequest(fullUrl, "POST", nil, request,
		GenTroveHeader(tenantId, requestId), timeout, 202)
	return err
}