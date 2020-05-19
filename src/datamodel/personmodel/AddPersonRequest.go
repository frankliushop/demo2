package personmodel

import (
	"myproject1/datacommon"
	"gopkg.in/guregu/null.v3"
)

type AddPersonRequest struct {
	Name        null.String `json:"name"`
	Phone       null.String `json:"phone"`
	MobilePhone null.String `json:"mobilePhone"`
	Address     null.String `json:"address"`
	Birthday    datacommon.NullDateExtend   `json:"birthday"`
}
//檢查必須輸入值是否存在
func (this *AddPersonRequest) CheckValue() error {
	if datacommon.IsNullOrEmpty(this.Name) {
		return datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		}
	}
	return nil
}
