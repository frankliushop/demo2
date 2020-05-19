package datacontroller

import (
	"fmt"
	"myproject1/datacommon"
	"myproject1/datamodel"
	"net/http"
	"log"

	jsoniter "github.com/json-iterator/go"
)

//處理瀏覽器crossdomain問題
func ControllerCrossDomain(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
}

//處理錯誤
func ErrorHandler(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err)
		switch err.(type) {
		case datacommon.ExceptionData:
			exp := err.(datacommon.ExceptionData)
			resultResponse := datamodel.ResultResponse{
				Result:     false,
				Code:       exp.ErrorCode,
				Message:    exp.ErrorMessage,
				DataResult: nil,
			}
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			jsondata, _ := json.Marshal(&resultResponse)
			jsonString := string(jsondata)			
			fmt.Fprintf(w, jsonString)
		case error:
			resultResponse := datamodel.ResultResponse{
				Result:     false,
				Code:       datacommon.ErrCodeSystemError,
				Message:    err.Error(),
				DataResult: nil,
			}
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			jsondata, _ := json.Marshal(&resultResponse)
			jsonString := string(jsondata)
			fmt.Fprintf(w, jsonString)
		}
		return true
	}
	return false
}
