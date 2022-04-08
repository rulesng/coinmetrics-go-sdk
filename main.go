package main

import (
	"context"
	"fmt"

	"github.com/Improwised/coinmetrics-go-sdk/Openapi"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
)

func Init(endpoint, apikey string) *Openapi.ClientWithResponses {

	client, err := Openapi.NewClientWithResponses(fmt.Sprintf(`%s/%s/`, endpoint, constants.API_VERSION))
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	c := Init(`https://community-api.coinmetrics.io/v4/?api_key=sdfsdf`, ``)
	fre1d := Openapi.CandleFrequency(`1d`)
	start := Openapi.StartTime(`2021-03-07`)
	end := Openapi.EndTime(`2022-03-08`)
	mk := Openapi.GetTimeseriesMarketCandlesParams{
		Markets:   Openapi.MarketId(`bibox-aaa-usdt-spot`),
		Frequency: &fre1d,
		StartTime: &start,
		EndTime:   &end,
	}
	res, err := c.GetTimeseriesMarketCandlesWithResponse(context.Background(), &mk)
	fmt.Println(err)
	fmt.Println(res.JSON200)

}
