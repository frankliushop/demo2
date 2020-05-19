package personmodel

import (
	"myproject1/datacommon"
	"gopkg.in/guregu/null.v3"
)

type GetAllRequest struct {
	PageIndex null.Int `json:"pageIndex"`
	PageSize null.Int `json:"pageSize"`
}

//檢查必須輸入值是否存在
func (this *GetAllRequest) CheckValue() {
	if datacommon.IsNullOrZero(this.PageIndex) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "pageIndex can not be empty or zero",
		})
	}
	if datacommon.IsNullOrZero(this.PageSize) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "pageSize can not be empty or zero",
		})
	}
}