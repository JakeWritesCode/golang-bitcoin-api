// A csv file parser for large quantities of historical data. So we don't have to spam the API to death.

package api

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ParseCSV with parse a csv file from the entered path and add historical bitcoin price data to the DB.
func ParseCSV(filename string) {
	// Open the csv file
	csv_raw, _ := os.ReadFile(filename)
	csv_string := string(csv_raw)
	csv := csv.NewReader(strings.NewReader(csv_string))

	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
