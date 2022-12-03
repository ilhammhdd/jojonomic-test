package utils

import (
	"log"
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
)

const (
	ENV_PORT                     = "PORT"
	ENV_KAFKA_BROKER             = "KAFKA_BROKER"
	ENV_DB_HOST                  = "DB_HOST"
	ENV_DB_PORT                  = "DB_PORT"
	ENV_DB_USER                  = "DB_USER"
	ENV_DB_PASSWORD              = "DB_PASSWORD"
	ENV_DB_DBNAME                = "DB_DBNAME"
	ENV_INPUT_HARGA_STORAGE_HOST = "INPUT_HARGA_STORAGE_HOST"
)

var ENV map[string]string
var SID *shortid.Shortid
var KafkaClientConfig *sarama.Config
var PGPool *pgxpool.Pool

func InitKafkaClientConfig(clientID string) {
	clientConfig := sarama.NewConfig()
	clientConfig.ClientID = clientID
	clientConfig.Producer.Retry.Max = 5
	clientConfig.Producer.Return.Successes = true
	clientConfig.Producer.Return.Errors = true
	clientConfig.Consumer.MaxWaitTime = 300 * time.Millisecond
	KafkaClientConfig = clientConfig
}

func SetupUtils(clientID string) {
	sid, err := shortid.New(1, shortid.DefaultABC, rand.Uint64())
	if err != nil {
		log.Fatalln(err)
	}
	SID = sid
	ENV, err = godotenv.Read()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	InitKafkaClientConfig(clientID)
}
