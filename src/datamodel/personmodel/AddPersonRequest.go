package personmodel

import (
	"myproject1/datacommon"
)

type AddPersonRequest struct {
	Name        datacommon.NullString `json:"name"`
	Phone       datacommon.NullString `json:"phone"`
	MobilePhone datacommon.NullString `json:"mobilePhone"`
	Address     datacommon.NullString `json:"address"`
	Birthday    datacommon.NullTime   `json:"birthday"`
}

func (addPersonRequest *AddPersonRequest) CheckValue() {
	if *addPersonRequest.Name.Data == "" || addPersonRequest.Name.Data == nil {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		})
	}
}
