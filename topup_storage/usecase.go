package main

import (
	"errors"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
	"github.com/jackc/pgx/v5"
)

type DBOperator interface {
	SelectRekeningBy(norek string) (*model.Rekening, error)
	InsertRekening(norek string) (int64, error)
	InsertTransaksiWithRel(rekeningID int64, txWithRel *model.TransaksiWithRel) (int64, error)
	UpdateSaldo(*model.Rekening) (int64, error)
}

var ErrFinalBalanceNotFulfill = errors.New("final balance doesn't fulfill minimal and multiplication of 0.001")

func HandleTransaksi(txWithRel *model.TransaksiWithRel, dbo DBOperator) error {
	rek, err := dbo.SelectRekeningBy(txWithRel.Rekening.Norek)
	if err != nil && err != pgx.ErrNoRows {
		return err
	} else if rek == nil || err == pgx.ErrNoRows {
		rekeningID, err := dbo.InsertRekening(txWithRel.Rekening.Norek)
		rek = &model.Rekening{TmplCols: model.TmplCols{ID: rekeningID}, Norek: txWithRel.Rekening.Norek, Saldo: 0.000}
		if err != nil {
			return err
		}
	}

	rek.Saldo = utils.AddWithDecimalPlaces(3, rek.Saldo, txWithRel.Transaksi.Gram)
	if !utils.CheckMaxDecimalPlaces(3, rek.Saldo) {
		return ErrFinalBalanceNotFulfill
	}
	txWithRel.Rekening.Saldo = rek.Saldo

	_, err = dbo.InsertTransaksiWithRel(rek.ID, txWithRel)
	if err != nil {
		return err
	}

	_, err = dbo.UpdateSaldo(rek)
	if err != nil {
		return err
	}

	return nil
}
