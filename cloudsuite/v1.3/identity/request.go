package identity

type IdRequest struct {
	Id string `json:"id"`
}

type ListVdcRequest struct {
	PageNumber int             `json:"pageNumber'`
	PageSize   int             `json:"pageSize"`
	Tags       map[string]bool `json:"tags"`
}
