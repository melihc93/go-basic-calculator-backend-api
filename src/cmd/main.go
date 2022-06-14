package main

import (
	"fmt"
	"go-basic-calculator-backend-api/src/endpointhandlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type responseObject struct {
	Result string `json:"result"`
	User   string `json:"user"`
}

type errorObject struct {
	Message string `json:"message"`
}

var responseObjects responseObject
var errorObjects errorObject

func sumResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sum response object creating...")
	query := r.URL.Query()
	fmt.Println(query["params"])
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func addResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add response object creating...")
	query := r.URL.Query()
	fmt.Println(query["params"])
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func subsResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("subtraction response object creating...")
	query := r.URL.Query()
	fmt.Println(query["params"])
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func multResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("multiplication response object creating...")
	query := r.URL.Query()
	fmt.Println(query["params"])
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/sum", endpointhandlers.SumResponse).Methods("GET")
	myRouter.HandleFunc("/add", endpointhandlers.AddResponse).Methods("POST")
	myRouter.HandleFunc("/subtraction", endpointhandlers.SubResponse).Methods("POST")
	myRouter.HandleFunc("/multiplication", endpointhandlers.MultResponse).Methods("POST")
	myRouter.HandleFunc("/division", endpointhandlers.DivResponse).Methods("POST")
	log.Fatal(http.ListenAndServe(":8383", myRouter))
}

func main() {
	handleRequests()
}
