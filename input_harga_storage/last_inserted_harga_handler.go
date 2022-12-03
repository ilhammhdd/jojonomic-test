package main

import (
	"encoding/json"
	"net/http"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

var lastInsertedHargaID model.LastInsertedID = model.LastInsertedID{ID: -1}

func HandleLastInsertedHarga(w http.ResponseWriter, r *http.Request) {
	respBody, err := json.Marshal(lastInsertedHargaID)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
