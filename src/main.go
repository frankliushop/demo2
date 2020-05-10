package main

import (
	"fmt"
	"net/http"
	"myproject1/datacontroller"
	"myproject1/datamodel"

	"github.com/gorilla/mux"
)

var globalData []datamodel.PersonResponse

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello word")
}

func main() {
	personHandler := datacontroller.PersonController{}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/GetAll", personHandler.GetAll).Methods("POST")
	r.HandleFunc("/GetPerson", personHandler.GetPerson).Methods("POST")
	r.HandleFunc("/AddPerson", personHandler.AddPerson).Methods("POST")
	r.HandleFunc("/UpdatePerson", personHandler.UpdatePerson).Methods("POST")
	r.HandleFunc("/DeletePerson", personHandler.DeletePerson).Methods("POST")
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	http.ListenAndServe(":8081", r)
	//http.Handle("/", r)

}
