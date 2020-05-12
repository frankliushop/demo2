package main

import (
	"fmt"
	"myproject1/datacontroller/personcontroller"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}

func globalPanicHandler() {
	if r := recover(); r != nil {
		log.Println(r)
	}
}

func main() {
	defer globalPanicHandler()
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetAll", personcontroller.GetAll).Methods("POST", "OPTIONS")
	r.HandleFunc("/GetPerson", personcontroller.GetPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/AddPerson", personcontroller.AddPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdatePerson", personcontroller.UpdatePerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/DeletePerson", personcontroller.DeletePerson).Methods("POST", "OPTIONS")
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	http.ListenAndServe(":8081", r)
	//http.Handle("/", r)

}
