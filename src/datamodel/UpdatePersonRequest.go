package datamodel

import (
	"myproject1/datacommon"
)

type UpdatePersonRequest struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Phone       string              `json:"phone"`
	MobilePhone string              `json:"mobilePhone"`
	Address     string              `json:"address"`
	Birthday    datacommon.NullTime `json:"birthday"`
}

func (updatePersonRequest *UpdatePersonRequest) CheckValue() {
	if updatePersonRequest.ID == 0 {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}

	if updatePersonRequest.Name == "" {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		})
	}
}
