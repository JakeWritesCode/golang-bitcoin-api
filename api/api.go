package api

import (
	"fmt"
	"golang-bitcoin-api/config"
	"golang-bitcoin-api/utils"
	"io/ioutil"
	"net/http"
	"time"
)

// BitcoinPriceData is the model structure for storing Bitcoin price data. Modelled on the response data from
// the CryptoCompare API.
type BitcoinPriceData struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	}

	Time       time.Time `json:"time"`
	PriceHigh  float64   `json:"high"`
	PriceLow   float64   `json:"low"`
	Open       float64   `json:"open"`
	VolumeFrom float64   `json:"volumefrom"`
	VolumeTo   float64   `json:"volumeto"`
	Close      float64   `json:"close"`
}

// DownloadHistoricalData is a dummy endpoint we can hit at regular intervals to remove
func DownloadHistoricalData(w http.ResponseWriter, r *http.Request) {
	const base_download_url = "https://min-api.cryptocompare.com/data/v2/histominute?fsym=BTC&tsym=GBP&limit=2000&" +
		"toTs={time_to}&api_key={api_key}"
	const base_epoch string = "1647359128"
	var download_url string = utils.FString(base_download_url, "time_to", base_epoch, "api_key", config.CRYPTOCOMPARE_API_KEY)

	response, _ := http.Get(download_url)
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Fprintf(w, string(body))
}
