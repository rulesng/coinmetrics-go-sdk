package main

import (
	"context"
	"fmt"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
)

func Init(endpoint, apikey string) coinmetrics.CoinMetrics {
	return coinmetrics.InitClient(endpoint)
}

func main() {
	c := Init(`https://community-api.coinmetrics.io/v4/`, ``)
	mk := api.GetTimeseriesMarketImpliedVolatilityParams{
		Markets: api.MarketId(`bibox-aaa-usdt-spot`),
	}
	res, err := c.GetTimeseriesMarketImpliedVolatilityWithResponseAsync(context.Background(), &mk)
	for {
		select {
		case rr := <-res:
			fmt.Println(rr)
		case errc := <-err:
			fmt.Println("received", errc)
			return
		}
	}
}
