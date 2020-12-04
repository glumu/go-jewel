package resource

import (
	"encoding/json"
	"time"

	"github.com/go-jewel/cloudsuite/resource/request"
	"github.com/go-jewel/cloudsuite/resource/response"
)

func (c *ResourceClient) CreateDisk(request CreateDiksRequest, requestId string, timeout time.Duration) (*CreateDiskResponse, error) {
	fullUrl := c.Endpoint + CreateDiskUrl

	var response CreateDiskResponse
	head := map[string]string{
		"X-Request-Id": requestId,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *ResourceClient) DeleteDisk(ids []string, requestId string, timeout time.Duration) error {
	fullUrl := c.Endpoint + DeleteDiskUrl

	request := DeleteDisksRequest{
		Ids: ids,
	}
	head := map[string]string{
		"X-Request-Id": requestId,
	}
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) GetDisk(id string, requestId string, timeout time.Duration) (*GetDiskResponse, error) {
	fullUrl := c.Endpoint + GetDiskUrl

	var response GetDiskResponse
	head := map[string]string{
		"X-Request-Id": requestId,
	}

	request := GetDiskRequest{
		Id: id,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ResourceClient) ResizeDisk(id string, extendSize int, requestId string, timeout time.Duration) (*ResizeDiskResponse, error) {
	fullUrl := c.Endpoint + ResizeDiskUrl

	var response ResizeDiskResponse
	if c.IsMock {
		response = ResizeDiskResponse{
			RequestId: requestId,
			TimeStamp: 123,
		}
	} else {
		request := ResizeDiskRequest{
			Id:         id,
			ExtendSize: extendSize,
		}
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

func (c *ResourceClient) AttachDisk(vmId string, diskId string, requestId string, timeout time.Duration) (*AttachDiskResponse, error) {
	fullUrl := c.Endpoint + AttachDiskUrl

	var response AttachDiskResponse
	if c.IsMock {
		response = AttachDiskResponse{
			RequestId: requestId,
			TimeStamp: 123,
		}
	} else {
		request := AttachDiskRequest{
			VirtualMachineId: vmId,
			DiskId:           diskId,
		}
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

func (c *ResourceClient) DetachDisk(vmId string, diskId string, reqId string, timeout time.Duration) error {
	fullUrl := c.Endpoint + DetachDiskUrl
	head := c.GenHeaders(reqId, nil)
	request := DetachDiskRequest{
		DiskId:           diskId,
		VirtualMachineId: vmId,
	}
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) UpdateDiskQos(request UpdateDiskQosRequest, requestId string, timeout time.Duration) (*UpdateDiskQosResponse, error) {
	fullUrl := c.Endpoint + UpdateDiskQosUrl

	var response UpdateDiskQosResponse
	if c.IsMock {
		response = UpdateDiskQosResponse{
			RequestId: requestId,
			TimeStamp: 123,
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

func (c *ResourceClient) CreateDiskSnapshot(name string, diskId string, desc string, reqId string, timeout time.Duration) (string, error) {
	fullUrl := c.Endpoint + CreateDiskSnapshotUrl
	var response CreateDiskSnapshotResponse
	head := c.GenHeaders(reqId, nil)
	request := CreateDiskSnapshotRequest{
		Name:        name,
		DiskId:      diskId,
		Description: desc,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return "", err
	}

	return response.Data.Id, nil
}

func (c *ResourceClient) DeleteDiskSnapshot(snapId string, reqId string, timeout time.Duration) error {
	fullUrl := c.Endpoint + DeleteDiskSnapshotUrl
	head := c.GenHeaders(reqId, nil)
	request := DeleteDiskSnapshotRequest{
		Id: snapId,
	}
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) ListDiskType(envId string, zoneId string, requestId string, timeout time.Duration) (*responseData.ListDiskTypeResponse, error) {
	fullUrl := c.Endpoint + ListDiskTypeUrl
	head := c.GenHeaders(requestId, nil)
	var response responseData.ListDiskTypeResponse
	request := requestData.ListDiskTypeRequest{
		EnvironmentId: envId,
		ZoneId:        zoneId,
	}
	resp, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
