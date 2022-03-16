package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang-bitcoin-api/config"
	"os"
	"time"
)

var db *sql.DB

// InitDB initialises database.
func InitDB() {
	connection, err := sql.Open("postgres", config.POSTGRES_DB_URL)
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
	var id = 0
	err := db.QueryRow("SELECT ID FROM historical_btc_data WHERE timestamp=$1", time.String()).Scan(&id)
	if err != nil {
		return false, id
	} else {
		return true, id
	}
}

// Adds a new BTC price record. Returns the id of the new record.
func AddNewBTCPriceRecord(record PriceData) int {
	// TODO - This is shit, surely there's a better way to iterate through the struct?
	base_query := "INSERT INTO historical_btc (timestamp, open, high, low, close, volume_btc, " +
		"volume_currency, weighted_price) VALUES (%s)"
	values := fmt.Sprintf(
		"%s, %s, %s, %s, %s, %s, %s, %s",
		record.Timestamp,
		record.Open,
		record.High,
		record.Low,
		record.Close,
		record.VolumeBTC,
		record.VolumeCurrency,
		record.WeightedPrice,
	)
	sql_query := fmt.Sprintf(base_query, values)
	fmt.Fprintf(os.Stdout, sql_query)
	return 1
}

// InsertBTCPriceRecordIfDoesNotExist checks to see if there's already a price record, and inserts one if not.
// Returns the id of the row in either case.
func InsertBTCPriceRecordIfDoesNotExist(record PriceData) int {
	exists, id := CheckForBTCPriceRecord(record.Timestamp)
	if !exists {

	}
	return id
}
