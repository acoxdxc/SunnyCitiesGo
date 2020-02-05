package main

import (
    "log"
    "net/http"
    "fmt"

    "github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}


func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}

func serve() {
    r := mux.NewRouter()
    r.HandleFunc("/", FlightCodes).Methods(http.MethodGet)
    r.HandleFunc("/", notFound)
    log.Fatal(http.ListenAndServe(":8080", r))
}

var flightCodes []string = multipleFlightMockAppend2()

func FlightCodes(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, flightCodes)
}