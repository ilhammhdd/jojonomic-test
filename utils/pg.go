package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenConnection(host, port, user, password, dbname string) {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname))
	if err != nil {
		log.Println(err)
		return
	}
	PGPool = conn
}
