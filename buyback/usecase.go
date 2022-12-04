package main

import (
	"errors"

	"github.com/ilhammhdd/jojonomic_test/cek_harga/cek_harga_mod"
	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

var ErrPriceDoesntMatch = errors.New("price doesn't match current buyback price")
var ErrBalanceInsufficient = errors.New("balance insufficient")

func Buyback(reqBody *model.TransaksiRequest) (string, error) {
	harga, err := cek_harga_mod.CekHarga()
	if err != nil {
		return "", err
	}

	if reqBody.Harga != harga.Buyback {
		return "", ErrPriceDoesntMatch
	}

	rekening, err := SelectRekening(reqBody.NoRek)
	if err != nil {
		return "", err
	}

	if rekening.Saldo < reqBody.Gram {
		return "", ErrBalanceInsufficient
	}

	transaksiWithRel := model.TransaksiWithRel{Transaksi: &model.Transaksi{Gram: reqBody.Gram, Type: model.TransaksiType_BUYBACK, HargaID: harga.ID}, Harga: harga, Rekening: &model.Rekening{Norek: reqBody.NoRek}}
	key, err := utils.Produce[model.TransaksiWithRel]("topup", &transaksiWithRel, utils.ValueByteConverterFunc[model.TransaksiWithRel](ConvertToByteVal), "buyback")
	if err != nil {
		return "", err
	}

	return key, nil
}
