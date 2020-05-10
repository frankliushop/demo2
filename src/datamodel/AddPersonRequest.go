package datamodel

import ("myproject1/datacommon"
)

type AddPersonRequest struct {
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	MobilePhone string    `json:"mobilePhone"`
	Address     string    `json:"address"`
	Birthday    datacommon.NullTime `json:"birthday"`
}

func (addPersonRequest *AddPersonRequest) CheckValue() {
	if addPersonRequest.Name == "" {
		panic(datacommon.ExceptionData{
			ErrorCode : datacommon.ErrCodeParamterError,
			ErrorMessage : "Name can not be empty",
		})
	}
}
