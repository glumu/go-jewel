package monitor

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type DeleteGroupRequest struct {
	Id int `json:"id"`
}

type UpdateGroupRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BindGroupRequest struct {
	Hosts []string `json:"hosts"`
	GrpId int      `json:"grpId"`
}

type UnbindGroupRequest struct {
	Hosts []string `json:"hosts"`
	GrpId int      `json:"grpId"`
}

type StrategyBase struct {
	Name        string `json:"name"`
	Metric      string `json:"metric"`
	Func        string `json:"func"`
	Op          string `json:"op"`
	RightValue  string `json:"rightValue"`
	DelaySecond int    `json:"delaySecond"`
	CallBack    string `json:"callback"`
}

type CreateStrategyRequest struct {
	GroupId int `json:"grpId"`
	StrategyBase
}

type UpdateStrategyRequest struct {
	Id int `json:"id"`
	StrategyBase
}

type DeleteStrategyRequest struct {
	Id int `json:"id"`
}

type SyncHostToGroupRequest struct {
	GroupId int      `json:"grpId"`
	Hosts   []string `json:"hosts"`
}
