package main

import (
	"context"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

type BuybackDBOperator struct {
	Ctx context.Context
}

func (tdo BuybackDBOperator) SelectRekeningBy(norek string) (*model.Rekening, error) {
	var rekening model.Rekening
	err := utils.PGPool.QueryRow(tdo.Ctx, "SELECT * FROM rekening WHERE norek = $1", norek).Scan(&rekening.ID, &rekening.CreatedAt, &rekening.Norek, &rekening.Saldo)
	if err != nil {
		return nil, err
	}
	return &rekening, nil
}

func (tdo BuybackDBOperator) InsertRekening(norek string) (int64, error) {
	var lastInsertedID int64
	err := utils.PGPool.QueryRow(tdo.Ctx, "INSERT INTO rekening (norek,saldo) VALUES ($1,$2) RETURNING id", norek, 0.000).Scan(&lastInsertedID)
	if err != nil {
		return -1, err
	}
	return lastInsertedID, nil
}

func (tdo BuybackDBOperator) InsertTransaksiWithRel(rekeningID int64, txWithRel *model.TransaksiWithRel) (int64, error) {
	var lastInsertedID int64
	err := utils.PGPool.QueryRow(tdo.Ctx, "INSERT INTO transaksi (saldo,gram,type,harga_id,rekening_id) VALUES ($1,$2,$3,$4,$5) RETURNING id", txWithRel.Rekening.Saldo, txWithRel.Transaksi.Gram, txWithRel.Transaksi.Type, txWithRel.Harga.ID, rekeningID).Scan(&lastInsertedID)
	if err != nil {
		return -1, err
	}
	return lastInsertedID, nil
}

func (tdo BuybackDBOperator) UpdateSaldo(rek *model.Rekening) (int64, error) {
	cmdTag, err := utils.PGPool.Exec(tdo.Ctx, "UPDATE rekening SET saldo = $1 WHERE id = $2", rek.Saldo, rek.ID)
	if err != nil {
		return -1, err
	}
	return cmdTag.RowsAffected(), nil
}
