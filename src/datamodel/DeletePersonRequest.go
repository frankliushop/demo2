package datamodel

import "myproject1/datacommon"

type DeletePersonRequest struct {
	ID          int32     `json:"id"`	
}

func (deletePersonRequest *DeletePersonRequest) CheckValue() {
	if deletePersonRequest.ID == 0 {
		panic(datacommon.ExceptionData{
			ErrorCode : datacommon.ErrCodeParamterError,
			ErrorMessage : "ID can not be empty",
		})
	}
}