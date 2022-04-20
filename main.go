package main

import (
	"context"
	"fmt"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
)

// Init is used to initialize client object with the given endpoint and apikey
func Init(endpoint, apikey string) (coinmetrics.CoinMetrics, error) {
	return coinmetrics.InitClient(endpoint, apikey)
}

func main() {
	c, err := Init(`https://api.coinmetrics.io/`, ``)
	if err != nil {
		panic(err)
	}
	mk := api.GetTimeseriesInstitutionMetricsParams{}
	c.Limit(210)
	res, errr := c.GetTimeseriesInstitutionMetricsWithResponseAsync(context.Background(), &mk)
	for {
		select {
		case rr := <-res:
			fmt.Println(`---------data-------`, rr)
		case errc := <-errr:
			fmt.Println("error", errc)
			return
		}
	}
}
