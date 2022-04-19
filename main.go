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
	mk := api.GetTimeseriesMarketGreeksParams{
		Markets: api.MarketId(`deribit-ETH-25MAR22-1200-P-option`),
	}
	c.Limit(110)
	res, errr := c.GetTimeseriesMarketGreeksWithResponseAsync(context.Background(), &mk)
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
