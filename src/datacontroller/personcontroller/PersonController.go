package personcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	datacontroller "myproject1/datacontroller"
	datamodel "myproject1/datamodel"
	personmodel "myproject1/datamodel/personmodel"
	personservice "myproject1/dataservice/personservice"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

//取得所有資料
func GetAll(w http.ResponseWriter, r *http.Request) {
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req personmodel.GetAllRequest
	err := decoder.Decode(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	err = req.CheckValue()
	if datacontroller.ErrorHandler(err, w) {
		return
	}

	list, err := personservice.GetAll(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	resultResponse := datamodel.ResultResponse{
		Result:     true,
		DataResult: *list,
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsondata, err := json.Marshal(&resultResponse)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//取得單筆資料
func GetPerson(w http.ResponseWriter, r *http.Request) {
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req personmodel.GetPersonRequest
	err := decoder.Decode(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	err = req.CheckValue()
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	data, err := personservice.GetPerson(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	result := true
	if data == nil {
		result = false
	}
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: data,
	}
	jsondata, err := json.Marshal(resultResponse)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//新增資料
func AddPerson(w http.ResponseWriter, r *http.Request) {
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	var req personmodel.AddPersonRequest
	reqString, err := ioutil.ReadAll(r.Body)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(reqString, &req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	err = req.CheckValue()
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	result, err := personservice.AddPerson(req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, err := json.Marshal(resultResponse)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//更新資料
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	datacontroller.ControllerCrossDomain(w, r)
	decoder := json.NewDecoder(r.Body)
	var req personmodel.UpdatePersonRequest
	err := decoder.Decode(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	err = req.CheckValue()
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	result, err := personservice.UpdatePerson(req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, err := json.Marshal(resultResponse)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//刪除資料
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req personmodel.DeletePersonRequest
	err := decoder.Decode(&req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	err = req.CheckValue()
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	result, err := personservice.DeletePerson(req)
	if datacontroller.ErrorHandler(err, w) {
		return
	}
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}
