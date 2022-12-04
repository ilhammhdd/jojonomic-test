package main

import (
	"context"
	"errors"
	"time"

	"github.com/ilhammhdd/jojonomic_test/utils"
	"github.com/jackc/pgx/v5"
)

var ErrNoTransaksi = errors.New("no transaksi exists from specified start and end datetime")

func SelectTransactions(start, end time.Time) ([]*outTransaction, error) {
	rows, err := utils.PGPool.Query(context.Background(), "SELECT t.*,h.topup,h.buyback FROM transaksi t JOIN harga h ON t.harga_id = h.id WHERE t.created_at >= $1 AND t.created_at <= $2 ORDER BY t.created_at", start, end)
	if err != nil && err == pgx.ErrNoRows {
		return nil, ErrNoTransaksi
	} else if err != nil {
		return nil, err
	}

	var id int64
	var createdAt *time.Time
	var saldo float64
	var gram float64
	var typ string
	var hargaID int64
	var rekeningID int64
	var hargaTopup int64
	var hargaBuyback int64

	var transactions []*outTransaction
	_, err = pgx.ForEachRow(rows, []any{&id, &createdAt, &saldo, &gram, &typ, &hargaID, &rekeningID, &hargaTopup, &hargaBuyback}, func() error {
		transactions = append(transactions, &outTransaction{Date: createdAt.UTC().Unix(), Type: typ, Gram: gram, HargaTopup: hargaTopup, HargaBuyback: hargaBuyback, Saldo: saldo})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
