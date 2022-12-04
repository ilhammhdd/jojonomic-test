package cek_harga_mod

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

var ErrNoPriceExists = errors.New("no price exists")

func CekHarga() (*model.Harga, error) {
	harga, err := SelectLatestHarga()
	if err != nil {
		return nil, err
	}

	lastInsertedHargaID, err := getLastInsertedHargaID()
	if err != nil {
		return nil, err
	}

	if harga == nil && lastInsertedHargaID == -1 {
		return nil, ErrNoPriceExists
	}

	if lastInsertedHargaID != -1 && lastInsertedHargaID > harga.ID {
		harga, err = SelectLastInsertedHarga(lastInsertedHargaID)
		if err != nil {
			return nil, err
		}
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
