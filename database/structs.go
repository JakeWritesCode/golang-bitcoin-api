// Scructs contains structs of objects to go to / from the DB

package database

import "time"

// PriceData is a struct to represent each line in the import csv file.
type PriceData struct {
	Timestamp      time.Time `db_column:"timestamp"`
	Open           float64   `db_column:"open"`
	High           float64   `db_column:"high"`
	Low            float64   `db_column:"low"`
	Close          float64   `db_column:"close"`
	VolumeBTC      float64   `db_column:"volume_btc"`
	VolumeCurrency float64   `db_column:"volume_currency"`
	WeightedPrice  float64   `db_column:"weighted_price"`
}
