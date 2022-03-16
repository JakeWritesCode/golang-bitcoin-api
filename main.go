/*
golang-bitcoin-api is a small bitcoin API I have built as a toy project in order to try and learn golang.

Essentially it's just going to download from the coinmarketcap API and cache the data so I don't have to pay for a
proper API key.
*/
package main

import (
	"flag"
	"fmt"
	"golang-bitcoin-api/api"
	"golang-bitcoin-api/file_parser"
	"log"
	"net/http"
)

// homePage is a homepage endpoint for our API.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Endpoint hit: homePage")
}

// handleRequests is the main requests handler for our API.
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/download", api.DownloadHistoricalData)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	parse_flag := flag.String("parse", "none", "Parse an incoming dataset")
	flag.Parse()

	if *parse_flag != "none" {
		file_parser.ParseCSV(*parse_flag)
	} else {
		fmt.Println("Starting web server on port 10000...")
		handleRequests()
	}
}
