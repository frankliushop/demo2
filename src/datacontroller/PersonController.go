package datacontroller

import (
	datamodel "myproject1/datamodel"
	dataservice "myproject1/dataservice"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type PersonController struct {
}

func (personController *PersonController) GetAll(w http.ResponseWriter, r *http.Request) {
	defer ControllerErrorHandler(w,r)
	serviceInst := dataservice.GetPersonServiceInstance()
	list := serviceInst.GetAll()
	resultResponse := datamodel.ResultResponse{
		Result: true,
		Data:   list,
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsondata, _ := json.Marshal(&resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func (personController *PersonController) GetPerson(w http.ResponseWriter, r *http.Request) {
	defer ControllerErrorHandler(w,r)
	serviceInst := dataservice.GetPersonServiceInstance()
	decoder := json.NewDecoder(r.Body)
	var req datamodel.GetPersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	data := serviceInst.GetPerson(&req)
	result := true
	if data == nil {
		result = false
	}
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   data,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func (personController *PersonController) AddPerson(w http.ResponseWriter, r *http.Request) {
	defer ControllerErrorHandler(w,r)
	serviceInst := dataservice.GetPersonServiceInstance()
	var req datamodel.AddPersonRequest
	reqString, _ := ioutil.ReadAll(r.Body)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(reqString, &req)
	req.CheckValue()
	result := serviceInst.AddPerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func (personController *PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	defer ControllerErrorHandler(w,r)
	serviceInst := dataservice.GetPersonServiceInstance()
	decoder := json.NewDecoder(r.Body)
	var req datamodel.UpdatePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := serviceInst.UpdatePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func (personController *PersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer ControllerErrorHandler(w,r)
	serviceInst := dataservice.GetPersonServiceInstance()
	decoder := json.NewDecoder(r.Body)
	var req datamodel.DeletePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := serviceInst.DeletePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}
