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

func GetAll(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	list := personservice.GetAll()
	resultResponse := datamodel.ResultResponse{
		Result: true,
		Data:   list,
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsondata, _ := json.Marshal(&resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
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
		Result: result,
		Data:   data,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	var req personmodel.AddPersonRequest
	reqString, _ := ioutil.ReadAll(r.Body)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(reqString, &req)
	req.CheckValue()
	result := personservice.AddPerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	decoder := json.NewDecoder(r.Body)
	var req personmodel.UpdatePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := personservice.UpdatePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer datacontroller.ControllerErrorHandler(w, r)
	decoder := json.NewDecoder(r.Body)
	var req personmodel.DeletePersonRequest
	decoder.Decode(&req)
	req.CheckValue()
	result := personservice.DeletePerson(req)
	resultResponse := datamodel.ResultResponse{
		Result: result,
		Data:   nil,
	}
	jsondata, _ := json.Marshal(resultResponse)
	jsonString := string(jsondata)
	fmt.Fprintf(w, jsonString)
}
