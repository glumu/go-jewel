package monitor

type GroupDataResponse struct {
	Id   int    `json:"id"`
	Name string `json:"grp_name"`
}

type CreateGroupResponse struct {
	Code int               `json:"code"`
	Data GroupDataResponse `json:"data"`
}

type DeleteGroupResponse struct {
	Code int               `json:"code"`
	Data GroupDataResponse `json:"data"`
}

type UpdateGroupResponse struct {
	Code int               `json:"code"`
	Data GroupDataResponse `json:"data"`
}

type BindGroupResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type UnbindGroupResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type StrategyData struct {
	Id int `json:"id"`
	//Name string `json:"name"`
	//MonitorType        string `json:"monitorType"`
	//ParentId           string `json:"parentId"`
	//ActionId           int    `json:"actionId"`
	//UserName           string `json:"userName"`
	//UserRole           string `json:"userRole"`
	//IsShare            int    `json:"isShare"`
	//MetricCatalogIdStr string `json:"metricCatalogIdStr"`
	//Status             string `json:"status"`
}
type CreateStrategyResponse struct {
	Code int          `json:"code"`
	Data StrategyData `json:"data"`
}
type UpdateStrategyResponse struct {
	Code int          `json:"code"`
	Data StrategyData `json:"data"`
}

type DeleteStrategyResponse struct {
	Code int `json:"code"`
	Id   int `json:"id"`
}

type SyncHostToGroupResponse struct {
	Code int `json:"code"`
}
