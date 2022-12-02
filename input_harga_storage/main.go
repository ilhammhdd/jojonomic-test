package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello input_harga_storage")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello input_harga_storage"))
	})
	server := http.Server{Addr: ":9998", Handler: r}
	defer server.Close()
	server.ListenAndServe()
}
