package v4_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var _coinmetrics coinmetrics.CoinMetrics

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var _err error
	_coinmetrics, _err = coinmetrics.InitClient(constants.TEST_ENDPOINT, constants.TEST_KEY)
	if _err != nil {
		panic(_err)
	}
	os.Exit(m.Run())
}

func TestAssetNotFoundForGetCatalogAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetsWithoutParams(t *testing.T) {
	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]},{"asset":"10set","full_name":"Tenset","exchanges":["gate.io","lbank"],"markets":["gate.io-10set-usdt-spot","lbank-10set-usdt-spot"]},{"asset":"18c","full_name":"Block 18","exchanges":["huobi"],"markets":["huobi-18c-btc-spot","huobi-18c-eth-spot"]},{"asset":"1art","full_name":"ArtWallet","exchanges":["gate.io"],"markets":["gate.io-1art-usdt-spot"]},{"asset":"1box","full_name":"1BOX","exchanges":["zb.com"],"markets":["zb.com-1box-usdt-spot"]},{"asset":"1earth","full_name":"EarthFund","exchanges":["gate.io","kucoin"],"markets":["gate.io-1earth-usdt-spot","kucoin-1earth-usdt-spot"]}]}`)
	param := api.GetCatalogAssetsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetsWithParams(t *testing.T) {
	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]}]}`)
	param := api.GetCatalogAssetsParams{
		Assets:  &api.CatalogAssetId{`100x`},
		Include: &api.CatalogAssetIncludeFields{`markets`, `exchanges`},
		Exclude: &api.CatalogAssetExcludeFields{`metrics`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

func TestAssetNotFoundForGetCatalogMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'metrics'. Value 'asdgwav' is not supported.`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`asdgwav`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMetricsWithoutParamsResponse(t *testing.T) {
	data := getCatalogMetricsResponse(`{"data":[{"metric":"AdrActCnt","full_name":"Addresses, active, count","description":"The sum count of unique addresses that were active in the network (either as a recipient or originator of a ledger change) that interval. All parties in a ledger change action (recipients and originators) are counted. Individual addresses are not double-counted if previously active.","category":"Addresses","subcategory":"Active","unit":"Addresses","data_type":"bigint","type":"Sum","frequencies":[{}],"display_name":"Active Addr Cnt"}]}`)
	param := api.GetCatalogMetricsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMetricsWithParamsResponse(t *testing.T) {
	data := getCatalogMetricsResponse(`{"data":[{"metric":"AdrActCnt","full_name":"Addresses, active, count","description":"The sum count of unique addresses that were active in the network (either as a recipient or originator of a ledger change) that interval. All parties in a ledger change action (recipients and originators) are counted. Individual addresses are not double-counted if previously active.","category":"Addresses","subcategory":"Active","unit":"Addresses","data_type":"bigint","type":"Sum","frequencies":[{}],"display_name":"Active Addr Cnt"}]}`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`AdrActCnt`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange

func TestExchangeNotFoundForGetCatalogExchangesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'id'. Exchange 'sdvwbtc' is not supported.`)
	param := api.GetCatalogExchangesParams{
		Exchanges: &api.CatalogExchangeId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangesWithoutParams(t *testing.T) {
	data := getCatalogExchangeResponse(`{"data":[{"exchange":"bitbank","markets":["bitbank-bch-btc-spot","bitbank-bch-jpy-spot","bitbank-btc-jpy-spot","bitbank-eth-btc-spot","bitbank-eth-jpy-spot","bitbank-ltc-btc-spot","bitbank-ltc-jpy-spot","bitbank-mona-btc-spot","bitbank-mona-jpy-spot","bitbank-xrp-btc-spot","bitbank-xrp-jpy-spot"],"min_time":"2017-02-14T03:26:18.512000000Z","max_time":"2022-04-25T07:01:12.566000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2017-02-14T04:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2017-02-15T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]},{"exchange":"itbit","markets":["itbit-aave-usd-spot","itbit-bch-usd-spot","itbit-btc-eur-spot","itbit-btc-sgd-spot","itbit-btc-usd-spot","itbit-eth-eur-spot","itbit-eth-sgd-spot","itbit-eth-usd-spot","itbit-link-usd-spot","itbit-ltc-usd-spot","itbit-matic-usd-spot","itbit-paxg-usd-spot","itbit-uni-usd-spot"],"min_time":"2019-03-13T07:53:05.963000000Z","max_time":"2022-04-25T07:01:02.866000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2019-03-13T08:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2019-03-14T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]}]}
	`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangesWithParams(t *testing.T) {
	data := getCatalogExchangeResponse(`{"data":[{"exchange":"bitbank","markets":["bitbank-bch-btc-spot","bitbank-bch-jpy-spot","bitbank-btc-jpy-spot","bitbank-eth-btc-spot","bitbank-eth-jpy-spot","bitbank-ltc-btc-spot","bitbank-ltc-jpy-spot","bitbank-mona-btc-spot","bitbank-mona-jpy-spot","bitbank-xrp-btc-spot","bitbank-xrp-jpy-spot"],"min_time":"2017-02-14T03:26:18.512000000Z","max_time":"2022-04-25T07:01:12.566000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2017-02-14T04:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2017-02-15T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangesParams{
		Exchanges: &api.CatalogExchangeId{`bitbank`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogExchangesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange Asset Pairs

func TestExchangeAssetsNotFoundForGetCatalogExchangeAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogExchangeAssetsParams{
		ExchangeAssets: &api.CatalogExchangeAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangeAssetsWithoutParams(t *testing.T) {
	data := getCatalogExchangeAssetsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]},{"exchange_asset":"coinbase-eth","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-11T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-10T19:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangeAssetsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangeAssetsWithParams(t *testing.T) {
	data := getCatalogExchangeAssetsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangeAssetsParams{
		ExchangeAssets: &api.CatalogExchangeAssetId{`binance-btc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogExchangeAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange Asset Pairs

func TestPairsNotFoundForGetCatalogAssetPairsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetPairsParams{
		Pairs: &api.CatalogPairId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetPairsWithoutParams(t *testing.T) {
	data := getCatalogAssetPairsResponse(`{"data":[{"pair": "1inch-btc", "metrics": [{ "metric": "volume_reported_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_reported_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }] }, { "pair": "1inch-busd", "metrics": [{ "metric": "volume_reported_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2021-02-23T12:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }, { "metric": "volume_reported_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2021-02-24T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2021-02-24T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2021-02-23T12:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogAssetPairsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetPairsWithParams(t *testing.T) {
	data := getCatalogAssetPairsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogAssetPairsParams{
		Pairs: &api.CatalogPairId{`1inch-btc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogAssetPairsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available Institution

func TestInstitutionNotFoundForGetCatalogInstitutionsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogInstitutionsParams{
		Institutions: &api.CatalogInstitutionId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogInstitutionsWithoutParams(t *testing.T) {
	data := getCatalogInstitutionsResponse(`{"data":[{"institution":"grayscale","metrics":[{"metric":"batfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2021-03-17T00:00:00.000000000Z","max_time":"2022-02-11T00:00:00.000000000Z"}]},{"metric":"xlmfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2018-12-06T00:00:00.000000000Z","max_time":"2021-10-19T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogInstitutionsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogInstitutionsWithParams(t *testing.T) {
	data := getCatalogInstitutionsResponse(`{"data":[{"institution":"grayscale","metrics":[{"metric":"batfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2021-03-17T00:00:00.000000000Z","max_time":"2022-02-11T00:00:00.000000000Z"}]},{"metric":"xlmfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2018-12-06T00:00:00.000000000Z","max_time":"2021-10-19T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogInstitutionsParams{
		Institutions: &api.CatalogInstitutionId{`grayscale`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogInstitutionsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogInstitutionsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available indexes

func TestIndexesNotFoundForGetCatalogIndexesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'indexes'. Value 'asdgwav' is not supported.`)
	param := api.GetCatalogIndexesParams{
		Indexes: &api.CatalogIndexId{`asdgwav`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/indexes`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogIndexesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func buildErrorMessage(message, errorType string) api.ErrorResponse {
	errObject := api.ErrorObject{
		Message: &message,
		Type:    errorType,
	}
	errResponse := api.ErrorResponse{errObject}
	return errResponse
}

func getCatalogMetricsResponse(res string) *api.MetricsResponse {
	responseStruct := api.MetricsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.MetricsResponse{}
	}
	return &responseStruct
}

func getCatalogAssetResponse(res string) *api.AssetsResponse {
	responseStruct := api.AssetsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.AssetsResponse{}
	}
	return &responseStruct
}

func getCatalogExchangeResponse(res string) *api.ExchangesResponse {
	responseStruct := api.ExchangesResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.ExchangesResponse{}
	}
	return &responseStruct
}

func getCatalogExchangeAssetsResponse(res string) *api.ExchangeAssetsResponse {
	responseStruct := api.ExchangeAssetsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.ExchangeAssetsResponse{}
	}
	return &responseStruct
}

func getCatalogAssetPairsResponse(res string) *api.PairsResponse {
	responseStruct := api.PairsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.PairsResponse{}
	}
	return &responseStruct
}

func getCatalogInstitutionsResponse(res string) *api.InstitutionsResponse {
	responseStruct := api.InstitutionsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.InstitutionsResponse{}
	}
	return &responseStruct
}
