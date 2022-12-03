package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("topup")
	utils.OpenDBConnection()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/topup", HandleTopup).Methods(http.MethodPost)
	utils.StartServerMux(r)
}
