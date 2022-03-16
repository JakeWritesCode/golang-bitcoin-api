package config

import (
	"golang-bitcoin-api/utils"
)

const (
	POSTGRES_DB_USER      string = "bitcoin_api"
	POSTGRES_DB_PASSWORD  string = "bitcoin_api"
	POSTGRES_DB_IP        string = "localhost"
	POSTGRES_DB_PORT      string = "3456"
	POSTGRES_DB_NAME      string = "bitcoin_api"
	CRYPTOCOMPARE_API_KEY string = "ef9d4c3f1ec774fc1058ad6446b0e36c548c6ce9647faa31167a72adc9cc0038"
)

var POSTGRES_DB_URL string = utils.FString(
	"postgresql://user:password@netloc:port/dbname",
	"user", POSTGRES_DB_USER,
	"password", POSTGRES_DB_PASSWORD,
	"netloc", POSTGRES_DB_IP,
	"port", POSTGRES_DB_PORT,
	"dbname", POSTGRES_DB_NAME,
)
