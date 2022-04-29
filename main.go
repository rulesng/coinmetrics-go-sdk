package main

import (
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
)

// Init is used to initialize client object with the given endpoint and apikey
func Init(endpoint, apikey string) (coinmetrics.CoinMetrics, error) {
	return coinmetrics.InitClient(endpoint, apikey)
}
