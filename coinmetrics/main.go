package coinmetrics

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
	"github.com/mitchellh/mapstructure"
)

var limit int32 = -1
var defaultPageSize int32 = 100

// CoinMetrics struct contains client object
type CoinMetrics struct {
	*api.ClientWithResponses
}

type InstitutionMetrics struct {
	Institution string  `json:"institution"`
	Time        string  `json:"time"`
	TotalAssets float64 `json:"total_assets"`
}

// InitClient will accept endpoint and apikey as parameter and it will return CoinMetrics struct which allows to access cliebt object.
func InitClient(endpoint, apiKey string) (CoinMetrics, error) {
	// TODO: Detect endpoint and based on that client option like api key should applied
	client, err := api.NewClientWithResponses(fmt.Sprintf(`%s%s/`, endpoint, constants.API_VERSION), addClientOptions(apiKey))
	if err != nil {
		return CoinMetrics{}, err
	}
	return CoinMetrics{client}, nil
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

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if limit != -1 {
				if pageSize == 0 {
					break
				}
				if defaultPageSize > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(defaultPageSize)
					params.PageSize = &defaultSize
					pageSize = pageSize - int32(defaultSize)
				}
			}
			res, err := c.GetTimeseriesMarketImpliedVolatilityWithResponse(ctx, params, reqEditors...)
			if err != nil {
				marketImpliedVolatilityError <- err
				break
			}
			if res.JSON200 != nil && len(res.JSON200.Data) > 0 {
				params.NextPageToken = res.JSON200.NextPageToken
				for _, data := range res.JSON200.Data {
					marketImpliedVolatility <- data
				}
				continue
			}
			marketImpliedVolatilityError <- errors.New(constants.NO_DATA_FOUND)
			break
		}
		close(marketImpliedVolatility)
		close(marketImpliedVolatilityError)
	}()

	return marketImpliedVolatility, marketImpliedVolatilityError
}

/*
	GetTimeseriesInstitutionMetricsWithResponseAsync To get all records over channel for institution metrics
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesPairMetrics
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
*/
func (c CoinMetrics) GetTimeseriesInstitutionMetricsWithResponseAsync(ctx context.Context, params *api.GetTimeseriesInstitutionMetricsParams, reqEditors ...api.RequestEditorFn) (chan api.InstitutionMetricsResponse, chan error) {
	institutionMetricsResponse := make(chan api.InstitutionMetricsResponse)
	institutionMetricsError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if limit != -1 {
				if pageSize == 0 {
					break
				}
				if defaultPageSize > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(defaultPageSize)
					params.PageSize = &defaultSize
					pageSize = pageSize - int32(defaultSize)
				}
			}
			res, err := c.GetTimeseriesInstitutionMetricsWithResponse(ctx, params, reqEditors...)
			if err != nil {
				institutionMetricsError <- err
				break
			}
			if res.JSON200 != nil {
				params.NextPageToken = res.JSON200.NextPageToken
				var arr []InstitutionMetrics
				err := mapstructure.Decode(res.JSON200.Data, &arr)
				if err != nil {
					institutionMetricsError <- err
					break
				}
				// for _, data := range res.JSON200.Data.(map[string]).(string) {
				// 	institutionMetricsResponse <- data
				// }
				continue
			}
			institutionMetricsError <- errors.New(constants.NO_DATA_FOUND)
			break
		}
		close(institutionMetricsResponse)
		close(institutionMetricsError)
	}()

	return institutionMetricsResponse, institutionMetricsError
}

func (c *CoinMetrics) Limit(l int32) {
	limit = l
}

func addClientOptions(apiKey string) api.ClientOption {
	addApiKey := func(ctx context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Add(constants.PARAMS_API_KEY, apiKey)
		req.URL.RawQuery = q.Encode()
		return nil
	}
	return api.WithRequestEditorFn(addApiKey)
}
