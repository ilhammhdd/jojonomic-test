package main

import (
	"context"
	"log"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func InsertHarga(harga *model.Harga) (int64, error) {
	row := utils.PGPool.QueryRow(context.Background(), "INSERT INTO harga (admin_id,topup,buyback) VALUES($1,$2,$3) RETURNING id", harga.AdminID, harga.Topup, harga.Buyback)

	var lastInsertedID int64
	if err := row.Scan(&lastInsertedID); err != nil {
		log.Println(err)
		return -1, err
	}

	return lastInsertedID, nil
}
