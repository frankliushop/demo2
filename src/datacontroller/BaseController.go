package datacontroller

import (
	"fmt"
	"myproject1/datacommon"
	"myproject1/datamodel"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

//Controller例外處理
func ControllerErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := recover()
	if err != nil {
		exp := err.(datacommon.ExceptionData)
		resultResponse := datamodel.ResultResponse{
			Result:  false,
			Code:    exp.ErrorCode,
			Message: exp.ErrorMessage,
			Data:    nil,
		}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		jsondata, _ := json.Marshal(&resultResponse)
		jsonString := string(jsondata)
		fmt.Fprintf(w, jsonString)
	}
}
