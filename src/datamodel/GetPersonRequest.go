package datamodel

import "myproject1/datacommon"

type GetPersonRequest struct {
	ID int32 `json:"id"`
}

func (getPersonRequest *GetPersonRequest) CheckValue() {
	if getPersonRequest.ID == 0 {
		panic(datacommon.ExceptionData{
			ErrorCode : datacommon.ErrCodeParamterError,
			ErrorMessage : "ID can not be empty",
		})
	}
}