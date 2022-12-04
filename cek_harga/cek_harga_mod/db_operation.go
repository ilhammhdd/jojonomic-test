package cek_harga_mod

import (
	"context"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func SelectLatestHarga() (*model.Harga, error) {
	row := utils.PGPool.QueryRow(context.Background(), "SELECT * FROM harga WHERE created_at = (SELECT MAX(created_at) FROM harga)")
	var harga model.Harga
	if err := row.Scan(&harga.ID, &harga.CreatedAt, &harga.AdminID, &harga.Topup, &harga.Buyback); err != nil {
		return nil, err
	}

	return &harga, nil
}

func SelectLastInsertedHarga(lastInsertedHargaID int64) (*model.Harga, error) {
	var harga model.Harga
	if err := utils.PGPool.QueryRow(context.Background(), "SELECT * FROM harga WHERE id = $1", lastInsertedHargaID).Scan(&harga.ID, &harga.CreatedAt, &harga.AdminID, &harga.Topup, &harga.Buyback); err != nil {
		return nil, err
	}
	return &harga, nil
}
