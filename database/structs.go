// Scructs contains structs of objects to go to / from the DB

package database

import "time"

// PriceData is a struct to represent each line in the import csv file.
type PriceData struct {
	Timestamp time.Time `name:"timestamp" json:"timestamp"`
	Open      float64   `name:"open" json:"open"`
	High      float64   `name:"high" json:"high"`
	Low       float64   `name:"low" json:"low"`
	Close     float64   `name:"close" json:"close"`
	VolumeBTC float64   `name:"volume_btc" db:"volume_btc" json:"volumebtc"`
}
