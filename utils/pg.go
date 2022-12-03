package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDBConnection() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", ENV[ENV_DB_HOST], ENV[ENV_DB_PORT], ENV[ENV_DB_USER], ENV[ENV_DB_PASSWORD], ENV[ENV_DB_DBNAME]))
	if err != nil {
		log.Println(err)
		return
	}
	PGPool = conn
}
