package cek_harga_mod

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func CekHarga() (*model.Harga, error) {
	lastInsertedHargaID, err := getLastInsertedHargaID()
	if err != nil {
		return nil, err
	}

	harga, err := SelectLatestHarga(lastInsertedHargaID)
	if err != nil {
		return nil, err
	}

	return harga, nil
}

func getLastInsertedHargaID() (int64, error) {
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/last-inserted/harga/id", utils.ENV[utils.ENV_INPUT_HARGA_STORAGE_HOST]), nil)
	if err != nil {
		return -1, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	var respBodyRaw []byte = make([]byte, resp.ContentLength)
	resp.Body.Read(respBodyRaw)
	defer resp.Body.Close()

	var respBody model.LastInsertedID
	err = json.Unmarshal(respBodyRaw, &respBody)
	if err != nil {
		return -1, err
	}

	return respBody.ID, nil
}
