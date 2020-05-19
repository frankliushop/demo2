package personmodel

import (
	"myproject1/datacommon"
	"gopkg.in/guregu/null.v3"
)

type GetPersonRequest struct {
	ID null.Int `json:"id"`
}

//檢查必須輸入值是否存在
func (this *GetPersonRequest) CheckValue() {
	if datacommon.IsNullOrZero(this.ID) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}
}
