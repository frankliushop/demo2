package personmodel

import (
	"myproject1/datacommon"
	"gopkg.in/guregu/null.v4"
)

type DeletePersonRequest struct {
	ID null.Int `json:"id"`
}

func (deletePersonRequest *DeletePersonRequest) CheckValue() {
	if datacommon.IsNullOrZero(deletePersonRequest.ID) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}
}
