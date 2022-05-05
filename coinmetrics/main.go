package coinmetrics

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
	"go.uber.org/ratelimit"
)

var limit int32 = -1
var rl ratelimit.Limiter

// CoinMetrics struct contains client object
type CoinMetrics struct {
	*api.ClientWithResponses
}

// InitClient will accept endpoint and apikey as parameter and it will return CoinMetrics struct which allows to access client object.
func InitClient(endpoint, apiKey string) (CoinMetrics, error) {
	var client *api.ClientWithResponses
	var err error
	clientOptions := addClientOptions(apiKey)
	client, err = api.NewClientWithResponses(fmt.Sprintf(`%s%s/`, endpoint, constants.API_VERSION), clientOptions)
	if err != nil {
		return CoinMetrics{}, err
	}
	return CoinMetrics{client}, nil
}

/*
	GetTimeseriesMarketImpliedVolatilityWithResponseSync To get all records for market volatility response
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketImpliedVolatility
	Returning: api.GetTimeseriesMarketImpliedVolatilityResponse, error
*/
func (c CoinMetrics) GetTimeseriesMarketImpliedVolatilityWithResponseSync(ctx context.Context, params *api.GetTimeseriesMarketImpliedVolatilityParams, reqEditors ...api.RequestEditorFn) (api.GetTimeseriesMarketImpliedVolatilityResponse, error) {
	var response api.GetTimeseriesMarketImpliedVolatilityResponse
	var responseError error
	marketImpliedVolatility := make(chan api.MarketImpliedVolatility)
	marketImpliedVolatilityError := make(chan error)

	var pageSize int32 = limit
	go func() {
		defer close(marketImpliedVolatility)
		defer close(marketImpliedVolatilityError)
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if constants.DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(constants.DEFAULT_PAGE_SIZE)
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
				if pageSize == 0 {
					break
				}
			} else {
				// Adding other errors to response maintain orignal implemenation of api
				if res.JSON400 != nil {
					response.JSON400 = res.JSON400
				} else if res.JSON401 != nil {
					response.JSON401 = res.JSON401
				} else if res.JSON403 != nil {
					response.JSON403 = res.JSON403
				}
				marketImpliedVolatilityError <- err
				break
			}
		}
	}()

	var i int64 = 0
	for {
		select {
		case record := <-marketImpliedVolatility:
			response.JSON200.Data[i] = record
			i++
		case responseError = <-marketImpliedVolatilityError:
			return response, responseError
		}
	}
}

/*
	GetTimeseriesInstitutionMetricsWithResponseSync To get all records for institution metrics
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesInstitutionMetrics
	Returning: []interface{}, error
*/
func (c CoinMetrics) GetTimeseriesInstitutionMetricsWithResponseSync(ctx context.Context, params *api.GetTimeseriesInstitutionMetricsParams, reqEditors ...api.RequestEditorFn) ([]interface{}, error) {
	var response []interface{}
	var responseError error
	institutionMetricsResponse := make(chan interface{})
	institutionMetricsError := make(chan error)

	var pageSize int32 = limit
	go func() {
		defer close(institutionMetricsResponse)
		defer close(institutionMetricsError)
		for {
			// This condition will trigger when limit is set
			if limit != -1 {
				if constants.DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(constants.DEFAULT_PAGE_SIZE)
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
				if pageSize == 0 {
					break
				}
			} else {
				institutionMetricsError <- errors.New(constants.NO_DATA_FOUND)
				break
			}
		}
	}()

	for {
		select {
		case record := <-institutionMetricsResponse:
			response = append(response, record)
		case responseError = <-institutionMetricsError:
			return response, responseError
		}
	}
}

/*
	GetTimeseriesMarketOpenInteresetWithResponseSync To get all records for market open interest
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketOpenIntereset
	Returning: api.GetTimeseriesMarketOpenInteresetResponse, error
*/
func (c CoinMetrics) GetTimeseriesMarketOpenInteresetWithResponseSync(ctx context.Context, params *api.GetTimeseriesMarketOpenInteresetParams, reqEditors ...api.RequestEditorFn) (api.GetTimeseriesMarketOpenInteresetResponse, error) {
	var response api.GetTimeseriesMarketOpenInteresetResponse
	var responseError error
	marketOpenInterest := make(chan api.MarketOpenInterest)
	marketOpenInterestError := make(chan error)
	var pageSize int32 = limit

	go func() {
		defer close(marketOpenInterest)
		defer close(marketOpenInterestError)
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if constants.DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(constants.DEFAULT_PAGE_SIZE)
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
				if pageSize == 0 {
					break
				}
			} else {
				if res.JSON400 != nil {
					response.JSON400 = res.JSON400
				} else if res.JSON401 != nil {
					response.JSON401 = res.JSON401
				} else if res.JSON403 != nil {
					response.JSON403 = res.JSON403
				}
				marketOpenInterestError <- err
				break
			}
		}
	}()

	var i int64 = 0
	for {
		select {
		case record := <-marketOpenInterest:
			response.JSON200.Data[i] = record
			i++
		case responseError = <-marketOpenInterestError:
			return response, responseError
		}
	}
}

/*
	GetTimeseriesMarketGreeksWithResponseSync To get market greeks
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getTimeseriesMarketGreeks
	Returning: api.GetTimeseriesMarketGreeksResponse, error
*/
func (c CoinMetrics) GetTimeseriesMarketGreeksWithResponseSync(ctx context.Context, params *api.GetTimeseriesMarketGreeksParams, reqEditors ...api.RequestEditorFn) (api.GetTimeseriesMarketGreeksResponse, error) {
	var response api.GetTimeseriesMarketGreeksResponse
	var responseError error
	marketGreeks := make(chan api.MarketGreeks)
	marketGreeksError := make(chan error)

	var pageSize int32 = limit
	go func() {
		defer close(marketGreeks)
		defer close(marketGreeksError)
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if constants.DEFAULT_PAGE_SIZE > pageSize {
					cc := api.PageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.PageSize(constants.DEFAULT_PAGE_SIZE)
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
				if res.JSON400 != nil {
					response.JSON400 = res.JSON400
				} else if res.JSON401 != nil {
					response.JSON401 = res.JSON401
				} else if res.JSON403 != nil {
					response.JSON403 = res.JSON403
				}
				marketGreeksError <- err
				break
			}
		}
	}()

	var i int64 = 0
	for {
		select {
		case record := <-marketGreeks:
			response.JSON200.Data[i] = record
			i++
		case responseError = <-marketGreeksError:
			return response, responseError
		}
	}
}

/*
	GetMempoolFeeratesWithResponseSync To get mempool feerates
 	ApiEndpoint: https://docs.coinmetrics.io/api/v4#operation/getMempoolFeerates
	Returning: api.GetMempoolFeeratesResponse, error
*/
func (c CoinMetrics) GetMempoolFeeratesWithResponseSync(ctx context.Context, params *api.GetMempoolFeeratesParams, reqEditors ...api.RequestEditorFn) (api.GetMempoolFeeratesResponse, error) {
	var response api.GetMempoolFeeratesResponse
	var responseError error
	marketGreeks := make(chan api.MempoolFeerate)
	marketGreeksError := make(chan error)

	var pageSize int32 = limit
	go func() {
		for {
			// This condition will trigger when limit is set
			if pageSize != -1 {
				if constants.DEFAULT_PAGE_SIZE > pageSize {
					cc := api.MempoolFeeratesPageSize(pageSize)
					pageSize = 0
					params.PageSize = &cc
				} else {
					defaultSize := api.MempoolFeeratesPageSize(constants.DEFAULT_PAGE_SIZE)
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
			} else {
				if res.JSON400 != nil {
					response.JSON400 = res.JSON400
				} else if res.JSON401 != nil {
					response.JSON401 = res.JSON401
				} else if res.JSON403 != nil {
					response.JSON403 = res.JSON403
				}
				marketGreeksError <- err
				break
			}
		}
		close(marketGreeks)
		close(marketGreeksError)
	}()

	var i int64 = 0
	for {
		select {
		case record := <-marketGreeks:
			response.JSON200.Data[i] = record
			i++
		case responseError = <-marketGreeksError:
			return response, responseError
		}
	}
}

// Limit you can set global limit for Sync method to get particular number of records
func (c *CoinMetrics) Limit(l int32) {
	limit = l
}

func addClientOptions(apiKey string) api.ClientOption {
	var clientOptions api.ClientOption
	rl = ratelimit.New(100)
	rateLimit := func(ctx context.Context, req *http.Request) error {
		rl.Take()
		return nil
	}
	clientOptions = api.WithRequestEditorFn(rateLimit)
	if apiKey == `` {
		return clientOptions
	}
	addApiKey := func(ctx context.Context, req *http.Request) error {
		q := req.URL.Query()
		q.Add(constants.PARAMS_API_KEY, apiKey)
		req.URL.RawQuery = q.Encode()
		return nil
	}
	clientOptions = api.WithRequestEditorFn(addApiKey)
	return clientOptions
}
