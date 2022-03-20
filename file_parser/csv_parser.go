// A csv file parser for large quantities of historical data. So we don't have to spam the API to death.

package file_parser

import (
	"encoding/csv"
	"fmt"
	"golang-bitcoin-api/database"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// ParseCSV will parse a csv file from the entered path and add historical bitcoin price data to the DB.
func ParseCSV(filename string) {
	// Open the csv file
	fmt.Println("Parsing csv at " + filename)
	csv_raw, _ := os.ReadFile(filename)
	csv_string := string(csv_raw)
	csv_data := csv.NewReader(strings.NewReader(csv_string))

	for {
		record, err := csv_data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Skip first row
		if record[0] == "Timestamp" || record[0] == "Date" {
			continue
		}

		// Check for NaN, if it exists, continue.
		if CheckForNaN(record) {
			continue
		}

		// Parse to struct
		var parsedObject = database.PriceData{
			Timestamp: ParseTimestamp(record[0]),
			Open:      ParseFloat(record[1]),
			High:      ParseFloat(record[2]),
			Low:       ParseFloat(record[3]),
			Close:     ParseFloat(record[4]),
			VolumeBTC: ParseFloat(record[5]),
		}
		database.AddNewBTCPriceRecord(parsedObject)
	}
	fmt.Println("Finished parsing all records.")
}

// ParseFloat parses an input string to float64 with error handling.
func ParseFloat(inputString string) float64 {
	output, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return output
}

// ParseTimestamp parses an input string to time.Time with error handling.
func ParseTimestamp(inputString string) time.Time {
	int_input, err := strconv.ParseInt(inputString, 10, 64)
	if err != nil {
		// Not an int, check for date string.
		if strings.Contains(err.Error(), "invalid syntax") {
			time, err := time.Parse("02/01/2006, 15:04:05", inputString)
			if err != nil {
				log.Fatal(err.Error())
				os.Exit(1)
			}
			return time
		}
		log.Fatal(err)
	}
	output := time.Unix(int_input, 0)
	return output
}

// CheckForNaN checks for NaNs in the row, returns True if there are.
func CheckForNaN(inputRow []string) bool {
	for i := 0; i < len(inputRow); i++ {
		if inputRow[i] == "NaN" {
			return true
		}
		parsedFloat, _ := strconv.ParseFloat(inputRow[i], 64)
		if parsedFloat == math.NaN() {
			return true
		}
	}
	return false
}
