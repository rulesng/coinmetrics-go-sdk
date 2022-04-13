package coinmetrics

import (
	"context"
	"errors"
	"fmt"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
)

type CoinMetrics struct {
	*api.ClientWithResponses
}

func InitClient(endpoint string) CoinMetrics {
	client, err := api.NewClientWithResponses(fmt.Sprintf(`%s/%s/`, endpoint, constants.API_VERSION))
	if err != nil {
		return CoinMetrics{}
	}
	return CoinMetrics{client}
}

/*
	GetTimeseriesMarketImpliedVolatilityWithResponseAsync To get all records over channel for market volatility response
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketImpliedVolatility
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
*/
func (c CoinMetrics) GetTimeseriesMarketImpliedVolatilityWithResponseAsync(ctx context.Context, params *api.GetTimeseriesMarketImpliedVolatilityParams, reqEditors ...api.RequestEditorFn) (chan api.MarketImpliedVolatility, chan error) {
	marketImpliedVolatility := make(chan api.MarketImpliedVolatility)
	marketImpliedVolatilityError := make(chan error)

	go func() {
		for {
			res, err := c.GetTimeseriesMarketImpliedVolatilityWithResponse(ctx, params, reqEditors...)
			if err != nil {
				marketImpliedVolatilityError <- err
				close(marketImpliedVolatility)
				close(marketImpliedVolatilityError)
				break
			}
			if res.JSON200 != nil && len(res.JSON200.Data) > 0 {
				params.NextPageToken = res.JSON200.NextPageToken
				*params.PageSize++

				for _, data := range res.JSON200.Data {
					marketImpliedVolatility <- data
				}
				continue
			}
			marketImpliedVolatilityError <- errors.New(`no data found`)
			close(marketImpliedVolatility)
			close(marketImpliedVolatilityError)
			break
		}
	}()

	return marketImpliedVolatility, marketImpliedVolatilityError
}
