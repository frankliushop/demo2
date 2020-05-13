package datamodel

type PagingResponse struct {
	PageIndex int64       `json:"pageIndex"`
	PageSize  int64       `json:"pageSize"`
	PageCount int64       `json:"pageCount"`
	DataList  interface{} `json:"dataList"`
}
