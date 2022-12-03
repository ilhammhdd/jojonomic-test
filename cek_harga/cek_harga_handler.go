package main

import (
	"net/http"

	"github.com/ilhammhdd/jojonomic_test/cek_harga/cek_harga_mod"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleCekHarga(w http.ResponseWriter, r *http.Request) {
	harga, err := cek_harga_mod.CekHarga()
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	respJSON, err := BuildResponseBody(harga)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}
