package api

import (
	"encoding/json"
	"fmt"
	"golang-bitcoin-api/config"
	"golang-bitcoin-api/database"
	"golang-bitcoin-api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// BitcoinPriceData mirrors the incoming JSON data from the API. It will be transferred to PriceData before
// saving to the database.
type BitcoinPriceData struct {
	Response   string   `json:"Response"`
	Message    string   `json:"Message"`
	HasWarning bool     `json:"HasWarning"`
	Type       int      `json:"Type"`
	RateLimit  struct{} `json:"RateLimit"`
	Data       struct {
		Aggregated bool  `json:"Aggregated"`
		TimeFrom   int64 `json:"TimeFrom"`
		TimeTo     int64 `json:"TimeTo"`
		Data       []struct {
			Time             int64   `json:"time"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Open             float64 `json:"open"`
			VolumeFrom       float64 `json:"volumefrom"`
			VolumeTo         float64 `json:"volumeto"`
			Close            float64 `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}

// DownloadHistoricalData is a dummy endpoint we can hit at regular intervals to remove
func DownloadHistoricalData(w http.ResponseWriter, r *http.Request) {
	timestamp := database.GetLatestBTCPriceRecordTimestamp() + (2000 * 60)
	fmt.Println("Start time is " + time.Unix(timestamp, 0).String())
	for {

		timestamp_unix := strconv.FormatInt(timestamp, 10)
		var download_url string = fmt.Sprintf(
			"https://min-api.cryptocompare.com/data/v2/histominute?fsym=BTC&tsym=USD&limit=2000&toTs=%s&api_key=%s",
			timestamp_unix,
			config.CRYPTOCOMPARE_API_KEY,
		)

		response, _ := http.Get(download_url)
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, string(body))

		jsonData := BitcoinPriceData{}
		json.Unmarshal(body, &jsonData)

		fmt.Println("Earliest time is " + time.Unix(jsonData.Data.Data[0].Time, 0).String())
		fmt.Println("Latest time is " + time.Unix(jsonData.Data.Data[len(jsonData.Data.Data)-1].Time, 0).String())

		for i := 0; i < len(jsonData.Data.Data); i++ {
			dbStruct := database.PriceData{}
			dbStruct.Timestamp = time.Unix(jsonData.Data.Data[i].Time, 0)
			dbStruct.Open = jsonData.Data.Data[i].Open
			dbStruct.High = jsonData.Data.Data[i].High
			dbStruct.Low = jsonData.Data.Data[i].Low
			dbStruct.Close = jsonData.Data.Data[i].Close
			dbStruct.VolumeBTC = jsonData.Data.Data[i].VolumeFrom
			database.AddNewBTCPriceRecord(dbStruct)
		}
		timestamp = database.GetLatestBTCPriceRecordTimestamp()
		if time.Now().Sub(time.Unix(timestamp, 0)).Seconds() < 600 {
			break
		}
		timestamp += 2000 * 60
	}
}

// GetBTCPriceHistoryData returns the BTC price history data between two timestamps.
func GetBTCPriceHistoryData(w http.ResponseWriter, r *http.Request) {
	fromDateList, present := r.URL.Query()["fromDate"]
	if !present || len(fromDateList) == 0 {
		utils.GenericAPI400Response(w, "You did not enter a fromDate.")
		return
	}
	fromDate := fromDateList[0]
	intFromDate, err := strconv.ParseInt(fromDate, 10, 64)
	failed := utils.GenericAPIErrorHandler(err, w, "The fromDate could not be parsed. Please enter a unix timestamp.")
	if failed {
		return
	}

	toDateList, present := r.URL.Query()["toDate"]
	if !present || len(toDateList) == 0 {
		utils.GenericAPI400Response(w, "You did not enter a to Date.")
	}
	toDate := toDateList[0]
	intToDate, err := strconv.ParseInt(toDate, 10, 64)
	failed = utils.GenericAPIErrorHandler(err, w, "The toDate could not be parsed. Please enter a unix timestamp.")
	if failed {
		return
	}

	data := database.FetchBTCDataBetweenTimestamps(time.Unix(intFromDate, 0), time.Unix(intToDate, 0))

	json, err := json.Marshal(data)
	failed = utils.GenericAPIErrorHandler(err, w, "Failed to parse data to JSON: ")
	if failed {
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(json)
	failed = utils.GenericAPIErrorHandler(err, w, "Failed to write response: ")
	if failed {
		return
	}
}
