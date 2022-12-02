CREATE TABLE IF NOT EXISTS harga(
  id BIGSERIAL NOT NULL PRIMARY KEY,
  created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  admin_id TEXT NOT NULL,
  topup BIGINT NOT NULL,
  buyback BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS rekening(
  id BIGSERIAL NOT NULL PRIMARY KEY,
  created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  norek TEXT NOT NULL UNIQUE,
  saldo NUMERIC(10,3) NOT NULL
);

CREATE TYPE transaksi_type AS ENUM('topup','buyback');

CREATE TABLE IF NOT EXISTS transaksi(
  id BIGSERIAL NOT NULL PRIMARY KEY,
  created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  saldo NUMERIC(10,3) NOT NULL,
  gram NUMERIC(10,3) NOT NULL,
  harga_id BIGINT NOT NULL REFERENCES harga (id),
  "type" transaksi_type NOT NULL
);