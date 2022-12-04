package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("buyback")
	utils.OpenDBConnection()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/buyback", HandleBuyback).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	utils.StartServerMux(r)
}
