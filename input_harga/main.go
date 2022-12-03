package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("input_harga")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/input-harga", HandleInputHarga).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	utils.StartServerMux(r)
}
