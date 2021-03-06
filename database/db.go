package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang-bitcoin-api/config"
	"os"
	"time"
)

var db *sqlx.DB

// InitDB initialises database.
func InitDB() {
	connection, err := sqlx.Connect("postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable",
			config.POSTGRES_DB_USER,
			config.POSTGRES_DB_PASSWORD,
			config.POSTGRES_DB_NAME,
		),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db = connection
}

// Create and update Historical BTC price data.

// CheckForBTCPriceRecord checks to see if there is already a record against the given timestamp.
// It returns true and the ID if the record exists, false and 0 if not.
func CheckForBTCPriceRecord(time time.Time) (bool, int) {
	var id int
	err := db.Get(&id, "SELECT id FROM historical_btc_data WHERE timestamp=$1", time)
	if err != nil {
		return false, 0
	} else {
		return true, id
	}
}

// AddNewBTCPriceRecord adds a new BTC price record.
func AddNewBTCPriceRecord(record PriceData) {
	statement := `
		INSERT INTO historical_btc_data
		(timestamp, open, high, low, close, volume_btc, volume_currency, weighted_price)
		SELECT :timestamp, :open, :high, :low, :close, :volumebtc, :volumecurrency, :weightedprice
		WHERE
			NOT EXISTS (
					SELECT timestamp FROM historical_btc_data WHERE timestamp = :timestamp
				);
	`

	_, err := db.NamedExec(statement, record)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// GetLatestBTCPriceRecordTimestamp gets the latest BTC price record timestamp as Unix.
func GetLatestBTCPriceRecordTimestamp() int64 {
	var timestamp = time.Time{}
	db.Get(&timestamp, "SELECT timestamp FROM historical_btc_data ORDER BY timestamp DESC LIMIT 1")
	fmt.Println("Latest DB time is " + timestamp.String())
	return timestamp.Unix()
}

// FetchBTCDataBetweenTimestamps fetches bitcoin price data between two timestamps.
// Returns []PriceData
func FetchBTCDataBetweenTimestamps(fromTime time.Time, toTime time.Time) []PriceData {
	records := []PriceData{}
	fmt.Println(fromTime, toTime)
	err := db.Select(&records, `SELECT timestamp, open, high, close, low, volume_btc from historical_btc_data
		WHERE timestamp BETWEEN $1 AND $2`, fromTime, toTime)
	fmt.Println(len(records))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("%s records fetched.", len(records)))
	return records
}
