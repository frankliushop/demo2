package main

import (
	"fmt"
	"myproject1/datacommon"
	dataconfig "myproject1/dataconfig"
	"myproject1/datacontroller"
	"myproject1/datacontroller/personcontroller"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

func globalPanicHandler() {
	if r := recover(); r != nil {
		switch r.(type) {
		case string:
			log.Println(r)
		case datacommon.ExceptionData:
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			jsondata, _ := json.Marshal(&r)
			jsonString := string(jsondata)
			log.Println(jsonString)
		}

	}
}

func main() {
	defer globalPanicHandler()
	r := mux.NewRouter()
	r.HandleFunc("/", datacontroller.Home).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetAll", personcontroller.GetAll).Methods("POST", "OPTIONS")
	r.HandleFunc("/GetPerson", personcontroller.GetPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/AddPerson", personcontroller.AddPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdatePerson", personcontroller.UpdatePerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/DeletePerson", personcontroller.DeletePerson).Methods("POST", "OPTIONS")
	port := fmt.Sprintf(":%d", dataconfig.GlobalConfigData.Port)
	http.ListenAndServe(port, r)
}
