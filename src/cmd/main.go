package main

import (
	"go-basic-calculator-backend-api/src/endpointhandlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
