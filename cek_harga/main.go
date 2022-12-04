package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("cek_harga")
	utils.OpenDBConnection()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/check-harga", HandleCekHarga).Methods(http.MethodGet)
	utils.StartServerMux(r)
}
