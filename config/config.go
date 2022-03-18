package config

import (
	"fmt"
)

const (
	POSTGRES_DB_USER      string = "bitcoin_api"
	POSTGRES_DB_PASSWORD  string = "bitcoin_api_password"
	POSTGRES_DB_IP        string = "127.0.0.1"
	POSTGRES_DB_PORT      string = "5432"
	POSTGRES_DB_NAME      string = "bitcoin_api"
	CRYPTOCOMPARE_API_KEY string = "ef9d4c3f1ec774fc1058ad6446b0e36c548c6ce9647faa31167a72adc9cc0038"
)

var POSTGRES_DB_URL string = fmt.Sprintf(
	"postgresql://%s:%s@%s:%s/%s",
	POSTGRES_DB_USER,
	POSTGRES_DB_PASSWORD,
	POSTGRES_DB_IP,
	POSTGRES_DB_PORT,
	POSTGRES_DB_NAME,
)
