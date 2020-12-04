package resource

import (
	"encoding/json"
	"time"

	"github.com/glumu/go-jewel/cloudsuite/resource/request"
	"github.com/glumu/go-jewel/cloudsuite/resource/response"
)

func (c *ResourceClient) AttachDisk(vmId string, diskId string, requestId string, timeout time.Duration) (*responseData.AttachDiskResponse, error) {
	fullUrl := c.Endpoint + AttachDiskUrl

	var response responseData.AttachDiskResponse

	if c.IsMock {
		response = responseData.AttachDiskResponse{
			responseData.BaseResponse{
				RequestId: requestId,
				TimeStamp: 123,
			},
		}
	} else {
		request := requestData.AttachDiskRequest{
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
	request := requestData.DetachDiskRequest{
		DiskId:           diskId,
		VirtualMachineId: vmId,
	}
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) CreateDisk(request requestData.CreateDiskRequest, requestId string, timeout time.Duration) (*responseData.CreateDiskResponse, error) {
	fullUrl := c.Endpoint + CreateDiskUrl

	var response responseData.CreateDiskResponse
	head := c.GenHeaders(requestId, nil)
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

	request := requestData.DeleteDisksRequest{
		Ids: ids,
	}
	head := c.GenHeaders(requestId, nil)
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) GetDisk(id string, requestId string, timeout time.Duration) (*responseData.GetDiskResponse, error) {
	fullUrl := c.Endpoint + GetDiskUrl

	var response responseData.GetDiskResponse
	head := c.GenHeaders(requestId, nil)

	request := requestData.GetDiskRequest{
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

func (c *ResourceClient) ResizeDisk(id string, extendSize int, requestId string, timeout time.Duration) (*responseData.ResizeDiskResponse, error) {
	fullUrl := c.Endpoint + ResizeDiskUrl

	var response responseData.ResizeDiskResponse
	if c.IsMock {
		response = responseData.ResizeDiskResponse{
			responseData.BaseResponse{
				RequestId: requestId,
				TimeStamp: 123,
			},
		}
	} else {
		request := requestData.ResizeDiskRequest{
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

func (c *ResourceClient) UpdateDiskQos(request requestData.UpdateDiskQosRequest, requestId string, timeout time.Duration) (*responseData.UpdateDiskQosResponse, error) {
	fullUrl := c.Endpoint + UpdateDiskQosUrl

	var response responseData.UpdateDiskQosResponse
	if c.IsMock {
		response = responseData.UpdateDiskQosResponse{
			responseData.BaseResponse{
				RequestId: requestId,
				TimeStamp: 123,
			},
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

func (c *ResourceClient) CreateDiskSnapshot(request requestData.CreateDiskSnapshotRequest, reqId string, timeout time.Duration) (string, error) {
	fullUrl := c.Endpoint + CreateDiskSnapshotUrl
	var response responseData.CreateDiskSnapshotResponse
	head := c.GenHeaders(reqId, nil)
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
	request := requestData.DeleteDiskSnapshotRequest{
		Id: snapId,
	}
	_, err := c.Request(fullUrl, "POST", nil, request, head, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *ResourceClient) ListDiskType(envId string, zoneId string,
	requestId string, timeout time.Duration) (*responseData.ListDiskTypeResponse, error) {
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
