// Scructs contains structs of objects to go to / from the DB

package database

import "time"

// PriceData is a struct to represent each line in the import csv file.
type PriceData struct {
	Timestamp      time.Time `name:"timestamp"`
	Open           float64   `name:"open"`
	High           float64   `name:"high"`
	Low            float64   `name:"low"`
	Close          float64   `name:"close"`
	VolumeBTC      float64   `name:"volumebtc"`
	VolumeCurrency float64   `name:"volumecurrency"`
	WeightedPrice  float64   `name:"weightedprice"`
}
