package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("cek-mutasi")
	utils.OpenDBConnection()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/mutasi", HandleCekMutasi).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	utils.StartServerMux(r)
}
