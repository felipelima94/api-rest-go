package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Info struct {
	Text string `json:"info"`
}

type Custumer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
	Info *Info  `json:"info"`
}

var infos []Info

var customers []Custumer

func fncHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infos)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	infos = append(infos, Info{Text: "Felipe"})

	r.HandleFunc("/hello", fncHello).Methods("GET")

	r.HandleFunc("/custumer", getCustomer).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
