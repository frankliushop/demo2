package personmodel

import (
	"myproject1/datacommon"
	"gopkg.in/guregu/null.v3"
)

type GetPersonRequest struct {
	ID null.Int `json:"id"`
}

func (getPersonRequest *GetPersonRequest) CheckValue() {
	if datacommon.IsNullOrZero(getPersonRequest.ID) {
		panic(datacommon.ExceptionData{
			ErrorCode:    datacommon.ErrCodeParamterError,
			ErrorMessage: "ID can not be empty",
		})
	}
}
