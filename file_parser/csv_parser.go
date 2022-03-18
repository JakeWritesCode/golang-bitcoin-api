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

// ParseCSV with parse a csv file from the entered path and add historical bitcoin price data to the DB.
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
		if record[0] == "Timestamp" {
			continue
		}

		// Check for NaN, if it exists, continue.
		if CheckForNaN(record) {
			continue
		}

		// Parse to struct
		var parsedObject = database.PriceData{
			Timestamp:      ParseUnixTimestamp(record[0]),
			Open:           ParseFloat(record[1]),
			High:           ParseFloat(record[2]),
			Low:            ParseFloat(record[3]),
			Close:          ParseFloat(record[4]),
			VolumeBTC:      ParseFloat(record[5]),
			VolumeCurrency: ParseFloat(record[6]),
			WeightedPrice:  ParseFloat(record[7]),
		}

		exists, id := database.CheckForBTCPriceRecord(parsedObject.Timestamp)
		if exists {
			fmt.Println("Exists!")
		} else {
			fmt.Println("New Record!" + string(rune(id)))
			id := database.AddNewBTCPriceRecord(parsedObject)
			fmt.Println("Woop woop, inserted new record, id:" + strconv.FormatInt(int64(id), 10))
		}
	}
}

// ParseFloat parses an input string to float64 with error handling.
func ParseFloat(inputString string) float64 {
	output, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

// ParseUnixTimestamp parses an input string to time.Time with error handling.
func ParseUnixTimestamp(inputString string) time.Time {
	int_input, err := strconv.ParseInt(inputString, 10, 64)
	if err != nil {
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
