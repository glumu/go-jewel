package storage

import (
	"encoding/json"
	"time"
)

func (c *StorageClient) AddLun(request AddLunRequest, requestId string, timeout time.Duration) (*AddLunResponse, error) {
	fullUrl := c.Endpoint + AddLunUrl

	var response AddLunResponse
	if c.IsMock {
		response = AddLunResponse{
			Code: 0,
			Data: "",
			Msg:  "",
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

func (c *StorageClient) RemoveLun(request RemoveLunRequest, requestId string, timeout time.Duration) (*RemoveLunResponse, error) {
	fullUrl := c.Endpoint + RemoveLunUrl

	var response RemoveLunResponse
	if c.IsMock {
		response = RemoveLunResponse{
			Code: 0,
			Data: "",
			Msg:  "",
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

func (c *StorageClient) ListLun(request ListLunRequest, requestId string, timeout time.Duration) (*ListLunResponse, error) {
	fullUrl := c.Endpoint + ListLunUrl

	var response ListLunResponse
	if c.IsMock {
		response = ListLunResponse{
			Code: 0,
			Data: map[string]string{
				"77777": "test_lun",
			},
			Msg: "ok",
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
