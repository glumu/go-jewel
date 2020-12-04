package taskcube

import (
	"encoding/json"
	"time"
)

func (c *TaskcubeClient) CreateProcInstService(request []interface{},
	requestId string, timeout time.Duration) (*[]string, error) {
	fullUrl := c.Endpoint + CreateProcInstServiceUrl

	var response []string
	if c.IsMock {
		response = []string{}
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
