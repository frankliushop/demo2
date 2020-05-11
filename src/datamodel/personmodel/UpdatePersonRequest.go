package personmodel

import (
	"myproject1/datacommon"
)

type UpdatePersonRequest struct {
	ID          int32                 `json:"id"`
	Name        datacommon.NullString `json:"name"`
	Phone       datacommon.NullString `json:"phone"`
	MobilePhone datacommon.NullString `json:"mobilePhone"`
	Address     datacommon.NullString `json:"address"`
	Birthday    datacommon.NullTime   `json:"birthday"`
}

func (updatePersonRequest *UpdatePersonRequest) CheckValue() {
	if updatePersonRequest.ID == 0 {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}

	if *updatePersonRequest.Name.Data == "" || updatePersonRequest.Name.Data == nil {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "Name can not be empty",
		})
	}
}
