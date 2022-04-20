package coinmetrics

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
)

var limit int32 = -1

const DEFAULT_PAGE_SIZE int32 = 100

// CoinMetrics struct contains client object
type CoinMetrics struct {
	*api.ClientWithResponses
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
		- You can call Limit() before calling this function to set limit.
*/
func (c CoinMetrics) GetTimeseriesMarketImpliedVolatilityWithResponseAsync(ctx context.Context, params *api.GetTimeseriesMarketImpliedVolatilityParams, reqEditors ...api.RequestEditorFn) (chan api.MarketImpliedVolatility, chan error) {
	marketImpliedVolatility := make(chan api.MarketImpliedVolatility)
	marketImpliedVolatilityError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if pageSize == 0 {
					break
				}
				if DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(DEFAULT_PAGE_SIZE)
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
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesInstitutionMetrics
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
*/
func (c CoinMetrics) GetTimeseriesInstitutionMetricsWithResponseAsync(ctx context.Context, params *api.GetTimeseriesInstitutionMetricsParams, reqEditors ...api.RequestEditorFn) (chan interface{}, chan error) {
	institutionMetricsResponse := make(chan interface{})
	institutionMetricsError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if limit != -1 {
				if pageSize == 0 {
					break
				}
				if DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(DEFAULT_PAGE_SIZE)
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
				if err != nil {
					institutionMetricsError <- err
					break
				}
				institutionMetricsResponse <- res.JSON200.Data
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

/*
	GetTimeseriesMarketOpenInteresetWithResponseAsync To get all records over channel for market open interest
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketOpenIntereset
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
		- You can call Limit() before calling this function to set limit.
*/
func (c CoinMetrics) GetTimeseriesMarketOpenInteresetWithResponseAsync(ctx context.Context, params *api.GetTimeseriesMarketOpenInteresetParams, reqEditors ...api.RequestEditorFn) (chan api.MarketOpenInterest, chan error) {
	marketOpenInterest := make(chan api.MarketOpenInterest)
	marketOpenInterestError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if pageSize == 0 {
					break
				}
				if DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(DEFAULT_PAGE_SIZE)
					params.PageSize = &defaultSize
					pageSize = pageSize - int32(defaultSize)
				}
			}
			res, err := c.GetTimeseriesMarketOpenInteresetWithResponse(ctx, params, reqEditors...)
			if err != nil {
				marketOpenInterestError <- err
				break
			}
			if res.JSON200 != nil && len(res.JSON200.Data) > 0 {
				params.NextPageToken = res.JSON200.NextPageToken
				for _, data := range res.JSON200.Data {
					marketOpenInterest <- data
				}
				continue
			}
			marketOpenInterestError <- errors.New(constants.NO_DATA_FOUND)
			break
		}
		close(marketOpenInterest)
		close(marketOpenInterestError)
	}()

	return marketOpenInterest, marketOpenInterestError
}

/*
	GetTimeseriesMarketGreeksWithResponseAsync To get market greeks
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketGreeks
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
		- You can call Limit() before calling this function to set limit.
*/
func (c CoinMetrics) GetTimeseriesMarketGreeksWithResponseAsync(ctx context.Context, params *api.GetTimeseriesMarketGreeksParams, reqEditors ...api.RequestEditorFn) (chan api.MarketGreeks, chan error) {
	marketGreeks := make(chan api.MarketGreeks)
	marketGreeksError := make(chan error)

	var pageSize int32 = limit
	go func() {
		defer close(marketGreeks)
		defer close(marketGreeksError)
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(DEFAULT_PAGE_SIZE)
					params.PageSize = &defaultSize
					pageSize = pageSize - int32(defaultSize)
				}
			}
			res, err := c.GetTimeseriesMarketGreeksWithResponse(ctx, params, reqEditors...)
			if err != nil {
				marketGreeksError <- err
				break
			}
			if res.JSON200 != nil && len(res.JSON200.Data) > 0 {
				params.NextPageToken = res.JSON200.NextPageToken
				for _, data := range res.JSON200.Data {
					marketGreeks <- data
				}
				if pageSize == 0 {
					break
				}
			} else {
				marketGreeksError <- errors.New(constants.NO_DATA_FOUND)
				break
			}
		}
	}()

	return marketGreeks, marketGreeksError
}

/*
	GetMempoolFeeratesWithResponseAsync To get market greeks
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getMempoolFeerates
	Returning:
		- Method will be returning two channel, one having data and one with error.
  		- This method will continuously retrive data and send over channel
  		- Check error channel if is there any, when error will occured channel will be closed.
		- You can call Limit() before calling this function to set limit.
*/
func (c CoinMetrics) GetMempoolFeeratesWithResponseAsync(ctx context.Context, params *api.GetMempoolFeeratesParams, reqEditors ...api.RequestEditorFn) (chan api.MempoolFeerate, chan error) {
	marketGreeks := make(chan api.MempoolFeerate)
	marketGreeksError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if DEFAULT_PAGE_SIZE > pageSize {
					cc := api.MempoolFeeratesPageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.MempoolFeeratesPageSize(DEFAULT_PAGE_SIZE)
					params.PageSize = &defaultSize
					pageSize = pageSize - int32(defaultSize)
				}
			}
			res, err := c.GetMempoolFeeratesWithResponse(ctx, params, reqEditors...)
			if err != nil {
				marketGreeksError <- err
				break
			}
			if res.JSON200 != nil && len(res.JSON200.Data) > 0 {
				params.NextPageToken = res.JSON200.NextPageToken
				for _, data := range res.JSON200.Data {
					marketGreeks <- data
				}
				if pageSize == 0 {
					break
				}
				continue
			}
			marketGreeksError <- errors.New(constants.NO_DATA_FOUND)
			break
		}
		close(marketGreeks)
		close(marketGreeksError)
	}()

	return marketGreeks, marketGreeksError
}

func (c *CoinMetrics) Limit(l int32) {
	limit = l
}

func addClientOptions(apiKey string) api.ClientOption {
	if apiKey == `` {
		return nil
	}
	addApiKey := func(ctx context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Add(constants.PARAMS_API_KEY, apiKey)
		req.URL.RawQuery = q.Encode()
		return nil
	}
	return api.WithRequestEditorFn(addApiKey)
}
