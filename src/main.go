package main

import (
	"fmt"
	dataconfig "myproject1/dataconfig"
	"myproject1/datacontroller"
	"myproject1/datacontroller/personcontroller"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

//處理全域panic
func globalPanicHandler() {
	if r := recover(); r != nil {
		log.Println(r)
	}
}

func main() {
	defer globalPanicHandler()
	log.Println("start program")
	//定義路由
	r := mux.NewRouter()
	r.HandleFunc("/", datacontroller.Home).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetAll", personcontroller.GetAll).Methods("POST", "OPTIONS")
	r.HandleFunc("/GetPerson", personcontroller.GetPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/AddPerson", personcontroller.AddPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdatePerson", personcontroller.UpdatePerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/DeletePerson", personcontroller.DeletePerson).Methods("POST", "OPTIONS")
	log.Println("define router")
	port := fmt.Sprintf(":%d", dataconfig.GlobalConfigData.Port)
	http.ListenAndServe(port, r)
	log.Println("program finish")
}
