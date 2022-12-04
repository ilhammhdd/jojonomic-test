package main

import (
	"context"
	"errors"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
	"github.com/jackc/pgx"
)

var ErrNoAccountExists = errors.New("no account with given norek exists")

func SelectRekening(norek string) (*model.Rekening, error) {
	var rekening model.Rekening
	err := utils.PGPool.QueryRow(context.Background(), "SELECT * FROM rekening WHERE norek = $1", norek).Scan(&rekening.ID, &rekening.CreatedAt, &rekening.Norek, &rekening.Saldo)
	if err != nil && err == pgx.ErrNoRows {
		return nil, ErrNoAccountExists
	} else if err != nil {
		return nil, err
	}
	return &rekening, nil
}
