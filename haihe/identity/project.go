package identity

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *IdentityClient) GetProject(request GetProjectRequest, requestId string, timeout time.Duration) (*GetProjectResponse, error) {
	fullUrl := c.Endpoint + GetProjectUrl

	var response GetProjectResponse
	if c.IsMock {
		response = GetProjectResponse{
			RequestId: requestId,
			Data: Detail{
				Name: "TestProject",
			},
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

func (c *IdentityClient) CreateProject(request CreateProjectRequest, requestId string, sessionId string, timeout time.Duration) (*CreateProjectResponse, error) {
	fullUrl := c.Endpoint + CreateProjectUrl

	var response CreateProjectResponse
	if c.IsMock {
		response = CreateProjectResponse{
			RequestId: requestId,
			Data: Project{
				Name: "TestProject",
			},
			TimeStamp: 123,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *IdentityClient) DeleteProject(request DeleteProjectRequest, requestId string, sessionId string, timeout time.Duration) (*DeleteProjectResponse, error) {
	fullUrl := c.Endpoint + DeleteProjectUrl

	var response DeleteProjectResponse
	if c.IsMock {
		response = DeleteProjectResponse{
			RequestId: requestId,
			TimeStamp: 123,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *IdentityClient) Login(request LoginRequest, requestId string, timeout time.Duration) (*LoginResponse, error) {
	fullUrl := c.Endpoint + LoginUrl

	var response LoginResponse
	if c.IsMock {
		response = LoginResponse{
			RequestId: requestId,
			Data: Login{
				SessionId: "TestProject",
				UserId:    "TestProject",
			},
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

func (c *IdentityClient) Logout(request LogoutRequest, requestId string, timeout time.Duration) (*LogoutResponse, error) {
	fullUrl := c.Endpoint + LogoutUrl

	var response LogoutResponse
	if c.IsMock {
		response = LogoutResponse{
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

func (c *IdentityClient) ProjectExpirationTimeUpdate(request UpdateProjectExpirationTimeRequest, requestId string, sessionId string, timeout time.Duration) (*UpdateProjectExpirationTimeResponse, error) {
	fullUrl := c.Endpoint + UpdateProjectExpirationTimeUrl

	var response UpdateProjectExpirationTimeResponse
	if c.IsMock {
		response = UpdateProjectExpirationTimeResponse{
			RequestId: requestId,
			TimeStamp: 123,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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

func (c *IdentityClient) GetProjectInTest(request GetProjectRequest, requestId string,
	sessionId string, timeout time.Duration) (*GetProjectResponse, error) {
	fullUrl := c.Endpoint + GetProjectUrl

	var response GetProjectResponse
	if c.IsMock {
		response = GetProjectResponse{
			RequestId: requestId,
			Data: Detail{
				Name: "TestProject",
			},
			TimeStamp: 123,
		}
	} else {
		head := c.GenHeaders(requestId, [][2]string{
			{"Cookie", fmt.Sprintf("sessionId=" + sessionId)},
			{"X-Scope", "System"},
		})
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
