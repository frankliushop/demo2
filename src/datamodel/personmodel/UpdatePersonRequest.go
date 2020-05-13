package personmodel

import (
	"myproject1/datacommon"

	"gopkg.in/guregu/null.v3"
)

type UpdatePersonRequest struct {
	ID          null.Int                  `json:"id"`
	Name        null.String               `json:"name"`
	Phone       null.String               `json:"phone"`
	MobilePhone null.String               `json:"mobilePhone"`
	Address     null.String               `json:"address"`
	Birthday    datacommon.NullDateExtend `json:"birthday"`
}

func (updatePersonRequest *UpdatePersonRequest) CheckValue() {
	if datacommon.IsNullOrZero(updatePersonRequest.ID) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}

	if datacommon.IsNullOrEmpty(updatePersonRequest.Name) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		})
	}
}
