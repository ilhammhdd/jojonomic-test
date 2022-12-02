package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
)

func init() {
	sid, err := shortid.New(1, shortid.DefaultABC, rand.Uint64())
	if err != nil {
		log.Fatalln(err)
	}
	utils.SID = sid
	utils.ENV, err = godotenv.Read()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/input-harga", HandleInputHarga).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	utils.StartServerMux(r)
}
