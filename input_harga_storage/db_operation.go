package main

import (
	"context"
	"log"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func InsertHarga(harga *model.Harga) error {
	cmdTag, err := utils.PGPool.Exec(context.Background(), "INSERT INTO harga (admin_id,topup,buyback) VALUES($1,$2,$3);", harga.AdminID, harga.Topup, harga.Buyback)
	if err != nil {
		log.Println(cmdTag.String(), err)
		return err
	}
	return nil
}
