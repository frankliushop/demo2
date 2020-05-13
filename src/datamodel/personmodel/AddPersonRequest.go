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

func (addPersonRequest *AddPersonRequest) CheckValue() {
	if datacommon.IsNullOrEmpty(addPersonRequest.Name) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		})
	}
}
