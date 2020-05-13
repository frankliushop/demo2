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
	defer datacontroller.ControllerErrorHandler(w, r)
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req personmodel.GetAllRequest
	decoder.Decode(&req)
	req.CheckValue()
	list := personservice.GetAll(&req)
	resultResponse := datamodel.ResultResponse{
		Result:     true,
		DataResult: *list,
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsondata, _ := json.Marshal(&resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//取得單筆資料
func GetPerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	datacontroller.ControllerCrossDomain(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req personmodel.GetPersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	data := personservice.GetPerson(&req)
	result := true
	if data == nil {
		result = false
	}
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: data,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//新增資料
func AddPerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	datacontroller.ControllerCrossDomain(w, r)
	var req personmodel.AddPersonRequest
	reqString, _ := ioutil.ReadAll(r.Body)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(reqString, &req)
	req.CheckValue()
	result := personservice.AddPerson(req)
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//更新資料
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	datacontroller.ControllerCrossDomain(w, r)
	decoder := json.NewDecoder(r.Body)
	var req personmodel.UpdatePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := personservice.UpdatePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

//刪除資料
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	datacontroller.ControllerCrossDomain(w, r)
	decoder := json.NewDecoder(r.Body)
	var req personmodel.DeletePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := personservice.DeletePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result:     result,
		DataResult: nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}
