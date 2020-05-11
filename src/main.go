package main

import (
	"fmt"
	"myproject1/datacontroller/personcontroller"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/GetAll", personcontroller.GetAll).Methods("POST")
	r.HandleFunc("/GetPerson", personcontroller.GetPerson).Methods("POST")
	r.HandleFunc("/AddPerson", personcontroller.AddPerson).Methods("POST")
	r.HandleFunc("/UpdatePerson", personcontroller.UpdatePerson).Methods("POST")
	r.HandleFunc("/DeletePerson", personcontroller.DeletePerson).Methods("POST")
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	http.ListenAndServe(":8081", r)
	//http.Handle("/", r)

}
